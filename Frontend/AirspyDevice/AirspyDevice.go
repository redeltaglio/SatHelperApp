/* ----------------------------------------------------------------------------
 * This file was automatically generated by SWIG (http://www.swig.org).
 * Version 3.0.12
 *
 * This file is not intended to be easily readable and contains a number of
 * coding conventions designed to improve portability and efficiency. Do not make
 * changes to this file unless you know what you are doing--modify the SWIG
 * interface file instead.
 * ----------------------------------------------------------------------------- */

// source: Frontend/AirspyDevice/AirspyDevice.i

package AirspyDevice

/*
#define intgo swig_intgo
typedef void *swig_voidp;

#include <stdint.h>


typedef long long intgo;
typedef unsigned long long uintgo;



typedef struct { char *p; intgo n; } _gostring_;
typedef struct { void* array; intgo len; intgo cap; } _goslice_;



#cgo CXXFLAGS: -std=c++11
#cgo LDFLAGS: -lairspy -lusb-1.0

typedef long long swig_type_1;
typedef long long swig_type_2;
typedef long long swig_type_3;
typedef long long swig_type_4;
typedef long long swig_type_5;
typedef long long swig_type_6;
typedef long long swig_type_7;
typedef long long swig_type_8;
typedef long long swig_type_9;
typedef long long swig_type_10;
typedef long long swig_type_11;
typedef long long swig_type_12;
typedef long long swig_type_13;
typedef long long swig_type_14;
typedef long long swig_type_15;
typedef long long swig_type_16;
typedef _gostring_ swig_type_17;
extern void _wrap_Swig_free_AirspyDevice_1c209d1ee771b908(uintptr_t arg1);
extern uintptr_t _wrap_Swig_malloc_AirspyDevice_1c209d1ee771b908(swig_intgo arg1);
extern uintptr_t _wrap__swig_NewDirectorAirspyDeviceCallbackAirspyDeviceCallback_AirspyDevice_1c209d1ee771b908(int);
extern void _wrap__swig_DirectorAirspyDeviceCallback_upcall_CbFloatIQ_AirspyDevice_1c209d1ee771b908(uintptr_t, uintptr_t data, swig_intgo length);
extern void _wrap__swig_DirectorAirspyDeviceCallback_upcall_CbS16IQ_AirspyDevice_1c209d1ee771b908(uintptr_t, uintptr_t data, swig_intgo length);
extern void _wrap__swig_DirectorAirspyDeviceCallback_upcall_CbS8IQ_AirspyDevice_1c209d1ee771b908(uintptr_t, uintptr_t data, swig_intgo length);
extern void _wrap_DeleteDirectorAirspyDeviceCallback_AirspyDevice_1c209d1ee771b908(uintptr_t arg1);
extern void _wrap_AirspyDeviceCallback_cbFloatIQ_AirspyDevice_1c209d1ee771b908(uintptr_t arg1, uintptr_t arg2, swig_intgo arg3);
extern void _wrap_AirspyDeviceCallback_cbS16IQ_AirspyDevice_1c209d1ee771b908(uintptr_t arg1, uintptr_t arg2, swig_intgo arg3);
extern void _wrap_AirspyDeviceCallback_cbS8IQ_AirspyDevice_1c209d1ee771b908(uintptr_t arg1, uintptr_t arg2, swig_intgo arg3);
extern void _wrap_delete_AirspyDeviceCallback_AirspyDevice_1c209d1ee771b908(uintptr_t arg1);
extern uintptr_t _wrap_new_AirspyDeviceCallback_AirspyDevice_1c209d1ee771b908(void);
extern uintptr_t _wrap_new_Vector32u__SWIG_0_AirspyDevice_1c209d1ee771b908(void);
extern uintptr_t _wrap_new_Vector32u__SWIG_1_AirspyDevice_1c209d1ee771b908(swig_type_1 arg1);
extern swig_type_2 _wrap_Vector32u_size_AirspyDevice_1c209d1ee771b908(uintptr_t arg1);
extern swig_type_3 _wrap_Vector32u_capacity_AirspyDevice_1c209d1ee771b908(uintptr_t arg1);
extern void _wrap_Vector32u_reserve_AirspyDevice_1c209d1ee771b908(uintptr_t arg1, swig_type_4 arg2);
extern _Bool _wrap_Vector32u_isEmpty_AirspyDevice_1c209d1ee771b908(uintptr_t arg1);
extern void _wrap_Vector32u_clear_AirspyDevice_1c209d1ee771b908(uintptr_t arg1);
extern void _wrap_Vector32u_add_AirspyDevice_1c209d1ee771b908(uintptr_t arg1, swig_intgo arg2);
extern swig_intgo _wrap_Vector32u_get_AirspyDevice_1c209d1ee771b908(uintptr_t arg1, swig_intgo arg2);
extern void _wrap_Vector32u_set_AirspyDevice_1c209d1ee771b908(uintptr_t arg1, swig_intgo arg2, swig_intgo arg3);
extern void _wrap_delete_Vector32u_AirspyDevice_1c209d1ee771b908(uintptr_t arg1);
extern uintptr_t _wrap_new_Vector32f__SWIG_0_AirspyDevice_1c209d1ee771b908(void);
extern uintptr_t _wrap_new_Vector32f__SWIG_1_AirspyDevice_1c209d1ee771b908(swig_type_5 arg1);
extern swig_type_6 _wrap_Vector32f_size_AirspyDevice_1c209d1ee771b908(uintptr_t arg1);
extern swig_type_7 _wrap_Vector32f_capacity_AirspyDevice_1c209d1ee771b908(uintptr_t arg1);
extern void _wrap_Vector32f_reserve_AirspyDevice_1c209d1ee771b908(uintptr_t arg1, swig_type_8 arg2);
extern _Bool _wrap_Vector32f_isEmpty_AirspyDevice_1c209d1ee771b908(uintptr_t arg1);
extern void _wrap_Vector32f_clear_AirspyDevice_1c209d1ee771b908(uintptr_t arg1);
extern void _wrap_Vector32f_add_AirspyDevice_1c209d1ee771b908(uintptr_t arg1, float arg2);
extern float _wrap_Vector32f_get_AirspyDevice_1c209d1ee771b908(uintptr_t arg1, swig_intgo arg2);
extern void _wrap_Vector32f_set_AirspyDevice_1c209d1ee771b908(uintptr_t arg1, swig_intgo arg2, float arg3);
extern void _wrap_delete_Vector32f_AirspyDevice_1c209d1ee771b908(uintptr_t arg1);
extern uintptr_t _wrap_new_Vector16i__SWIG_0_AirspyDevice_1c209d1ee771b908(void);
extern uintptr_t _wrap_new_Vector16i__SWIG_1_AirspyDevice_1c209d1ee771b908(swig_type_9 arg1);
extern swig_type_10 _wrap_Vector16i_size_AirspyDevice_1c209d1ee771b908(uintptr_t arg1);
extern swig_type_11 _wrap_Vector16i_capacity_AirspyDevice_1c209d1ee771b908(uintptr_t arg1);
extern void _wrap_Vector16i_reserve_AirspyDevice_1c209d1ee771b908(uintptr_t arg1, swig_type_12 arg2);
extern _Bool _wrap_Vector16i_isEmpty_AirspyDevice_1c209d1ee771b908(uintptr_t arg1);
extern void _wrap_Vector16i_clear_AirspyDevice_1c209d1ee771b908(uintptr_t arg1);
extern void _wrap_Vector16i_add_AirspyDevice_1c209d1ee771b908(uintptr_t arg1, short arg2);
extern short _wrap_Vector16i_get_AirspyDevice_1c209d1ee771b908(uintptr_t arg1, swig_intgo arg2);
extern void _wrap_Vector16i_set_AirspyDevice_1c209d1ee771b908(uintptr_t arg1, swig_intgo arg2, short arg3);
extern void _wrap_delete_Vector16i_AirspyDevice_1c209d1ee771b908(uintptr_t arg1);
extern uintptr_t _wrap_new_Vector8i__SWIG_0_AirspyDevice_1c209d1ee771b908(void);
extern uintptr_t _wrap_new_Vector8i__SWIG_1_AirspyDevice_1c209d1ee771b908(swig_type_13 arg1);
extern swig_type_14 _wrap_Vector8i_size_AirspyDevice_1c209d1ee771b908(uintptr_t arg1);
extern swig_type_15 _wrap_Vector8i_capacity_AirspyDevice_1c209d1ee771b908(uintptr_t arg1);
extern void _wrap_Vector8i_reserve_AirspyDevice_1c209d1ee771b908(uintptr_t arg1, swig_type_16 arg2);
extern _Bool _wrap_Vector8i_isEmpty_AirspyDevice_1c209d1ee771b908(uintptr_t arg1);
extern void _wrap_Vector8i_clear_AirspyDevice_1c209d1ee771b908(uintptr_t arg1);
extern void _wrap_Vector8i_add_AirspyDevice_1c209d1ee771b908(uintptr_t arg1, char arg2);
extern char _wrap_Vector8i_get_AirspyDevice_1c209d1ee771b908(uintptr_t arg1, swig_intgo arg2);
extern void _wrap_Vector8i_set_AirspyDevice_1c209d1ee771b908(uintptr_t arg1, swig_intgo arg2, char arg3);
extern void _wrap_delete_Vector8i_AirspyDevice_1c209d1ee771b908(uintptr_t arg1);
extern uintptr_t _wrap_new_AirspyDevice_AirspyDevice_1c209d1ee771b908(void);
extern void _wrap_delete_AirspyDevice_AirspyDevice_1c209d1ee771b908(uintptr_t arg1);
extern void _wrap_AirspyDevice_Initialize_AirspyDevice_1c209d1ee771b908(void);
extern void _wrap_AirspyDevice_DeInitialize_AirspyDevice_1c209d1ee771b908(void);
extern swig_intgo _wrap_AirspyDevice_SetSampleRate_AirspyDevice_1c209d1ee771b908(uintptr_t arg1, swig_intgo arg2);
extern swig_intgo _wrap_AirspyDevice_SetCenterFrequency_AirspyDevice_1c209d1ee771b908(uintptr_t arg1, swig_intgo arg2);
extern uintptr_t _wrap_AirspyDevice_GetAvailableSampleRates_AirspyDevice_1c209d1ee771b908(uintptr_t arg1);
extern void _wrap_AirspyDevice_Start_AirspyDevice_1c209d1ee771b908(uintptr_t arg1);
extern void _wrap_AirspyDevice_Stop_AirspyDevice_1c209d1ee771b908(uintptr_t arg1);
extern void _wrap_AirspyDevice_SetAGC_AirspyDevice_1c209d1ee771b908(uintptr_t arg1, _Bool arg2);
extern _Bool _wrap_AirspyDevice_Init_AirspyDevice_1c209d1ee771b908(uintptr_t arg1);
extern void _wrap_AirspyDevice_Destroy_AirspyDevice_1c209d1ee771b908(uintptr_t arg1);
extern void _wrap_AirspyDevice_SetLNAGain_AirspyDevice_1c209d1ee771b908(uintptr_t arg1, char arg2);
extern void _wrap_AirspyDevice_SetVGAGain_AirspyDevice_1c209d1ee771b908(uintptr_t arg1, char arg2);
extern void _wrap_AirspyDevice_SetMixerGain_AirspyDevice_1c209d1ee771b908(uintptr_t arg1, char arg2);
extern void _wrap_AirspyDevice_SetBiasT_AirspyDevice_1c209d1ee771b908(uintptr_t arg1, char arg2);
extern swig_intgo _wrap_AirspyDevice_GetCenterFrequency_AirspyDevice_1c209d1ee771b908(uintptr_t arg1);
extern swig_type_17 _wrap_AirspyDevice_GetName_AirspyDevice_1c209d1ee771b908(uintptr_t arg1);
extern swig_intgo _wrap_AirspyDevice_GetSampleRate_AirspyDevice_1c209d1ee771b908(uintptr_t arg1);
extern void _wrap_AirspyDevice_SetSamplesAvailableCallback_AirspyDevice_1c209d1ee771b908(uintptr_t arg1, uintptr_t arg2);
#undef intgo
*/
import "C"

