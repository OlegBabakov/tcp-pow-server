package pow

type POW interface {
	Verifier
	Solver
}

// Verifier is an interface of pow verifier.
type Verifier interface {
	Challenge() []byte
	Verify(challenge, solution []byte) error
}

// Solver is an interface of pow solver.
type Solver interface {
	Solve(challenge []byte) []byte
}
