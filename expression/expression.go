package expression

type Visitor interface {
	VisitBinary(expression *Expression) interface{}
	VisitGrouping(expression *Expression) interface{}
	VisitLiteral(expression *Expression) interface{}
	VisitUnary(expression *Expression) interface{}
}

type Expression interface {
	Accept(v Visitor) interface{}
}