import "unsafe"
import _ "runtime/cgo"
import "sync"


type _ unsafe.Pointer



var Swig_escape_always_false bool
var Swig_escape_val interface{}


type _swig_fnptr *byte
type _swig_memberptr *byte


type _ sync.Mutex


type swig_gostring struct { p uintptr; n int }
func swigCopyString(s string) string {
  p := *(*swig_gostring)(unsafe.Pointer(&s))
  r := string((*[0x7fffffff]byte)(unsafe.Pointer(p.p))[:p.n])
  Swig_free(p.p)
  return r
}

func Swig_free(arg1 uintptr) {
	_swig_i_0 := arg1
	C._wrap_Swig_free_AirspyDevice_1c209d1ee771b908(C.uintptr_t(_swig_i_0))
}

func Swig_malloc(arg1 int) (_swig_ret uintptr) {
	var swig_r uintptr
	_swig_i_0 := arg1
	swig_r = (uintptr)(C._wrap_Swig_malloc_AirspyDevice_1c209d1ee771b908(C.swig_intgo(_swig_i_0)))
	return swig_r
}

const FRONTEND_SAMPLETYPE_FLOATIQ int = 0
const FRONTEND_SAMPLETYPE_S16IQ int = 1
const FRONTEND_SAMPLETYPE_S8IQ int = 2
type _swig_DirectorAirspyDeviceCallback struct {
	SwigcptrAirspyDeviceCallback
	v interface{}
}

