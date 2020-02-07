package ameda

import (
	"math"
	"sort"
	"strconv"
)

// Int64Slice int64 slice object
type Int64Slice []int64

// NewInt64Slice creates an Int64Slice object.
func NewInt64Slice(a []int64) *Int64Slice {
	i := Int64Slice(a)
	return &i
}

// Copy creates a copy of the int64 slice.
func (i Int64Slice) Copy() []int64 {
	b := make([]int64, len(i))
	copy(b, i)
	return b
}

// Strings converts int64 slice to string slice.
func (i Int64Slice) Strings() []string {
	r := make([]string, len(i))
	for k, v := range i {
		r[k] = strconv.FormatInt(v, 10)
	}
	return r
}

// Float32s converts int64 slice to float32 slice.
func (i Int64Slice) Float32s() []float32 {
	r := make([]float32, len(i))
	for k, v := range i {
		r[k] = float32(v)
	}
	return r
}

// Float64s converts int64 slice to float64 slice.
func (i Int64Slice) Float64s() []float64 {
	r := make([]float64, len(i))
	for k, v := range i {
		r[k] = float64(v)
	}
	return r
}

// Ints converts int64 slice to int slice.
func (i Int64Slice) Ints() ([]int, error) {
	r := make([]int, len(i))
	if is64BitPlatform {
		for k, v := range i {
			r[k] = int(v)
		}
	} else {
		for k, v := range i {
			if v > math.MaxInt32 {
				return nil, errOverflowValue
			}
			r[k] = int(v)
		}
	}
	return r, nil
}

// Int8s converts int64 slice to int8 slice.
func (i Int64Slice) Int8s() ([]int8, error) {
	r := make([]int8, len(i))
	for k, v := range i {
		if v > math.MaxInt8 || v < math.MinInt8 {
			return nil, errOverflowValue
		}
		r[k] = int8(v)
	}
	return r, nil
}

// Int16s converts int64 slice to int16 slice.
func (i Int64Slice) Int16s() ([]int16, error) {
	r := make([]int16, len(i))
	for k, v := range i {
		if v > math.MaxInt16 || v < math.MinInt16 {
			return nil, errOverflowValue
		}
		r[k] = int16(v)
	}
	return r, nil
}

// Int32s converts int64 slice to int32 slice.
func (i Int64Slice) Int32s() ([]int32, error) {
	r := make([]int32, len(i))
	for k, v := range i {
		if v > math.MaxInt32 || v < math.MinInt32 {
			return nil, errOverflowValue
		}
		r[k] = int32(v)
	}
	return r, nil
}

// Int64s converts to []int64.
func (i Int64Slice) Int64s() []int64 {
	return []int64(i)
}

// Uint8s converts int64 slice to uint slice.
func (i Int64Slice) Uints() ([]uint, error) {
	r := make([]uint, len(i))
	if is64BitPlatform {
		for k, v := range i {
			if v < 0 {
				return nil, errNegativeValue
			}
			r[k] = uint(v)
		}
	} else {
		for k, v := range i {
			if v < 0 {
				return nil, errNegativeValue
			}
			if v > math.MaxUint32 {
				return nil, errOverflowValue
			}
			r[k] = uint(v)
		}
	}
	return r, nil
}

// Uint8s converts int64 slice to uint8 slice.
func (i Int64Slice) Uint8s() ([]uint8, error) {
	r := make([]uint8, len(i))
	for k, v := range i {
		if v < 0 {
			return nil, errNegativeValue
		}
		if v > math.MaxUint8 {
			return nil, errOverflowValue
		}
		r[k] = uint8(v)
	}
	return r, nil
}

// Uint16s converts int64 slice to uint16 slice.
func (i Int64Slice) Uint16s() ([]uint16, error) {
	r := make([]uint16, len(i))
	for k, v := range i {
		if v < 0 {
			return nil, errNegativeValue
		}
		if v > math.MaxUint16 {
			return nil, errOverflowValue
		}
		r[k] = uint16(v)
	}
	return r, nil
}

// Uint32s converts int64 slice to uint32 slice.
func (i Int64Slice) Uint32s() ([]uint32, error) {
	r := make([]uint32, len(i))
	for k, v := range i {
		if v < 0 {
			return nil, errNegativeValue
		}
		if v > math.MaxUint32 {
			return nil, errOverflowValue
		}
		r[k] = uint32(v)
	}
	return r, nil
}

