package main

import (
	"fmt"
	"io"
	"math"
	"sort"

	cid "github.com/ipfs/go-cid"
	cbg "github.com/whyrusleeping/cbor-gen"
	xerrors "golang.org/x/xerrors"
)

type AMTRoot struct {
	BitWidth uint64
	Height   uint64
	Count    uint64
	AMTNode  AMTNode
}

type AMTNode struct {
	Bmap   []byte
	Links  []cid.Cid
	Values []*cbg.Deferred
}

// Code generated by github.com/whyrusleeping/cbor-gen. DO NOT EDIT.

var _ = xerrors.Errorf
var _ = cid.Undef
var _ = math.E
var _ = sort.Sort

var lengthBufAMTRoot = []byte{132}

func (t *AMTRoot) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}

	cw := cbg.NewCborWriter(w)

	if _, err := cw.Write(lengthBufAMTRoot); err != nil {
		return err
	}

	// t.BitWidth (uint64) (uint64)

	if err := cw.WriteMajorTypeHeader(cbg.MajUnsignedInt, uint64(t.BitWidth)); err != nil {
		return err
	}

	// t.Height (uint64) (uint64)

	if err := cw.WriteMajorTypeHeader(cbg.MajUnsignedInt, uint64(t.Height)); err != nil {
		return err
	}

	// t.Count (uint64) (uint64)

	if err := cw.WriteMajorTypeHeader(cbg.MajUnsignedInt, uint64(t.Count)); err != nil {
		return err
	}

	// t.AMTNode (internal.AMTNode) (struct)
	if err := t.AMTNode.MarshalCBOR(cw); err != nil {
		return err
	}
	return nil
}

func (t *AMTRoot) UnmarshalCBOR(r io.Reader) (err error) {
	*t = AMTRoot{}

	cr := cbg.NewCborReader(r)

	maj, extra, err := cr.ReadHeader()
	if err != nil {
		return err
	}
	defer func() {
		if err == io.EOF {
			err = io.ErrUnexpectedEOF
		}
	}()

	if maj != cbg.MajArray {
		return fmt.Errorf("cbor input should be of type array")
	}

	if extra != 4 {
		return fmt.Errorf("cbor input had wrong number of fields")
	}

	// t.BitWidth (uint64) (uint64)

	{

		maj, extra, err = cr.ReadHeader()
		if err != nil {
			return err
		}
		if maj != cbg.MajUnsignedInt {
			return fmt.Errorf("wrong type for uint64 field")
		}
		t.BitWidth = uint64(extra)

	}
	// t.Height (uint64) (uint64)

	{

		maj, extra, err = cr.ReadHeader()
		if err != nil {
			return err
		}
		if maj != cbg.MajUnsignedInt {
			return fmt.Errorf("wrong type for uint64 field")
		}
		t.Height = uint64(extra)

	}
	// t.Count (uint64) (uint64)

	{

		maj, extra, err = cr.ReadHeader()
		if err != nil {
			return err
		}
		if maj != cbg.MajUnsignedInt {
			return fmt.Errorf("wrong type for uint64 field")
		}
		t.Count = uint64(extra)

	}
	// t.AMTNode (internal.AMTNode) (struct)

	{

		if err := t.AMTNode.UnmarshalCBOR(cr); err != nil {
			return xerrors.Errorf("unmarshaling t.AMTNode: %w", err)
		}

	}
	return nil
}

var lengthBufAMTNode = []byte{131}

func (t *AMTNode) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}

	cw := cbg.NewCborWriter(w)

	if _, err := cw.Write(lengthBufAMTNode); err != nil {
		return err
	}

	// t.Bmap ([]uint8) (slice)
	if len(t.Bmap) > cbg.ByteArrayMaxLen {
		return xerrors.Errorf("Byte array in field t.Bmap was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajByteString, uint64(len(t.Bmap))); err != nil {
		return err
	}

	if _, err := cw.Write(t.Bmap[:]); err != nil {
		return err
	}

	// t.Links ([]cid.Cid) (slice)
	if len(t.Links) > cbg.MaxLength {
		return xerrors.Errorf("Slice value in field t.Links was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajArray, uint64(len(t.Links))); err != nil {
		return err
	}
	for _, v := range t.Links {
		if err := cbg.WriteCid(w, v); err != nil {
			return xerrors.Errorf("failed writing cid field t.Links: %w", err)
		}
	}

	// t.Values ([]*typegen.Deferred) (slice)
	if len(t.Values) > cbg.MaxLength {
		return xerrors.Errorf("Slice value in field t.Values was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajArray, uint64(len(t.Values))); err != nil {
		return err
	}
	for _, v := range t.Values {
		if err := v.MarshalCBOR(cw); err != nil {
			return err
		}
	}
	return nil
}

func (t *AMTNode) UnmarshalCBOR(r io.Reader) (err error) {
	*t = AMTNode{}

	cr := cbg.NewCborReader(r)

	maj, extra, err := cr.ReadHeader()
	if err != nil {
		return err
	}
	defer func() {
		if err == io.EOF {
			err = io.ErrUnexpectedEOF
		}
	}()

	if maj != cbg.MajArray {
		return fmt.Errorf("cbor input should be of type array")
	}

	if extra != 3 {
		return fmt.Errorf("cbor input had wrong number of fields")
	}

	// t.Bmap ([]uint8) (slice)

	maj, extra, err = cr.ReadHeader()
	if err != nil {
		return err
	}

	if extra > cbg.ByteArrayMaxLen {
		return fmt.Errorf("t.Bmap: byte array too large (%d)", extra)
	}
	if maj != cbg.MajByteString {
		return fmt.Errorf("expected byte array")
	}

	if extra > 0 {
		t.Bmap = make([]uint8, extra)
	}

	if _, err := io.ReadFull(cr, t.Bmap[:]); err != nil {
		return err
	}
	// t.Links ([]cid.Cid) (slice)

	maj, extra, err = cr.ReadHeader()
	if err != nil {
		return err
	}

	if extra > cbg.MaxLength {
		return fmt.Errorf("t.Links: array too large (%d)", extra)
	}

	if maj != cbg.MajArray {
		return fmt.Errorf("expected cbor array")
	}

	if extra > 0 {
		t.Links = make([]cid.Cid, extra)
	}

	for i := 0; i < int(extra); i++ {

		c, err := cbg.ReadCid(cr)
		if err != nil {
			return xerrors.Errorf("reading cid field t.Links failed: %w", err)
		}
		t.Links[i] = c
	}

	// t.Values ([]*typegen.Deferred) (slice)

	maj, extra, err = cr.ReadHeader()
	if err != nil {
		return err
	}

	if extra > cbg.MaxLength {
		return fmt.Errorf("t.Values: array too large (%d)", extra)
	}

	if maj != cbg.MajArray {
		return fmt.Errorf("expected cbor array")
	}

	if extra > 0 {
		t.Values = make([]*cbg.Deferred, extra)
	}

	for i := 0; i < int(extra); i++ {

		var v cbg.Deferred
		if err := v.UnmarshalCBOR(cr); err != nil {
			return err
		}

		t.Values[i] = &v
	}

	return nil
}
