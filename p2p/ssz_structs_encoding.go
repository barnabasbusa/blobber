// Code generated by fastssz. DO NOT EDIT.
// Hash: 8a5198ac11e4b6cf73f165d62294a2dc05507241b7c70d66f75d6840c5660a02
package p2p

import (
	ssz "github.com/prysmaticlabs/fastssz"
)

// MarshalSSZ ssz marshals the Status object
func (s *Status) MarshalSSZ() ([]byte, error) {
	return ssz.MarshalSSZ(s)
}

// MarshalSSZTo ssz marshals the Status object to a target array
func (s *Status) MarshalSSZTo(buf []byte) (dst []byte, err error) {
	dst = buf

	// Field (0) 'ForkDigest'
	if size := len(s.ForkDigest); size != 4 {
		err = ssz.ErrBytesLengthFn("--.ForkDigest", size, 4)
		return
	}
	dst = append(dst, s.ForkDigest...)

	// Field (1) 'FinalizedRoot'
	if size := len(s.FinalizedRoot); size != 32 {
		err = ssz.ErrBytesLengthFn("--.FinalizedRoot", size, 32)
		return
	}
	dst = append(dst, s.FinalizedRoot...)

	// Field (2) 'FinalizedEpoch'
	dst = ssz.MarshalUint64(dst, s.FinalizedEpoch)

	// Field (3) 'HeadRoot'
	if size := len(s.HeadRoot); size != 32 {
		err = ssz.ErrBytesLengthFn("--.HeadRoot", size, 32)
		return
	}
	dst = append(dst, s.HeadRoot...)

	// Field (4) 'HeadSlot'
	dst = ssz.MarshalUint64(dst, s.HeadSlot)

	return
}

// UnmarshalSSZ ssz unmarshals the Status object
func (s *Status) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size != 84 {
		return ssz.ErrSize
	}

	// Field (0) 'ForkDigest'
	if cap(s.ForkDigest) == 0 {
		s.ForkDigest = make([]byte, 0, len(buf[0:4]))
	}
	s.ForkDigest = append(s.ForkDigest, buf[0:4]...)

	// Field (1) 'FinalizedRoot'
	if cap(s.FinalizedRoot) == 0 {
		s.FinalizedRoot = make([]byte, 0, len(buf[4:36]))
	}
	s.FinalizedRoot = append(s.FinalizedRoot, buf[4:36]...)

	// Field (2) 'FinalizedEpoch'
	s.FinalizedEpoch = ssz.UnmarshallUint64(buf[36:44])

	// Field (3) 'HeadRoot'
	if cap(s.HeadRoot) == 0 {
		s.HeadRoot = make([]byte, 0, len(buf[44:76]))
	}
	s.HeadRoot = append(s.HeadRoot, buf[44:76]...)

	// Field (4) 'HeadSlot'
	s.HeadSlot = ssz.UnmarshallUint64(buf[76:84])

	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the Status object
func (s *Status) SizeSSZ() (size int) {
	size = 84
	return
}

// HashTreeRoot ssz hashes the Status object
func (s *Status) HashTreeRoot() ([32]byte, error) {
	return ssz.HashWithDefaultHasher(s)
}

// HashTreeRootWith ssz hashes the Status object with a hasher
func (s *Status) HashTreeRootWith(hh *ssz.Hasher) (err error) {
	indx := hh.Index()

	// Field (0) 'ForkDigest'
	if size := len(s.ForkDigest); size != 4 {
		err = ssz.ErrBytesLengthFn("--.ForkDigest", size, 4)
		return
	}
	hh.PutBytes(s.ForkDigest)

	// Field (1) 'FinalizedRoot'
	if size := len(s.FinalizedRoot); size != 32 {
		err = ssz.ErrBytesLengthFn("--.FinalizedRoot", size, 32)
		return
	}
	hh.PutBytes(s.FinalizedRoot)

	// Field (2) 'FinalizedEpoch'
	hh.PutUint64(s.FinalizedEpoch)

	// Field (3) 'HeadRoot'
	if size := len(s.HeadRoot); size != 32 {
		err = ssz.ErrBytesLengthFn("--.HeadRoot", size, 32)
		return
	}
	hh.PutBytes(s.HeadRoot)

	// Field (4) 'HeadSlot'
	hh.PutUint64(s.HeadSlot)

	if ssz.EnableVectorizedHTR {
		hh.MerkleizeVectorizedHTR(indx)
	} else {
		hh.Merkleize(indx)
	}
	return
}

// MarshalSSZ ssz marshals the MetaData object
func (m *MetaData) MarshalSSZ() ([]byte, error) {
	return ssz.MarshalSSZ(m)
}

