package ipam

import (
	"fmt"
)

// IPRange represents an IP range of [start, end]
type IPRange struct {
	start, end IP
}

func NewIPRange(start, end IP) *IPRange {
	return &IPRange{start, end}
}

func (r *IPRange) Clone() *IPRange {
	return NewIPRange(r.start, r.end)
}

func (r *IPRange) Start() IP {
	return r.start
}

func (r *IPRange) End() IP {
	return r.end
}

func (r *IPRange) SetStart(ip IP) {
	r.start = ip
}

func (r *IPRange) SetEnd(ip IP) {
	r.end = ip
}

func (r *IPRange) Contains(ip IP) bool {
	return !r.start.GreaterThan(ip) && !r.end.LessThan(ip)
}

func (r *IPRange) Add(ip IP) bool {
	if newStart := r.start.Sub(1); newStart.Equal(ip) {
		r.start = newStart
		return true
	}
	if newEnd := r.end.Add(1); newEnd.Equal(ip) {
		r.end = newEnd
		return true
	}

	return false
}

func (r *IPRange) Remove(ip IP) ([]*IPRange, bool) {
	if !r.Contains(ip) {
		return nil, false
	}

	ret := make([]*IPRange, 0, 2)
	r1 := NewIPRange(r.start, ip.Sub(1))
	r2 := NewIPRange(ip.Add(1), r.end)
	if !r1.start.GreaterThan(r1.end) {
		ret = append(ret, r1)
	}
	if !r2.start.GreaterThan(r2.end) {
		ret = append(ret, r2)
	}

	return ret, true
}

func (r *IPRange) String() string {
	if r.start.Equal(r.end) {
		return r.start.String()
	}
	return fmt.Sprintf("%s-%s", r.start.String(), r.end.String())
}