package generators

//AccessorsGenerator should implements all the method necessary for generate the code
//that can handle an omnibuff schema
type AccessorsGenerator interface {
	//Should return a function for access to a boolean
	Scalar(mutationAllowed bool) string
	ScalarVector(mutationAllowed bool) string
	Table(mutationAllowed bool) string
	TableUnion(mutationAllowed bool) string
	Resource(mutationAllowed bool) string
	ResourceUnion(mutationAllowed bool) string
	StringVector(mutationAllowed bool) string
	TableVector(mutationAllowed bool) string
	UnionVector(mutationAllowed bool) string
	ResourceVector(mutationAllowed bool) string
}