// Uint64s converts int64 slice to uint64 slice.
func (i Int64Slice) Uint64s() ([]uint64, error) {
	r := make([]uint64, len(i))
	for k, v := range i {
		if v < 0 {
			return nil, errNegativeValue
		}
		r[k] = uint64(v)
	}
	return r, nil
}

// Concat is used to merge two or more slices.
// This method does not change the existing slices, but instead returns a new slice.
func (i Int64Slice) Concat(a ...[]int64) []int64 {
	totalLen := len(i)
	for _, v := range a {
		totalLen += len(v)
	}
	ret := make([]int64, totalLen)
	n := copy(ret, i)
	dst := ret[n:]
	for _, v := range a {
		n := copy(dst, v)
		dst = dst[n:]
	}
	return ret
}

// CopyWithin copies part of an slice to another location in the current slice.
// @target
//  Zero-based index at which to copy the sequence to. If negative, target will be counted from the end.
// @start
//  Zero-based index at which to start copying elements from. If negative, start will be counted from the end.
// @end
//  Zero-based index at which to end copying elements from. CopyWithin copies up to but not including end.
//  If negative, end will be counted from the end.
//  If end is omitted, CopyWithin will copy until the last index (default to len(s)).
func (i Int64Slice) CopyWithin(target, start int, end ...int) {
	target = fixIndex(len(i), target, true)
	if target == len(i) {
		return
	}
	sub := i.Slice(start, end...)
	for k, v := range sub {
		i[target+k] = v
	}
}

// Every tests whether all elements in the slice pass the test implemented by the provided function.
// NOTE:
//  Calling this method on an empty slice will return true for any condition!
func (i Int64Slice) Every(fn func(curr Int64Slice, k int, v int64) bool) bool {
	for k, v := range i {
		if !fn(i, k, v) {
			return false
		}
	}
	return true
}

// Fill changes all elements in the current slice to a value, from a start index to an end index.
// @value
//  Zero-based index at which to copy the sequence to. If negative, target will be counted from the end.
// @start
//  Zero-based index at which to start copying elements from. If negative, start will be counted from the end.
// @end
//  Zero-based index at which to end copying elements from. CopyWithin copies up to but not including end.
//  If negative, end will be counted from the end.
//  If end is omitted, CopyWithin will copy until the last index (default to len(s)).
func (i Int64Slice) Fill(value int64, start int, end ...int) {
	fixedStart, fixedEnd, ok := fixRange(len(i), start, end...)
	if !ok {
		return
	}
	for k := fixedStart; k < fixedEnd; k++ {
		i[k] = value
	}
}

// Filter creates a new slice with all elements that pass the test implemented by the provided function.
func (i Int64Slice) Filter(fn func(curr Int64Slice, k int, v int64) bool) []int64 {
	ret := make([]int64, 0)
	for k, v := range i {
		if fn(i, k, v) {
			ret = append(ret, v)
		}
	}
	return ret
}

// Find returns the key-value of the first element in the provided slice that satisfies the provided testing function.
// NOTE:
//  If not found, k = -1
func (i Int64Slice) Find(fn func(curr Int64Slice, k int, v int64) bool) (k int, v int64) {
	for k, v := range i {
		if fn(i, k, v) {
			return k, v
		}
	}
	return -1, 0
}

// Includes determines whether an slice includes a certain value among its entries.
// @fromIndex
//  The index to start the search at. Defaults to 0.
func (i Int64Slice) Includes(valueToFind int64, fromIndex ...int) bool {
	return i.IndexOf(valueToFind, fromIndex...) > -1
}

// IndexOf returns the first index at which a given element can be found in the slice, or -1 if it is not present.
// @fromIndex
//  The index to start the search at. Defaults to 0.
func (i Int64Slice) IndexOf(searchElement int64, fromIndex ...int) int {
	idx := getFromIndex(len(i), fromIndex...)
	for k, v := range i[idx:] {
		if searchElement == v {
			return k + idx
		}
	}
	return -1
}