func (p *_swig_DirectorAirspyDeviceCallback) Swigcptr() uintptr {
	return p.SwigcptrAirspyDeviceCallback.Swigcptr()
}

func (p *_swig_DirectorAirspyDeviceCallback) SwigIsAirspyDeviceCallback() {
}

func (p *_swig_DirectorAirspyDeviceCallback) DirectorInterface() interface{} {
	return p.v
}

func NewDirectorAirspyDeviceCallback(v interface{}) AirspyDeviceCallback {
	p := &_swig_DirectorAirspyDeviceCallback{0, v}
	p.SwigcptrAirspyDeviceCallback = SwigcptrAirspyDeviceCallback(C._wrap__swig_NewDirectorAirspyDeviceCallbackAirspyDeviceCallback_AirspyDevice_1c209d1ee771b908(C.int(swigDirectorAdd(p))))
	return p
}

type _swig_DirectorInterfaceAirspyDeviceCallbackCbFloatIQ interface {
	CbFloatIQ(uintptr, int)
}

func (swig_p *_swig_DirectorAirspyDeviceCallback) CbFloatIQ(data uintptr, length int) {
	if swig_g, swig_ok := swig_p.v.(_swig_DirectorInterfaceAirspyDeviceCallbackCbFloatIQ); swig_ok {
		swig_g.CbFloatIQ(data, length)
		return
	}
	_swig_i_0 := data
	_swig_i_1 := length
	C._wrap__swig_DirectorAirspyDeviceCallback_upcall_CbFloatIQ_AirspyDevice_1c209d1ee771b908(C.uintptr_t(swig_p.SwigcptrAirspyDeviceCallback), C.uintptr_t(_swig_i_0), C.swig_intgo(_swig_i_1))
}

func DirectorAirspyDeviceCallbackCbFloatIQ(p AirspyDeviceCallback, arg2 uintptr, arg3 int) {
	_swig_i_0 := arg2
	_swig_i_1 := arg3
	C._wrap__swig_DirectorAirspyDeviceCallback_upcall_CbFloatIQ_AirspyDevice_1c209d1ee771b908(C.uintptr_t(p.(*_swig_DirectorAirspyDeviceCallback).SwigcptrAirspyDeviceCallback), C.uintptr_t(_swig_i_0), C.swig_intgo(_swig_i_1))
}

//export Swig_DirectorAirspyDeviceCallback_callback_cbFloatIQ_AirspyDevice_1c209d1ee771b908
func Swig_DirectorAirspyDeviceCallback_callback_cbFloatIQ_AirspyDevice_1c209d1ee771b908(swig_c int, arg2 uintptr, arg3 int) {
	swig_p := swigDirectorLookup(swig_c).(*_swig_DirectorAirspyDeviceCallback)
	swig_p.CbFloatIQ(arg2, arg3)
}

type _swig_DirectorInterfaceAirspyDeviceCallbackCbS16IQ interface {
	CbS16IQ(uintptr, int)
}

func (swig_p *_swig_DirectorAirspyDeviceCallback) CbS16IQ(data uintptr, length int) {
	if swig_g, swig_ok := swig_p.v.(_swig_DirectorInterfaceAirspyDeviceCallbackCbS16IQ); swig_ok {
		swig_g.CbS16IQ(data, length)
		return
	}
	_swig_i_0 := data
	_swig_i_1 := length
	C._wrap__swig_DirectorAirspyDeviceCallback_upcall_CbS16IQ_AirspyDevice_1c209d1ee771b908(C.uintptr_t(swig_p.SwigcptrAirspyDeviceCallback), C.uintptr_t(_swig_i_0), C.swig_intgo(_swig_i_1))
}

func DirectorAirspyDeviceCallbackCbS16IQ(p AirspyDeviceCallback, arg2 uintptr, arg3 int) {
	_swig_i_0 := arg2
	_swig_i_1 := arg3
	C._wrap__swig_DirectorAirspyDeviceCallback_upcall_CbS16IQ_AirspyDevice_1c209d1ee771b908(C.uintptr_t(p.(*_swig_DirectorAirspyDeviceCallback).SwigcptrAirspyDeviceCallback), C.uintptr_t(_swig_i_0), C.swig_intgo(_swig_i_1))
}

//export Swig_DirectorAirspyDeviceCallback_callback_cbS16IQ_AirspyDevice_1c209d1ee771b908
func Swig_DirectorAirspyDeviceCallback_callback_cbS16IQ_AirspyDevice_1c209d1ee771b908(swig_c int, arg2 uintptr, arg3 int) {
	swig_p := swigDirectorLookup(swig_c).(*_swig_DirectorAirspyDeviceCallback)
	swig_p.CbS16IQ(arg2, arg3)
}

