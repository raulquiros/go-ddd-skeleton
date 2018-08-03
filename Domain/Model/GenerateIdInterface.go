package Model

type GenerateIdInterface interface {
	Generate() (string, error)
}
