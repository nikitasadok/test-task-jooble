package calculator

import (
	"calculator/stack"
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"
)

type calculator struct {
	expression string
	operations *stack.Stack
	values     *stack.Stack
}

func NewCalculator(expression string) Calculator {
	values := stack.NewStack()
	operations := stack.NewStack()
	return &calculator{expression: expression, values: values, operations: operations}
}

func (c *calculator) getPrecedence(ch string) int {
	if ch == "+" || ch == "-" {
		return 1
	}

	if ch == "*" || ch == "/" {
		return 2
	}

	return -1
}

func (c *calculator) evaluate() (float64, error) {
	s := c.expression
	sReader := strings.NewReader(s)
	for {
		ch, err := sReader.ReadByte()
		if err != nil {
			log.Println("evaluate() -> error reading file, possible EOF")
			break
		}

		switch ch {
		case ' ':
			continue
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			err := sReader.UnreadByte()
			if err != nil {
				log.Println("evaluate() -> error unreading byte err: ", err)
			}

			n, err := c.processNumber(sReader)
			if err != nil {
				return 0.0, err
			}
			c.values.Push(n)
		case '(':
			c.operations.Push(string(ch))
		case ')':
			for !c.operations.IsEmpty() && c.operations.Peek() != "(" {
				c.performCalculation()
			}

			if !c.operations.IsEmpty() {
				c.operations.Pop()
			}
		case '-':
			err := c.processMinus(sReader, ch)
			if err != nil {
				return 0.0, err
			}
		case '+', '*', '/':
			for !c.operations.IsEmpty() && c.getPrecedence(string(ch)) <=
				c.getPrecedence(c.operations.Peek()) {
				c.performCalculation()
			}

			c.operations.Push(string(ch))
		default:
			log.Println("evaluate() -> unknown symbol")
			return 0.0, ErrInvalidExpression
		}
	}

	for !c.operations.IsEmpty() {
		op2 := c.values.Pop()
		op1 := c.values.Pop()

		res, err := c.getResult(op1, op2)
		if err != nil {
			continue
		}
		c.values.Push(fmt.Sprint(res))
	}
	n, err := strconv.ParseFloat(c.values.Peek(), 64)
	if err != nil {
		log.Println("evaluate() -> wrong value on top of the stack!")
		return 0.0, err
	}
	return n, nil
}

func (c *calculator) getResult(arg1, arg2 string) (float64, error) {
	op1, err := c.convertOperand(arg1)
	if err != nil {
		log.Println("getResult() -> error converting operand from string, operand =", op1)
	}

	op2, err := c.convertOperand(arg2)
	if err != nil {
		log.Println("getResult() -> error converting operand from string, operand =", op2)
	}

	operator := c.operations.Pop()
	var res float64
	res = c.doOperation(op1, op2, operator)
	return res, nil
}

func (c *calculator) processMinus(sReader *strings.Reader, ch byte) error {
	chNext, err := sReader.ReadByte()
	if err != nil {
		log.Println("processOperator() -> error reading next symbol after arithmetic operator, err:", err)
		return ErrInvalidExpression
	}

	if c.isNumber(chNext) {
		sReader.UnreadByte()
		n, err := c.processNumber(sReader)
		if err != nil {
			log.Println("processOperator() -> error processing number from string, err:", err)
		}

		n = "-" + n
		c.values.Push(n)
		return nil
	}

	if chNext == ' ' {
		for !c.operations.IsEmpty() && c.getPrecedence(string(ch)) <=
			c.getPrecedence(c.operations.Peek()) {
			c.performCalculation()
		}

		c.operations.Push(string(ch))

		return nil
	}

	return ErrInvalidExpression
}

func (c *calculator) processNumber(reader *strings.Reader) (string, error) {
	var num string
	var last byte
	for {
		ch, err := reader.ReadByte()
		if err != nil {
			log.Println("processNumber() -> Error reading next digit, error:", err)
			return num, nil
		}

		if !c.isNumber(ch) {
			last = ch
			break
		}

		num += string(ch)
	}

	if !c.isValidDelim(last) {
		return "", ErrInvalidExpression
	}

	reader.UnreadByte()

	return num, nil
}

func (c *calculator) isValidDelim(ch byte) bool {
	validDelims := []byte{' ', ')', 0}
	for _, val := range validDelims {
		if ch == val {
			return true
		}
	}

	return false
}

func (c *calculator) convertOperand(op string) (float64, error) {
	val, err := strconv.ParseFloat(op, 64)
	if err != nil {
		log.Printf("convertOperand() -> error converting operand: operand = %s, err = %s\n", op, err)
		return .0, err
	}
	return val, nil
}

func (c *calculator) doOperation(op1, op2 float64, operator string) float64 {
	var res float64
	switch operator {
	case "+":
		res = op1 + op2
	case "-":
		res = op1 - op2
	case "*":
		res = op1 * op2
	case "/":
		res = c.handleDivision(op1, op2)
	default:
		res = op1
	}

	return res
}

func (c *calculator) handleDivision(dividend, divisor float64) float64 {
	if divisor == -0 {
		return math.Inf(-1)
	}

	return dividend / divisor
}

func (c *calculator) performCalculation() {
	op2 := c.values.Pop()
	op1 := c.values.Pop()

	res, err := c.getResult(op1, op2)
	if err != nil {
		return
	}
	c.values.Push(fmt.Sprint(res))
}

func (c *calculator) Evaluate() (float64, error) {
	return c.evaluate()
}

func (c *calculator) isNumber(ch byte) bool {
	return ch >= '0' && ch <= '9'
}