type _swig_DirectorInterfaceAirspyDeviceCallbackCbS8IQ interface {
	CbS8IQ(uintptr, int)
}

func (swig_p *_swig_DirectorAirspyDeviceCallback) CbS8IQ(data uintptr, length int) {
	if swig_g, swig_ok := swig_p.v.(_swig_DirectorInterfaceAirspyDeviceCallbackCbS8IQ); swig_ok {
		swig_g.CbS8IQ(data, length)
		return
	}
	_swig_i_0 := data
	_swig_i_1 := length
	C._wrap__swig_DirectorAirspyDeviceCallback_upcall_CbS8IQ_AirspyDevice_1c209d1ee771b908(C.uintptr_t(swig_p.SwigcptrAirspyDeviceCallback), C.uintptr_t(_swig_i_0), C.swig_intgo(_swig_i_1))
}

func DirectorAirspyDeviceCallbackCbS8IQ(p AirspyDeviceCallback, arg2 uintptr, arg3 int) {
	_swig_i_0 := arg2
	_swig_i_1 := arg3
	C._wrap__swig_DirectorAirspyDeviceCallback_upcall_CbS8IQ_AirspyDevice_1c209d1ee771b908(C.uintptr_t(p.(*_swig_DirectorAirspyDeviceCallback).SwigcptrAirspyDeviceCallback), C.uintptr_t(_swig_i_0), C.swig_intgo(_swig_i_1))
}

//export Swig_DirectorAirspyDeviceCallback_callback_cbS8IQ_AirspyDevice_1c209d1ee771b908
func Swig_DirectorAirspyDeviceCallback_callback_cbS8IQ_AirspyDevice_1c209d1ee771b908(swig_c int, arg2 uintptr, arg3 int) {
	swig_p := swigDirectorLookup(swig_c).(*_swig_DirectorAirspyDeviceCallback)
	swig_p.CbS8IQ(arg2, arg3)
}

func DeleteDirectorAirspyDeviceCallback(arg1 AirspyDeviceCallback) {
	_swig_i_0 := arg1.Swigcptr()
	C._wrap_DeleteDirectorAirspyDeviceCallback_AirspyDevice_1c209d1ee771b908(C.uintptr_t(_swig_i_0))
}

//export Swiggo_DeleteDirector_AirspyDeviceCallback_AirspyDevice_1c209d1ee771b908
func Swiggo_DeleteDirector_AirspyDeviceCallback_AirspyDevice_1c209d1ee771b908(c int) {
	swigDirectorLookup(c).(*_swig_DirectorAirspyDeviceCallback).SwigcptrAirspyDeviceCallback = 0
	swigDirectorDelete(c)
}

type SwigcptrAirspyDeviceCallback uintptr

func (p SwigcptrAirspyDeviceCallback) Swigcptr() uintptr {
	return (uintptr)(p)
}

func (p SwigcptrAirspyDeviceCallback) SwigIsAirspyDeviceCallback() {
}

func (p SwigcptrAirspyDeviceCallback) DirectorInterface() interface{} {
	return nil
}

func (arg1 SwigcptrAirspyDeviceCallback) CbFloatIQ(arg2 uintptr, arg3 int) {
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	_swig_i_2 := arg3
	C._wrap_AirspyDeviceCallback_cbFloatIQ_AirspyDevice_1c209d1ee771b908(C.uintptr_t(_swig_i_0), C.uintptr_t(_swig_i_1), C.swig_intgo(_swig_i_2))
}

func (arg1 SwigcptrAirspyDeviceCallback) CbS16IQ(arg2 uintptr, arg3 int) {
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	_swig_i_2 := arg3
	C._wrap_AirspyDeviceCallback_cbS16IQ_AirspyDevice_1c209d1ee771b908(C.uintptr_t(_swig_i_0), C.uintptr_t(_swig_i_1), C.swig_intgo(_swig_i_2))
}

func (arg1 SwigcptrAirspyDeviceCallback) CbS8IQ(arg2 uintptr, arg3 int) {
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	_swig_i_2 := arg3
	C._wrap_AirspyDeviceCallback_cbS8IQ_AirspyDevice_1c209d1ee771b908(C.uintptr_t(_swig_i_0), C.uintptr_t(_swig_i_1), C.swig_intgo(_swig_i_2))
}

func DeleteAirspyDeviceCallback(arg1 AirspyDeviceCallback) {
	_swig_i_0 := arg1.Swigcptr()
	C._wrap_delete_AirspyDeviceCallback_AirspyDevice_1c209d1ee771b908(C.uintptr_t(_swig_i_0))
}

func NewAirspyDeviceCallback() (_swig_ret AirspyDeviceCallback) {
	var swig_r AirspyDeviceCallback
	swig_r = (AirspyDeviceCallback)(SwigcptrAirspyDeviceCallback(C._wrap_new_AirspyDeviceCallback_AirspyDevice_1c209d1ee771b908()))
	return swig_r
}

type AirspyDeviceCallback interface {
	Swigcptr() uintptr
	SwigIsAirspyDeviceCallback()
	DirectorInterface() interface{}
	CbFloatIQ(arg2 uintptr, arg3 int)
	CbS16IQ(arg2 uintptr, arg3 int)
	CbS8IQ(arg2 uintptr, arg3 int)
}

type SwigcptrVector32u uintptr

func (p SwigcptrVector32u) Swigcptr() uintptr {
	return (uintptr)(p)
}

func (p SwigcptrVector32u) SwigIsVector32u() {
}

func NewVector32u__SWIG_0() (_swig_ret Vector32u) {
	var swig_r Vector32u
	swig_r = (Vector32u)(SwigcptrVector32u(C._wrap_new_Vector32u__SWIG_0_AirspyDevice_1c209d1ee771b908()))
	return swig_r
}

func NewVector32u__SWIG_1(arg1 int64) (_swig_ret Vector32u) {
	var swig_r Vector32u
	_swig_i_0 := arg1
	swig_r = (Vector32u)(SwigcptrVector32u(C._wrap_new_Vector32u__SWIG_1_AirspyDevice_1c209d1ee771b908(C.swig_type_1(_swig_i_0))))
	return swig_r
}

