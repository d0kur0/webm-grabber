package types

type Board string

func (board Board) String() string {
	return string(board)
}