// MarshalSSZTo ssz marshals the MetaData object to a target array
func (m *MetaData) MarshalSSZTo(buf []byte) (dst []byte, err error) {
	dst = buf
	offset := int(12)

	// Field (0) 'SeqNumber'
	dst = ssz.MarshalUint64(dst, m.SeqNumber)

	// Offset (1) 'AttNets'
	dst = ssz.WriteOffset(dst, offset)
	offset += len(m.AttNets)

	// Field (1) 'AttNets'
	if size := len(m.AttNets); size > 64 {
		err = ssz.ErrBytesLengthFn("--.AttNets", size, 64)
		return
	}
	dst = append(dst, m.AttNets...)

	return
}

// UnmarshalSSZ ssz unmarshals the MetaData object
func (m *MetaData) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size < 12 {
		return ssz.ErrSize
	}

	tail := buf
	var o1 uint64

	// Field (0) 'SeqNumber'
	m.SeqNumber = ssz.UnmarshallUint64(buf[0:8])

	// Offset (1) 'AttNets'
	if o1 = ssz.ReadOffset(buf[8:12]); o1 > size {
		return ssz.ErrOffset
	}

	if o1 < 12 {
		return ssz.ErrInvalidVariableOffset
	}

	// Field (1) 'AttNets'
	{
		buf = tail[o1:]
		if err = ssz.ValidateBitlist(buf, 64); err != nil {
			return err
		}
		if cap(m.AttNets) == 0 {
			m.AttNets = make([]byte, 0, len(buf))
		}
		m.AttNets = append(m.AttNets, buf...)
	}
	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the MetaData object
func (m *MetaData) SizeSSZ() (size int) {
	size = 12

	// Field (1) 'AttNets'
	size += len(m.AttNets)

	return
}

// HashTreeRoot ssz hashes the MetaData object
func (m *MetaData) HashTreeRoot() ([32]byte, error) {
	return ssz.HashWithDefaultHasher(m)
}

// HashTreeRootWith ssz hashes the MetaData object with a hasher
func (m *MetaData) HashTreeRootWith(hh *ssz.Hasher) (err error) {
	indx := hh.Index()

	// Field (0) 'SeqNumber'
	hh.PutUint64(m.SeqNumber)

	// Field (1) 'AttNets'
	if len(m.AttNets) == 0 {
		err = ssz.ErrEmptyBitlist
		return
	}
	hh.PutBitlist(m.AttNets, 64)

	if ssz.EnableVectorizedHTR {
		hh.MerkleizeVectorizedHTR(indx)
	} else {
		hh.Merkleize(indx)
	}
	return
}

// MarshalSSZ ssz marshals the BlocksByRangeRequest object
func (b *BlocksByRangeRequest) MarshalSSZ() ([]byte, error) {
	return ssz.MarshalSSZ(b)
}

// MarshalSSZTo ssz marshals the BlocksByRangeRequest object to a target array
func (b *BlocksByRangeRequest) MarshalSSZTo(buf []byte) (dst []byte, err error) {
	dst = buf

	// Field (0) 'StartSlot'
	dst = ssz.MarshalUint64(dst, b.StartSlot)

	// Field (1) 'Count'
	dst = ssz.MarshalUint64(dst, b.Count)

	// Field (2) 'Step'
	dst = ssz.MarshalUint64(dst, b.Step)

	return
}

// UnmarshalSSZ ssz unmarshals the BlocksByRangeRequest object
func (b *BlocksByRangeRequest) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size != 24 {
		return ssz.ErrSize
	}

	// Field (0) 'StartSlot'
	b.StartSlot = ssz.UnmarshallUint64(buf[0:8])

	// Field (1) 'Count'
	b.Count = ssz.UnmarshallUint64(buf[8:16])

	// Field (2) 'Step'
	b.Step = ssz.UnmarshallUint64(buf[16:24])

	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the BlocksByRangeRequest object
func (b *BlocksByRangeRequest) SizeSSZ() (size int) {
	size = 24
	return
}

// HashTreeRoot ssz hashes the BlocksByRangeRequest object
func (b *BlocksByRangeRequest) HashTreeRoot() ([32]byte, error) {
	return ssz.HashWithDefaultHasher(b)
}

// HashTreeRootWith ssz hashes the BlocksByRangeRequest object with a hasher
func (b *BlocksByRangeRequest) HashTreeRootWith(hh *ssz.Hasher) (err error) {
	indx := hh.Index()

	// Field (0) 'StartSlot'
	hh.PutUint64(b.StartSlot)

	// Field (1) 'Count'
	hh.PutUint64(b.Count)

	// Field (2) 'Step'
	hh.PutUint64(b.Step)

	if ssz.EnableVectorizedHTR {
		hh.MerkleizeVectorizedHTR(indx)
	} else {
		hh.Merkleize(indx)
	}
	return
}