func NewVector32u(a ...interface{}) Vector32u {
	argc := len(a)
	if argc == 0 {
		return NewVector32u__SWIG_0()
	}
	if argc == 1 {
		return NewVector32u__SWIG_1(a[0].(int64))
	}
	panic("No match for overloaded function call")
}

func (arg1 SwigcptrVector32u) Size() (_swig_ret int64) {
	var swig_r int64
	_swig_i_0 := arg1
	swig_r = (int64)(C._wrap_Vector32u_size_AirspyDevice_1c209d1ee771b908(C.uintptr_t(_swig_i_0)))
	return swig_r
}

func (arg1 SwigcptrVector32u) Capacity() (_swig_ret int64) {
	var swig_r int64
	_swig_i_0 := arg1
	swig_r = (int64)(C._wrap_Vector32u_capacity_AirspyDevice_1c209d1ee771b908(C.uintptr_t(_swig_i_0)))
	return swig_r
}

func (arg1 SwigcptrVector32u) Reserve(arg2 int64) {
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	C._wrap_Vector32u_reserve_AirspyDevice_1c209d1ee771b908(C.uintptr_t(_swig_i_0), C.swig_type_4(_swig_i_1))
}

func (arg1 SwigcptrVector32u) IsEmpty() (_swig_ret bool) {
	var swig_r bool
	_swig_i_0 := arg1
	swig_r = (bool)(C._wrap_Vector32u_isEmpty_AirspyDevice_1c209d1ee771b908(C.uintptr_t(_swig_i_0)))
	return swig_r
}

func (arg1 SwigcptrVector32u) Clear() {
	_swig_i_0 := arg1
	C._wrap_Vector32u_clear_AirspyDevice_1c209d1ee771b908(C.uintptr_t(_swig_i_0))
}

func (arg1 SwigcptrVector32u) Add(arg2 uint) {
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	C._wrap_Vector32u_add_AirspyDevice_1c209d1ee771b908(C.uintptr_t(_swig_i_0), C.swig_intgo(_swig_i_1))
}

func (arg1 SwigcptrVector32u) Get(arg2 int) (_swig_ret uint) {
	var swig_r uint
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	swig_r = (uint)(C._wrap_Vector32u_get_AirspyDevice_1c209d1ee771b908(C.uintptr_t(_swig_i_0), C.swig_intgo(_swig_i_1)))
	return swig_r
}

func (arg1 SwigcptrVector32u) Set(arg2 int, arg3 uint) {
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	_swig_i_2 := arg3
	C._wrap_Vector32u_set_AirspyDevice_1c209d1ee771b908(C.uintptr_t(_swig_i_0), C.swig_intgo(_swig_i_1), C.swig_intgo(_swig_i_2))
}

func DeleteVector32u(arg1 Vector32u) {
	_swig_i_0 := arg1.Swigcptr()
	C._wrap_delete_Vector32u_AirspyDevice_1c209d1ee771b908(C.uintptr_t(_swig_i_0))
}

type Vector32u interface {
	Swigcptr() uintptr
	SwigIsVector32u()
	Size() (_swig_ret int64)
	Capacity() (_swig_ret int64)
	Reserve(arg2 int64)
	IsEmpty() (_swig_ret bool)
	Clear()
	Add(arg2 uint)
	Get(arg2 int) (_swig_ret uint)
	Set(arg2 int, arg3 uint)
}

type SwigcptrVector32f uintptr

func (p SwigcptrVector32f) Swigcptr() uintptr {
	return (uintptr)(p)
}

func (p SwigcptrVector32f) SwigIsVector32f() {
}

func NewVector32f__SWIG_0() (_swig_ret Vector32f) {
	var swig_r Vector32f
	swig_r = (Vector32f)(SwigcptrVector32f(C._wrap_new_Vector32f__SWIG_0_AirspyDevice_1c209d1ee771b908()))
	return swig_r
}

func NewVector32f__SWIG_1(arg1 int64) (_swig_ret Vector32f) {
	var swig_r Vector32f
	_swig_i_0 := arg1
	swig_r = (Vector32f)(SwigcptrVector32f(C._wrap_new_Vector32f__SWIG_1_AirspyDevice_1c209d1ee771b908(C.swig_type_5(_swig_i_0))))
	return swig_r
}

func NewVector32f(a ...interface{}) Vector32f {
	argc := len(a)
	if argc == 0 {
		return NewVector32f__SWIG_0()
	}
	if argc == 1 {
		return NewVector32f__SWIG_1(a[0].(int64))
	}
	panic("No match for overloaded function call")
}

func (arg1 SwigcptrVector32f) Size() (_swig_ret int64) {
	var swig_r int64
	_swig_i_0 := arg1
	swig_r = (int64)(C._wrap_Vector32f_size_AirspyDevice_1c209d1ee771b908(C.uintptr_t(_swig_i_0)))
	return swig_r
}

func (arg1 SwigcptrVector32f) Capacity() (_swig_ret int64) {
	var swig_r int64
	_swig_i_0 := arg1
	swig_r = (int64)(C._wrap_Vector32f_capacity_AirspyDevice_1c209d1ee771b908(C.uintptr_t(_swig_i_0)))
	return swig_r
}

func (arg1 SwigcptrVector32f) Reserve(arg2 int64) {
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	C._wrap_Vector32f_reserve_AirspyDevice_1c209d1ee771b908(C.uintptr_t(_swig_i_0), C.swig_type_8(_swig_i_1))
}

func (arg1 SwigcptrVector32f) IsEmpty() (_swig_ret bool) {
	var swig_r bool
	_swig_i_0 := arg1
	swig_r = (bool)(C._wrap_Vector32f_isEmpty_AirspyDevice_1c209d1ee771b908(C.uintptr_t(_swig_i_0)))
	return swig_r
}

func (arg1 SwigcptrVector32f) Clear() {
	_swig_i_0 := arg1
	C._wrap_Vector32f_clear_AirspyDevice_1c209d1ee771b908(C.uintptr_t(_swig_i_0))
}

func (arg1 SwigcptrVector32f) Add(arg2 float32) {
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	C._wrap_Vector32f_add_AirspyDevice_1c209d1ee771b908(C.uintptr_t(_swig_i_0), C.float(_swig_i_1))
}

