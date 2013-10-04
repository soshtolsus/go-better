package WorkerPool

import "testing"
import . "launchpad.net/gocheck"

func Test(t *testing.T) {
	TestingT(t)
}

//================================

type TheOnlySuite struct {}
var _ = Suite(new(TheOnlySuite))

//================================

func (self *TheOnlySuite) TestLength(c *C) {
	fixture := new(dlList)

	fixture.length = uint(8)

	c.Check(fixture.Length(), Equals, uint(8))
}

func (self *TheOnlySuite) TestPushPop(c *C) {
	var returned *int

	fixture := new(dlList)

	fixture.LPush(func(){temp := 1; returned = &temp})
	c.Check(fixture.Length(), Equals, uint(1))
	fixture.LPush(func(){temp := 2; returned = &temp})
	c.Check(fixture.Length(), Equals, uint(2))
	fixture.LPush(func(){temp := 3; returned = &temp})
	c.Check(fixture.Length(), Equals, uint(3))

	fixture.RPop().Value()()
	c.Check(*returned, Equals, 1)
	c.Check(fixture.Length(), Equals, uint(2))
	fixture.RPop().Value()()
	c.Check(*returned, Equals, 2)
	c.Check(fixture.Length(), Equals, uint(1))
	fixture.RPop().Value()()
	c.Check(*returned, Equals, 3)
	c.Check(fixture.Length(), Equals, uint(0))
}

func (self *TheOnlySuite) TestPopWhenEmpty(c *C) {
	fixture := new(dlList)

	popped_item := fixture.RPop()

	c.Check(popped_item, IsNil)
}
