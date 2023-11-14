// Package consensus implements consensus algorithm
package consensus

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	pb "github.com/structx/structs/pb/protos/raft"
	"google.golang.org/grpc"
)

type state int

const (
	// follower state
	follower state = iota
	// candidate state
	candidate
	// leader state
	leader
)

type logEntry struct {
	term    int64
	index   int64
	command []byte
}

type stateMachine struct{}

// RaftNode is a node in a raft cluster
type RaftNode struct {
	pb.UnimplementedRaftServiceServer

	id          uuid.UUID     // id of the node
	state       state         // state of the node
	sm          *stateMachine // state machine
	term        int64         // current term
	votedFor    string        // last vote for
	log         []*logEntry   // log entries
	commitIndex int64         // index of highest log entry known to be committed
	lastApplied int64         // index of highest log entry applied to state machine
	nextIndex   []int64       // for each server, index of the next log entry to send to that server
	matchIndex  []int64       // for each server, index of highest log entry known to be replicated on server

	peers   map[string]string // peers in the cluster
	timeout int               // election timeout
}

// NewRaftNode creates a new raft node
func NewRaftNode() *RaftNode {
	return &RaftNode{
		id:          uuid.New(),
		state:       follower,
		term:        0,
		votedFor:    "",
		sm:          &stateMachine{},
		log:         make([]*logEntry, 0),
		commitIndex: 0,
		lastApplied: 0,
		nextIndex:   make([]int64, 0),
		matchIndex:  make([]int64, 0),
		peers:       make(map[string]string),
		timeout:     300,
	}
}

// GetState returns the state of the node
func (r *RaftNode) GetState() string {
	switch r.state {
	case follower:
		return "follower"
	case candidate:
		return "candidate"
	case leader:
		return "leader"
	default:
		return ""
	}
}

func (r *RaftNode) StartWithContext(ctx context.Context) {

}

func (r *RaftNode) Start() {}

func (r *RaftNode) run(ctx context.Context) {

	for {

		select {
		case <-ctx.Done():
			return
		default:
			switch r.state {
			case follower:
				if r.term > 0 {
					// start election
					r.state = candidate

					// increment term
					r.term++

					// vote for self
					r.votedFor = r.id.String()

					// reset election timeout
					r.timeout = 300
				}

			case candidate:

				// reset election timeout
				r.timeout = 300

				var count int

				// send vote requests
				for _, peer := range r.peers {
					vote, err := r.sendVoteRequest(ctx, peer)
					if err != nil {
						fmt.Printf("failed to send vote request: %v", err)
					}

					if vote {
						// increment vote count
						count++
					}
				}

				// check if majority of votes received
				if count > len(r.peers)/2 {
					// become leader
					r.state = leader
				}

			case leader:
				// send heartbeats
				for _, peer := range r.peers {
					err := r.sendHeartbeart(ctx, peer)
					if err != nil {
						fmt.Printf("failed to send heartbeat: %v", err)
					}
				}
			default:
				// do nothing
			}
		}
	}
}

// AppendEntries appends entries to the log
func (r *RaftNode) AppendEntries(_ context.Context, in *pb.AppendEntriesRequest) (*pb.AppendEntriesResponse, error) {

	switch r.state {
	case leader:
		// received heartbeat from new leader
		if in.Term > r.term {
			// update term
			r.term = in.Term
			// reset voted for
			r.votedFor = ""
			// become follower
			r.state = follower
			// reset election timeout
			r.timeout = 300
			return &pb.AppendEntriesResponse{Term: r.term, Success: true}, nil
		}
	case follower, candidate:
		if in.Term < r.term {
			return &pb.AppendEntriesResponse{Term: r.term, Success: false}, nil
		}

		if len(r.log) < int(in.PrevLogIndex) {
			return &pb.AppendEntriesResponse{Term: r.term, Success: false}, nil
		}

		if r.log[in.PrevLogIndex].term != in.PrevLogIndex {
			return &pb.AppendEntriesResponse{Term: r.term, Success: false}, nil
		}

		if r.log[in.PrevLogIndex].term != in.PrevLogTerm {
			r.log = r.log[:in.PrevLogIndex]
		}

		// append any new entries not already in the log
		if len(r.log) > int(in.PrevLogIndex) {
			for _, entry := range in.Entries {
				r.log = append(r.log, &logEntry{
					term:    entry.Term,
					index:   entry.Index,
					command: []byte(entry.Command),
				})
			}
		}

		if in.LeaderCommit > r.commitIndex {
			r.commitIndex = in.LeaderCommit
		} else {
			r.commitIndex = int64(len(r.log))
		}
	}

	return &pb.AppendEntriesResponse{Term: r.term, Success: true}, nil
}

// RequestVote requests vote from other nodes
func (r *RaftNode) RequestVote(_ context.Context, in *pb.RequestVoteRequest) (*pb.RequestVoteResponse, error) {

	if in.Term < r.term {
		return &pb.RequestVoteResponse{Term: r.term, VoteGranted: false}, nil
	}

	if r.votedFor == "" || r.votedFor == in.CandidateId {
		r.votedFor = in.CandidateId
		return &pb.RequestVoteResponse{Term: r.term, VoteGranted: true}, nil
	}

	return &pb.RequestVoteResponse{Term: r.term, VoteGranted: false}, nil
}

func (r *RaftNode) sendHeartbeart(ctx context.Context, address string) error {

	timeout, cancel := context.WithTimeout(ctx, time.Millisecond*time.Duration(r.timeout))
	defer cancel()

	conn, err := grpc.DialContext(timeout, address)
	if err != nil {
		return fmt.Errorf("failed to dial: %v", err)
	}
	defer conn.Close()

	client := pb.NewRaftServiceClient(conn)

	_, err = client.AppendEntries(timeout, &pb.AppendEntriesRequest{
		Term:         r.term,
		LeaderId:     r.id.String(),
		PrevLogIndex: r.commitIndex,
		PrevLogTerm:  r.log[r.commitIndex].term,
		Entries:      nil,
		LeaderCommit: r.commitIndex,
	})
	if err != nil {
		return fmt.Errorf("failed to send heartbeat: %v", err)
	}

	return nil
}

func (r *RaftNode) sendVoteRequest(ctx context.Context, address string) (bool, error) {

	timeout, cancel := context.WithTimeout(ctx, time.Millisecond*time.Duration(r.timeout))
	defer cancel()

	conn, err := grpc.DialContext(timeout, address)
	if err != nil {
		return false, fmt.Errorf("failed to dial: %v", err)
	}
	defer conn.Close()

	client := pb.NewRaftServiceClient(conn)

	resp, err := client.RequestVote(timeout, &pb.RequestVoteRequest{
		Term:         r.term,
		CandidateId:  r.id.String(),
		LastLogIndex: r.commitIndex,
		LastLogTerm:  r.log[r.commitIndex].term,
	})
	if err != nil {
		return false, fmt.Errorf("failed to send vote request: %v", err)
	}

	return resp.VoteGranted, nil
}