func (arg1 SwigcptrVector32f) Get(arg2 int) (_swig_ret float32) {
	var swig_r float32
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	swig_r = (float32)(C._wrap_Vector32f_get_AirspyDevice_1c209d1ee771b908(C.uintptr_t(_swig_i_0), C.swig_intgo(_swig_i_1)))
	return swig_r
}

func (arg1 SwigcptrVector32f) Set(arg2 int, arg3 float32) {
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	_swig_i_2 := arg3
	C._wrap_Vector32f_set_AirspyDevice_1c209d1ee771b908(C.uintptr_t(_swig_i_0), C.swig_intgo(_swig_i_1), C.float(_swig_i_2))
}

func DeleteVector32f(arg1 Vector32f) {
	_swig_i_0 := arg1.Swigcptr()
	C._wrap_delete_Vector32f_AirspyDevice_1c209d1ee771b908(C.uintptr_t(_swig_i_0))
}

type Vector32f interface {
	Swigcptr() uintptr
	SwigIsVector32f()
	Size() (_swig_ret int64)
	Capacity() (_swig_ret int64)
	Reserve(arg2 int64)
	IsEmpty() (_swig_ret bool)
	Clear()
	Add(arg2 float32)
	Get(arg2 int) (_swig_ret float32)
	Set(arg2 int, arg3 float32)
}

type SwigcptrVector16i uintptr

func (p SwigcptrVector16i) Swigcptr() uintptr {
	return (uintptr)(p)
}

func (p SwigcptrVector16i) SwigIsVector16i() {
}

func NewVector16i__SWIG_0() (_swig_ret Vector16i) {
	var swig_r Vector16i
	swig_r = (Vector16i)(SwigcptrVector16i(C._wrap_new_Vector16i__SWIG_0_AirspyDevice_1c209d1ee771b908()))
	return swig_r
}

func NewVector16i__SWIG_1(arg1 int64) (_swig_ret Vector16i) {
	var swig_r Vector16i
	_swig_i_0 := arg1
	swig_r = (Vector16i)(SwigcptrVector16i(C._wrap_new_Vector16i__SWIG_1_AirspyDevice_1c209d1ee771b908(C.swig_type_9(_swig_i_0))))
	return swig_r
}

func NewVector16i(a ...interface{}) Vector16i {
	argc := len(a)
	if argc == 0 {
		return NewVector16i__SWIG_0()
	}
	if argc == 1 {
		return NewVector16i__SWIG_1(a[0].(int64))
	}
	panic("No match for overloaded function call")
}

func (arg1 SwigcptrVector16i) Size() (_swig_ret int64) {
	var swig_r int64
	_swig_i_0 := arg1
	swig_r = (int64)(C._wrap_Vector16i_size_AirspyDevice_1c209d1ee771b908(C.uintptr_t(_swig_i_0)))
	return swig_r
}

func (arg1 SwigcptrVector16i) Capacity() (_swig_ret int64) {
	var swig_r int64
	_swig_i_0 := arg1
	swig_r = (int64)(C._wrap_Vector16i_capacity_AirspyDevice_1c209d1ee771b908(C.uintptr_t(_swig_i_0)))
	return swig_r
}

func (arg1 SwigcptrVector16i) Reserve(arg2 int64) {
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	C._wrap_Vector16i_reserve_AirspyDevice_1c209d1ee771b908(C.uintptr_t(_swig_i_0), C.swig_type_12(_swig_i_1))
}

func (arg1 SwigcptrVector16i) IsEmpty() (_swig_ret bool) {
	var swig_r bool
	_swig_i_0 := arg1
	swig_r = (bool)(C._wrap_Vector16i_isEmpty_AirspyDevice_1c209d1ee771b908(C.uintptr_t(_swig_i_0)))
	return swig_r
}

func (arg1 SwigcptrVector16i) Clear() {
	_swig_i_0 := arg1
	C._wrap_Vector16i_clear_AirspyDevice_1c209d1ee771b908(C.uintptr_t(_swig_i_0))
}

func (arg1 SwigcptrVector16i) Add(arg2 int16) {
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	C._wrap_Vector16i_add_AirspyDevice_1c209d1ee771b908(C.uintptr_t(_swig_i_0), C.short(_swig_i_1))
}

func (arg1 SwigcptrVector16i) Get(arg2 int) (_swig_ret int16) {
	var swig_r int16
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	swig_r = (int16)(C._wrap_Vector16i_get_AirspyDevice_1c209d1ee771b908(C.uintptr_t(_swig_i_0), C.swig_intgo(_swig_i_1)))
	return swig_r
}

func (arg1 SwigcptrVector16i) Set(arg2 int, arg3 int16) {
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	_swig_i_2 := arg3
	C._wrap_Vector16i_set_AirspyDevice_1c209d1ee771b908(C.uintptr_t(_swig_i_0), C.swig_intgo(_swig_i_1), C.short(_swig_i_2))
}

func DeleteVector16i(arg1 Vector16i) {
	_swig_i_0 := arg1.Swigcptr()
	C._wrap_delete_Vector16i_AirspyDevice_1c209d1ee771b908(C.uintptr_t(_swig_i_0))
}

type Vector16i interface {
	Swigcptr() uintptr
	SwigIsVector16i()
	Size() (_swig_ret int64)
	Capacity() (_swig_ret int64)
	Reserve(arg2 int64)
	IsEmpty() (_swig_ret bool)
	Clear()
	Add(arg2 int16)
	Get(arg2 int) (_swig_ret int16)
	Set(arg2 int, arg3 int16)
}

type SwigcptrVector8i uintptr

func (p SwigcptrVector8i) Swigcptr() uintptr {
	return (uintptr)(p)
}

func (p SwigcptrVector8i) SwigIsVector8i() {
}

func NewVector8i__SWIG_0() (_swig_ret Vector8i) {
	var swig_r Vector8i
	swig_r = (Vector8i)(SwigcptrVector8i(C._wrap_new_Vector8i__SWIG_0_AirspyDevice_1c209d1ee771b908()))
	return swig_r
}

