package request

type PersonRequestUri struct {
	PersonId int32 `uri:"person-id" binding:"required"`
}
