// Code generated by bpf2go; DO NOT EDIT.

package sockops

import (
	"bytes"
	_ "embed"
	"fmt"
	"io"

	"github.com/cilium/ebpf"
)

type bpf_redirAddr2Tuple struct {
	Ip4  uint32
	Port uint32
}

type bpf_redirSocket4Tuple struct {
	Local  bpf_redirAddr2Tuple
	Remote bpf_redirAddr2Tuple
}

// loadBpf_redir returns the embedded CollectionSpec for bpf_redir.
func loadBpf_redir() (*ebpf.CollectionSpec, error) {
	reader := bytes.NewReader(_Bpf_redirBytes)
	spec, err := ebpf.LoadCollectionSpecFromReader(reader)
	if err != nil {
		return nil, fmt.Errorf("can't load bpf_redir: %w", err)
	}

	return spec, err
}

// loadBpf_redirObjects loads bpf_redir and converts it into a struct.
//
// The following types are suitable as obj argument:
//
//     *bpf_redirObjects
//     *bpf_redirPrograms
//     *bpf_redirMaps
//
// See ebpf.CollectionSpec.LoadAndAssign documentation for details.
func loadBpf_redirObjects(obj interface{}, opts *ebpf.CollectionOptions) error {
	spec, err := loadBpf_redir()
	if err != nil {
		return err
	}

	return spec.LoadAndAssign(obj, opts)
}

// bpf_redirSpecs contains maps and programs before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type bpf_redirSpecs struct {
	bpf_redirProgramSpecs
	bpf_redirMapSpecs
}

// bpf_redirSpecs contains programs before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type bpf_redirProgramSpecs struct {
	BpfRedirProxy *ebpf.ProgramSpec `ebpf:"bpf_redir_proxy"`
}

// bpf_redirMapSpecs contains maps before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type bpf_redirMapSpecs struct {
	DebugMap       *ebpf.MapSpec `ebpf:"debug_map"`
	LocalIpMap     *ebpf.MapSpec `ebpf:"local_ip_map"`
	MapActiveEstab *ebpf.MapSpec `ebpf:"map_active_estab"`
	MapProxy       *ebpf.MapSpec `ebpf:"map_proxy"`
	MapRedir       *ebpf.MapSpec `ebpf:"map_redir"`
}

// bpf_redirObjects contains all objects after they have been loaded into the kernel.
//
// It can be passed to loadBpf_redirObjects or ebpf.CollectionSpec.LoadAndAssign.
type bpf_redirObjects struct {
	bpf_redirPrograms
	bpf_redirMaps
}

func (o *bpf_redirObjects) Close() error {
	return _Bpf_redirClose(
		&o.bpf_redirPrograms,
		&o.bpf_redirMaps,
	)
}

// bpf_redirMaps contains all maps after they have been loaded into the kernel.
//
// It can be passed to loadBpf_redirObjects or ebpf.CollectionSpec.LoadAndAssign.
type bpf_redirMaps struct {
	DebugMap       *ebpf.Map `ebpf:"debug_map"`
	LocalIpMap     *ebpf.Map `ebpf:"local_ip_map"`
	MapActiveEstab *ebpf.Map `ebpf:"map_active_estab"`
	MapProxy       *ebpf.Map `ebpf:"map_proxy"`
	MapRedir       *ebpf.Map `ebpf:"map_redir"`
}

func (m *bpf_redirMaps) Close() error {
	return _Bpf_redirClose(
		m.DebugMap,
		m.LocalIpMap,
		m.MapActiveEstab,
		m.MapProxy,
		m.MapRedir,
	)
}

// bpf_redirPrograms contains all programs after they have been loaded into the kernel.
//
// It can be passed to loadBpf_redirObjects or ebpf.CollectionSpec.LoadAndAssign.
type bpf_redirPrograms struct {
	BpfRedirProxy *ebpf.Program `ebpf:"bpf_redir_proxy"`
}

func (p *bpf_redirPrograms) Close() error {
	return _Bpf_redirClose(
		p.BpfRedirProxy,
	)
}

func _Bpf_redirClose(closers ...io.Closer) error {
	for _, closer := range closers {
		if err := closer.Close(); err != nil {
			return err
		}
	}
	return nil
}

// Do not access this directly.
//go:embed bpf_redir_bpf.o
var _Bpf_redirBytes []byte