func NewVector8i__SWIG_1(arg1 int64) (_swig_ret Vector8i) {
	var swig_r Vector8i
	_swig_i_0 := arg1
	swig_r = (Vector8i)(SwigcptrVector8i(C._wrap_new_Vector8i__SWIG_1_AirspyDevice_1c209d1ee771b908(C.swig_type_13(_swig_i_0))))
	return swig_r
}

func NewVector8i(a ...interface{}) Vector8i {
	argc := len(a)
	if argc == 0 {
		return NewVector8i__SWIG_0()
	}
	if argc == 1 {
		return NewVector8i__SWIG_1(a[0].(int64))
	}
	panic("No match for overloaded function call")
}

func (arg1 SwigcptrVector8i) Size() (_swig_ret int64) {
	var swig_r int64
	_swig_i_0 := arg1
	swig_r = (int64)(C._wrap_Vector8i_size_AirspyDevice_1c209d1ee771b908(C.uintptr_t(_swig_i_0)))
	return swig_r
}

func (arg1 SwigcptrVector8i) Capacity() (_swig_ret int64) {
	var swig_r int64
	_swig_i_0 := arg1
	swig_r = (int64)(C._wrap_Vector8i_capacity_AirspyDevice_1c209d1ee771b908(C.uintptr_t(_swig_i_0)))
	return swig_r
}

func (arg1 SwigcptrVector8i) Reserve(arg2 int64) {
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	C._wrap_Vector8i_reserve_AirspyDevice_1c209d1ee771b908(C.uintptr_t(_swig_i_0), C.swig_type_16(_swig_i_1))
}

func (arg1 SwigcptrVector8i) IsEmpty() (_swig_ret bool) {
	var swig_r bool
	_swig_i_0 := arg1
	swig_r = (bool)(C._wrap_Vector8i_isEmpty_AirspyDevice_1c209d1ee771b908(C.uintptr_t(_swig_i_0)))
	return swig_r
}

func (arg1 SwigcptrVector8i) Clear() {
	_swig_i_0 := arg1
	C._wrap_Vector8i_clear_AirspyDevice_1c209d1ee771b908(C.uintptr_t(_swig_i_0))
}

func (arg1 SwigcptrVector8i) Add(arg2 int8) {
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	C._wrap_Vector8i_add_AirspyDevice_1c209d1ee771b908(C.uintptr_t(_swig_i_0), C.char(_swig_i_1))
}

func (arg1 SwigcptrVector8i) Get(arg2 int) (_swig_ret int8) {
	var swig_r int8
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	swig_r = (int8)(C._wrap_Vector8i_get_AirspyDevice_1c209d1ee771b908(C.uintptr_t(_swig_i_0), C.swig_intgo(_swig_i_1)))
	return swig_r
}

func (arg1 SwigcptrVector8i) Set(arg2 int, arg3 int8) {
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	_swig_i_2 := arg3
	C._wrap_Vector8i_set_AirspyDevice_1c209d1ee771b908(C.uintptr_t(_swig_i_0), C.swig_intgo(_swig_i_1), C.char(_swig_i_2))
}

func DeleteVector8i(arg1 Vector8i) {
	_swig_i_0 := arg1.Swigcptr()
	C._wrap_delete_Vector8i_AirspyDevice_1c209d1ee771b908(C.uintptr_t(_swig_i_0))
}

type Vector8i interface {
	Swigcptr() uintptr
	SwigIsVector8i()
	Size() (_swig_ret int64)
	Capacity() (_swig_ret int64)
	Reserve(arg2 int64)
	IsEmpty() (_swig_ret bool)
	Clear()
	Add(arg2 int8)
	Get(arg2 int) (_swig_ret int8)
	Set(arg2 int, arg3 int8)
}

type SwigcptrAirspyDevice uintptr

func (p SwigcptrAirspyDevice) Swigcptr() uintptr {
	return (uintptr)(p)
}

func (p SwigcptrAirspyDevice) SwigIsAirspyDevice() {
}

func NewAirspyDevice() (_swig_ret AirspyDevice) {
	var swig_r AirspyDevice
	swig_r = (AirspyDevice)(SwigcptrAirspyDevice(C._wrap_new_AirspyDevice_AirspyDevice_1c209d1ee771b908()))
	return swig_r
}

func DeleteAirspyDevice(arg1 AirspyDevice) {
	_swig_i_0 := arg1.Swigcptr()
	C._wrap_delete_AirspyDevice_AirspyDevice_1c209d1ee771b908(C.uintptr_t(_swig_i_0))
}

func AirspyDeviceInitialize() {
	C._wrap_AirspyDevice_Initialize_AirspyDevice_1c209d1ee771b908()
}

func AirspyDeviceDeInitialize() {
	C._wrap_AirspyDevice_DeInitialize_AirspyDevice_1c209d1ee771b908()
}

func (arg1 SwigcptrAirspyDevice) SetSampleRate(arg2 uint) (_swig_ret uint) {
	var swig_r uint
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	swig_r = (uint)(C._wrap_AirspyDevice_SetSampleRate_AirspyDevice_1c209d1ee771b908(C.uintptr_t(_swig_i_0), C.swig_intgo(_swig_i_1)))
	return swig_r
}

func (arg1 SwigcptrAirspyDevice) SetCenterFrequency(arg2 uint) (_swig_ret uint) {
	var swig_r uint
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	swig_r = (uint)(C._wrap_AirspyDevice_SetCenterFrequency_AirspyDevice_1c209d1ee771b908(C.uintptr_t(_swig_i_0), C.swig_intgo(_swig_i_1)))
	return swig_r
}

func (arg1 SwigcptrAirspyDevice) GetAvailableSampleRates() (_swig_ret Vector32u) {
	var swig_r Vector32u
	_swig_i_0 := arg1
	swig_r = (Vector32u)(SwigcptrVector32u(C._wrap_AirspyDevice_GetAvailableSampleRates_AirspyDevice_1c209d1ee771b908(C.uintptr_t(_swig_i_0))))
	return swig_r
}

func (arg1 SwigcptrAirspyDevice) Start() {
	_swig_i_0 := arg1
	C._wrap_AirspyDevice_Start_AirspyDevice_1c209d1ee771b908(C.uintptr_t(_swig_i_0))
}