// LastIndexOf returns the last index at which a given element can be found in the slice, or -1 if it is not present.
// @fromIndex
//  The index to start the search at. Defaults to 0.
func (i Int64Slice) LastIndexOf(searchElement int64, fromIndex ...int) int {
	idx := getFromIndex(len(i), fromIndex...)
	for k := len(i) - 1; k >= idx; k-- {
		if searchElement == i[k] {
			return k
		}
	}
	return -1
}

// Map creates a new slice populated with the results of calling a provided function
// on every element in the calling slice.
func (i Int64Slice) Map(fn func(curr Int64Slice, k int, v int64) int64) []int64 {
	ret := make([]int64, len(i))
	for k, v := range i {
		ret[k] = fn(i, k, v)
	}
	return ret
}

// Pop removes the last element from an slice and returns that element.
// This method changes the length of the slice.
func (i *Int64Slice) Pop() (int64, bool) {
	a := *i
	if len(a) == 0 {
		return 0, false
	}
	lastIndex := len(a) - 1
	last := a[lastIndex]
	a = a[:lastIndex]
	*i = a[:len(a):len(a)]
	return last, true
}

// Push adds one or more elements to the end of an slice and returns the new length of the slice.
func (i *Int64Slice) Push(element ...int64) int {
	*i = append(*i, element...)
	return len(*i)
}

// PushOnce adds one or more new elements that do not exist in the current slice at the end
// and returns the new length of the slice.
func (i *Int64Slice) PushOnce(element ...int64) int {
	a := *i
L:
	for _, v := range element {
		for _, vv := range a {
			if vv == v {
				continue L
			}
		}
		a = append(a, v)
	}
	*i = a
	return len(a)
}

// Reduce executes a reducer function (that you provide) on each element of the slice,
// resulting in a single output value.
// @accumulator
//  The accumulator accumulates callback's return values.
//  It is the accumulated value previously returned in the last invocation of the callback—or initialValue,
//  if it was supplied (see below).
// @initialValue
//  A value to use as the first argument to the first call of the callback.
//  If no initialValue is supplied, the first element in the slice will be used and skipped.
func (i Int64Slice) Reduce(
	fn func(curr Int64Slice, k int, v int64, accumulator int64) int64, initialValue ...int64,
) int64 {
	if len(i) == 0 {
		return 0
	}
	start := 0
	acc := i[start]
	if len(initialValue) > 0 {
		acc = initialValue[0]
	} else {
		start += 1
	}
	for k := start; k < len(i); k++ {
		acc = fn(i, k, i[k], acc)
	}
	return acc
}

// ReduceRight applies a function against an accumulator and each value of the slice (from right-to-left)
// to reduce it to a single value.
// @accumulator
//  The accumulator accumulates callback's return values.
//  It is the accumulated value previously returned in the last invocation of the callback—or initialValue,
//  if it was supplied (see below).
// @initialValue
//  A value to use as the first argument to the first call of the callback.
//  If no initialValue is supplied, the first element in the slice will be used and skipped.
func (i Int64Slice) ReduceRight(
	fn func(curr Int64Slice, k int, v int64, accumulator int64) int64, initialValue ...int64,
) int64 {
	if len(i) == 0 {
		return 0
	}
	end := len(i) - 1
	acc := i[end]
	if len(initialValue) > 0 {
		acc = initialValue[0]
	} else {
		end -= 1
	}
	for k := end; k >= 0; k-- {
		acc = fn(i, k, i[k], acc)
	}
	return acc
}

// Reverse reverses an slice in place.
func (i Int64Slice) Reverse() {
	first := 0
	last := len(i) - 1
	for first < last {
		i[first], i[last] = i[last], i[first]
		first++
		last--
	}
}

// Shift removes the first element from an slice and returns that removed element.
// This method changes the length of the slice.
func (i *Int64Slice) Shift() (int64, bool) {
	a := *i
	if len(a) == 0 {
		return 0, false
	}
	first := a[0]
	a = a[1:]
	*i = a[:len(a):len(a)]
	return first, true
}

// Slice returns a copy of a portion of an slice into a new slice object selected
// from begin to end (end not included) where begin and end represent the index of items in that slice.
// The original slice will not be modified.
func (i Int64Slice) Slice(begin int, end ...int) []int64 {
	fixedStart, fixedEnd, ok := fixRange(len(i), begin, end...)
	if !ok {
		return []int64{}
	}
	return i[fixedStart:fixedEnd].Copy()
}

