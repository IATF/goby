//+build !release

package ast

import (
	"testing"
)

/*
 BaseNode
*/

// IsAssignExpression fails the test and returns nil by default
func (b *BaseNode) IsArrayExpression(t *testing.T) *TestableArrayExpression {
	t.Fatalf(nodeFailureMsgFormat, "array expression", b)
	return nil
}

// IsAssignExpression fails the test and returns nil by default
func (b *BaseNode) IsAssignExpression(t *testing.T) (ae *TestableAssignExpression) {
	t.Fatalf(nodeFailureMsgFormat, "assign expression", b)
	return
}

// IsCallExpression fails the test and returns nil by default
func (b *BaseNode) IsCallExpression(t *testing.T) (ce *TestableCallExpression) {
	t.Fatalf(nodeFailureMsgFormat, "call expression", b)
	return
}

// IsConditionalExpression fails the test and returns nil by default
func (b *BaseNode) IsConditionalExpression(t *testing.T) *TestableConditionalExpression {
	t.Fatalf(nodeFailureMsgFormat, "conditional expression", b)
	return nil
}

// IsConstant fails the test and returns nil by default
func (b *BaseNode) IsConstant(t *testing.T) (c *TestableConstant) {
	t.Fatalf(nodeFailureMsgFormat, "constant", b)
	return
}

// IsExpression fails the test and returns nil by default
func (b *BaseNode) IsExpression(t *testing.T) (te TestingExpression) {
	t.Helper()
	t.Fatalf(nodeFailureMsgFormat, "expression", b)
	return
}

// HashExpression fails the test and returns nil by default
func (b *BaseNode) IsHashExpression(t *testing.T) *TestableHashExpression {
	t.Fatalf(nodeFailureMsgFormat, "hash expression", b)
	return nil
}

// IsIdentifier fails the test and returns nil by default
func (b *BaseNode) IsIdentifier(t *testing.T) *TestableIdentifier {
	t.Fatalf(nodeFailureMsgFormat, "identifier", b)
	return nil
}

// IsIfExpression fails the test and returns nil by default
func (b *BaseNode) IsIfExpression(t *testing.T) *TestableIfExpression {
	t.Fatalf(nodeFailureMsgFormat, "if expression", b)
	return nil
}

// IsInfixExpression fails the test and returns nil by default
func (b *BaseNode) IsInfixExpression(t *testing.T) (ie *TestableInfixExpression) {
	t.Fatalf(nodeFailureMsgFormat, "infix expression", b)
	return
}

// IsIntegerLiteral fails the test and returns nil by default
func (b *BaseNode) IsIntegerLiteral(t *testing.T) (il *TestableIntegerLiteral) {
	t.Fatalf(nodeFailureMsgFormat, "integer literal", b)
	return
}

// IsStringLiteral fails the test and returns nil by default
func (b *BaseNode) IsStringLiteral(t *testing.T) (sl *StringLiteral) {
	t.Fatalf(nodeFailureMsgFormat, "string literal", b)
	return
}

// IsYieldExpression returns pointer of the receiver yield expression
func (b *BaseNode) IsYieldExpression(t *testing.T) (ye *YieldExpression) {
	t.Fatalf(nodeFailureMsgFormat, "yield expression", b)
	return
}

// IsArrayExpression returns pointer of the receiver array expression
func (ae *ArrayExpression) IsArrayExpression(t *testing.T) *TestableArrayExpression {
	return &TestableArrayExpression{ArrayExpression: ae, t: t}
}

// IsAssignExpression returns pointer of the receiver assign expression
func (ae *AssignExpression) IsAssignExpression(t *testing.T) *TestableAssignExpression {
	return &TestableAssignExpression{AssignExpression: ae, t: t}
}

// NameIs compares the assignment's variable name and expected name
func (ae *AssignExpression) NameIs(n string) bool {
	return ae.Variables[0].(testingNode).NameIs(n)
}

// IsCallExpression returns pointer of the receiver call expression
func (ce *CallExpression) IsCallExpression(t *testing.T) *TestableCallExpression {
	return &TestableCallExpression{CallExpression: ce, t: t}
}

func (ce *ConditionalExpression) IsConditionalExpression(t *testing.T) *TestableConditionalExpression {
	return &TestableConditionalExpression{ConditionalExpression: ce, t: t}
}

// IsConstant returns pointer of the current receiver constant
func (c *Constant) IsConstant(t *testing.T) *TestableConstant {
	return &TestableConstant{Constant: c, t: t}
}

// IsHashExpression returns pointer of the receiver hash expression
func (he *HashExpression) IsHashExpression(t *testing.T) *TestableHashExpression {
	return &TestableHashExpression{HashExpression: he, t: t}
}

// IsIdentifier returns pointer of the receiver identifier
func (i *Identifier) IsIdentifier(t *testing.T) *TestableIdentifier {
	return &TestableIdentifier{Identifier: i, t: t}
}

// NameIs compares the identifier's name and expected name
func (i *Identifier) NameIs(n string) bool {
	if i.Value == n {
		return true
	}

	return false
}

// IsIfExpression returns pointer of the receiver if expression
func (ie *IfExpression) IsIfExpression(t *testing.T) *TestableIfExpression {
	return &TestableIfExpression{IfExpression: ie, t: t}
}

// IsInfixExpression returns pointer of the receiver infix expression
func (ie *InfixExpression) IsInfixExpression(t *testing.T) *TestableInfixExpression {
	return &TestableInfixExpression{InfixExpression: ie, t: t}
}

// IsIntegerLiteral returns pointer of the receiver string literal
func (il *IntegerLiteral) IsIntegerLiteral(t *testing.T) *TestableIntegerLiteral {
	return &TestableIntegerLiteral{IntegerLiteral: il, t: t}
}

// IsStringLiteral returns pointer of the receiver string literal
func (sl *StringLiteral) IsStringLiteral(t *testing.T) *StringLiteral {
	return sl
}

// IsYieldExpression returns pointer of the receiver yield expression
func (ye *YieldExpression) IsYieldExpression(t *testing.T) *YieldExpression {
	return ye
}