func (arg1 SwigcptrAirspyDevice) Stop() {
	_swig_i_0 := arg1
	C._wrap_AirspyDevice_Stop_AirspyDevice_1c209d1ee771b908(C.uintptr_t(_swig_i_0))
}

func (arg1 SwigcptrAirspyDevice) SetAGC(arg2 bool) {
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	C._wrap_AirspyDevice_SetAGC_AirspyDevice_1c209d1ee771b908(C.uintptr_t(_swig_i_0), C._Bool(_swig_i_1))
}

func (arg1 SwigcptrAirspyDevice) Init() (_swig_ret bool) {
	var swig_r bool
	_swig_i_0 := arg1
	swig_r = (bool)(C._wrap_AirspyDevice_Init_AirspyDevice_1c209d1ee771b908(C.uintptr_t(_swig_i_0)))
	return swig_r
}

func (arg1 SwigcptrAirspyDevice) Destroy() {
	_swig_i_0 := arg1
	C._wrap_AirspyDevice_Destroy_AirspyDevice_1c209d1ee771b908(C.uintptr_t(_swig_i_0))
}

func (arg1 SwigcptrAirspyDevice) SetLNAGain(arg2 byte) {
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	C._wrap_AirspyDevice_SetLNAGain_AirspyDevice_1c209d1ee771b908(C.uintptr_t(_swig_i_0), C.char(_swig_i_1))
}

func (arg1 SwigcptrAirspyDevice) SetVGAGain(arg2 byte) {
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	C._wrap_AirspyDevice_SetVGAGain_AirspyDevice_1c209d1ee771b908(C.uintptr_t(_swig_i_0), C.char(_swig_i_1))
}

func (arg1 SwigcptrAirspyDevice) SetMixerGain(arg2 byte) {
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	C._wrap_AirspyDevice_SetMixerGain_AirspyDevice_1c209d1ee771b908(C.uintptr_t(_swig_i_0), C.char(_swig_i_1))
}

func (arg1 SwigcptrAirspyDevice) SetBiasT(arg2 byte) {
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	C._wrap_AirspyDevice_SetBiasT_AirspyDevice_1c209d1ee771b908(C.uintptr_t(_swig_i_0), C.char(_swig_i_1))
}

func (arg1 SwigcptrAirspyDevice) GetCenterFrequency() (_swig_ret uint) {
	var swig_r uint
	_swig_i_0 := arg1
	swig_r = (uint)(C._wrap_AirspyDevice_GetCenterFrequency_AirspyDevice_1c209d1ee771b908(C.uintptr_t(_swig_i_0)))
	return swig_r
}

func (arg1 SwigcptrAirspyDevice) GetName() (_swig_ret string) {
	var swig_r string
	_swig_i_0 := arg1
	swig_r_p := C._wrap_AirspyDevice_GetName_AirspyDevice_1c209d1ee771b908(C.uintptr_t(_swig_i_0))
	swig_r = *(*string)(unsafe.Pointer(&swig_r_p))
	var swig_r_1 string
 swig_r_1 = swigCopyString(swig_r) 
	return swig_r_1
}

func (arg1 SwigcptrAirspyDevice) GetSampleRate() (_swig_ret uint) {
	var swig_r uint
	_swig_i_0 := arg1
	swig_r = (uint)(C._wrap_AirspyDevice_GetSampleRate_AirspyDevice_1c209d1ee771b908(C.uintptr_t(_swig_i_0)))
	return swig_r
}

func (arg1 SwigcptrAirspyDevice) SetSamplesAvailableCallback(arg2 AirspyDeviceCallback) {
	_swig_i_0 := arg1
	_swig_i_1 := arg2.Swigcptr()
	C._wrap_AirspyDevice_SetSamplesAvailableCallback_AirspyDevice_1c209d1ee771b908(C.uintptr_t(_swig_i_0), C.uintptr_t(_swig_i_1))
}

type AirspyDevice interface {
	Swigcptr() uintptr
	SwigIsAirspyDevice()
	SetSampleRate(arg2 uint) (_swig_ret uint)
	SetCenterFrequency(arg2 uint) (_swig_ret uint)
	GetAvailableSampleRates() (_swig_ret Vector32u)
	Start()
	Stop()
	SetAGC(arg2 bool)
	Init() (_swig_ret bool)
	Destroy()
	SetLNAGain(arg2 byte)
	SetVGAGain(arg2 byte)
	SetMixerGain(arg2 byte)
	SetBiasT(arg2 byte)
	GetCenterFrequency() (_swig_ret uint)
	GetName() (_swig_ret string)
	GetSampleRate() (_swig_ret uint)
	SetSamplesAvailableCallback(arg2 AirspyDeviceCallback)
}


type SwigcptrSwigDirector_AirspyDeviceCallback uintptr
type SwigDirector_AirspyDeviceCallback interface {
	Swigcptr() uintptr;
}
func (p SwigcptrSwigDirector_AirspyDeviceCallback) Swigcptr() uintptr {
	return uintptr(p)
}



var swigDirectorTrack struct {
	sync.Mutex
	m map[int]interface{}
	c int
}

func swigDirectorAdd(v interface{}) int {
	swigDirectorTrack.Lock()
	defer swigDirectorTrack.Unlock()
	if swigDirectorTrack.m == nil {
		swigDirectorTrack.m = make(map[int]interface{})
	}
	swigDirectorTrack.c++
	ret := swigDirectorTrack.c
	swigDirectorTrack.m[ret] = v
	return ret
}

func swigDirectorLookup(c int) interface{} {
	swigDirectorTrack.Lock()
	defer swigDirectorTrack.Unlock()
	ret := swigDirectorTrack.m[c]
	if ret == nil {
		panic("C++ director pointer not found (possible	use-after-free)")
	}
	return ret
}

func swigDirectorDelete(c int) {
	swigDirectorTrack.Lock()
	defer swigDirectorTrack.Unlock()
	if swigDirectorTrack.m[c] == nil {
		if c > swigDirectorTrack.c {
			panic("C++ director pointer invalid (possible memory corruption")
		} else {
			panic("C++ director pointer not found (possible use-after-free)")
		}
	}
	delete(swigDirectorTrack.m, c)
}