// Some tests whether at least one element in the slice passes the test implemented by the provided function.
// NOTE:
//  Calling this method on an empty slice returns false for any condition!
func (i Int64Slice) Some(fn func(curr Int64Slice, k int, v int64) bool) bool {
	for k, v := range i {
		if fn(i, k, v) {
			return true
		}
	}
	return false
}

// Len is the number of elements in the collection.
func (i Int64Slice) Len() int {
	return len(i)
}

// Less reports whether the element with
// index m should sort before the element with index n.
func (i Int64Slice) Less(m, n int) bool {
	return i[m] < i[n]
}

// Swap swaps the elements with indexes m and n.
func (i Int64Slice) Swap(m, n int) {
	i[m], i[n] = i[n], i[m]
}

// Sort sorts the elements of an slice in place and returns the sorted slice.
func (i Int64Slice) Sort() {
	sort.Sort(i)
}

// Splice changes the contents of an slice by removing or replacing
// existing elements and/or adding new elements in place.
func (i *Int64Slice) Splice(start, deleteCount int, items ...int64) {
	a := *i
	if deleteCount < 0 {
		deleteCount = 0
	}
	start, end, _ := fixRange(len(a), start, start+1+deleteCount)
	deleteCount = end - start - 1
	for k := 0; k < len(items); k++ {
		if deleteCount > 0 {
			// replace
			a[start] = items[k]
			deleteCount--
			start++
		} else {
			// insert
			lastSlice := a[start:].Copy()
			items = items[k:]
			a = append(a[:start], items...)
			a = append(a[:start+len(items)], lastSlice...)
			*i = a[:len(a):len(a)]
			return
		}
	}
	if deleteCount > 0 {
		a = append(a[:start], a[start+1+deleteCount:]...)
	}
	*i = a[:len(a):len(a)]
}

// Unshift adds one or more elements to the beginning of an slice and returns the new length of the slice.
func (i *Int64Slice) Unshift(element ...int64) int {
	*i = append(element, *i...)
	return len(*i)
}

// UnshiftOnce adds one or more new elements that do not exist in the current slice to the beginning
// and returns the new length of the slice.
func (i *Int64Slice) UnshiftOnce(element ...int64) int {
	a := *i
	if len(element) == 0 {
		return len(a)
	}
	m := make(map[int64]bool, len(element))
	r := make([]int64, 0, len(a)+len(element))
L:
	for _, v := range element {
		if m[v] {
			continue
		}
		m[v] = true
		for _, vv := range a {
			if vv == v {
				continue L
			}
		}
		r = append(r, v)
	}
	r = append(r, a...)
	*i = r[:len(r):len(r)]
	return len(r)
}

// Distinct creates an new slice in place set that removes the same elements
// and returns the new length of the slice.
func (i *Int64Slice) Distinct() int {
	a := (*i)[:0]
	m := make(map[int64]bool, len(a))
	for _, v := range *i {
		if m[v] {
			continue
		}
		a = append(a, v)
		m[v] = true
	}
	n := len(m)
	*i = a[:n:n]
	return n
}

// RemoveOne removes the first matched elements from the slice,
// and returns the new length of the slice.
func (i *Int64Slice) RemoveOne(element ...int64) int {
	a := *i
	m := make(map[int64]bool, len(element))
	for _, v := range element {
		if m[v] {
			continue
		}
		m[v] = true
		for kk, vv := range a {
			if vv == v {
				a = append(a[:kk], a[kk+1:]...)
				break
			}
		}
	}
	n := len(a)
	*i = a[:n:n]
	return n
}

// RemoveAll removes all the elements from the slice,
// and returns the new length of the slice.
func (i *Int64Slice) RemoveAll(element ...int64) int {
	a := *i
	m := make(map[int64]bool, len(element))
	for _, v := range element {
		if m[v] {
			continue
		}
		m[v] = true
		for kk, vv := range a {
			if vv == v {
				a = append(a[:kk], a[kk+1:]...)
			}
		}
	}
	n := len(a)
	*i = a[:n:n]
	return n
}