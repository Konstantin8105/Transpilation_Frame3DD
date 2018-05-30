/*
	Package main - transpiled by c4go

	If you have found any issues, please raise an issue at:
	https://github.com/Konstantin8105/c4go/
*/

// Warning (*ast.EnumDecl):  /usr/include/x86_64-linux-gnu/bits/confname.h:71 :Add support of continues counter for type : *ast.BinaryExpr
// Warning (*ast.EnumDecl):  /usr/include/ctype.h:46 :Add support of continues counter for type : *ast.BinaryExpr
// Warning (*ast.SwitchStmt):  /home/konstantin/go/src/github.com/Konstantin8105/History_frame3DD/src/HPGutil.c:126 :Unsupport case
// Warning Function `noarch.InterfaceToFloat32SliceSlice` haven`t implementation
// Warning Function `noarch.InterfaceToFloat64SliceSlice` haven`t implementation
// Warning Function `noarch.InterfaceToIntSliceSlice` haven`t implementation
// Warning Function `noarch.InterfaceToFloat32SliceSlice` haven`t implementation
// Warning Function `noarch.InterfaceToFcomplexSliceSlice` haven`t implementation
// Warning Function `noarch.Float64SliceSliceSliceToFloat32SliceSliceSlice` haven`t implementation
// Warning Function `noarch.InterfaceToFloat64SliceSlice` haven`t implementation
// Warning Function `noarch.Float64SliceSliceToFloat32SliceSlice` haven`t implementation
// Warning Function `noarch.InterfaceToFloat64SliceSlice` haven`t implementation
// Warning Function `noarch.Float64SliceSliceToFloat32SliceSlice` haven`t implementation
// Warning Function `noarch.Float64SliceToFloat32Slice` haven`t implementation
// Warning Function `noarch.InterfaceToFloat64SliceSlice` haven`t implementation
// Warning (*ast.TranslationUnitDecl):  :0 :cannot transpileRecordDecl `__WAIT_STATUS`. could not determine the size of type `union __WAIT_STATUS` for that reason: Cannot determine sizeof : |union __WAIT_STATUS|. err = Cannot canculate `union` sizeof for `string`. Cannot determine sizeof : |union wait *|. err = error in union

package main

import "github.com/Konstantin8105/c4go/linux"
import "math"
import "math/rand"
import "unsafe"
import "os"
import "github.com/Konstantin8105/c4go/noarch"

const (
	FP_NAN       int = 0
	FP_INFINITE      = 1
	FP_ZERO          = 2
	FP_SUBNORMAL     = 3
	FP_NORMAL        = 4
)
const (
	_IEEE_  int = -1
	_SVID_      = 0
	_XOPEN_     = 1
	_POSIX_     = 2
	_ISOC_      = 3
)

type size_t uint32
type __time_t int32
type va_list int64
type __gnuc_va_list int64
type __codecvt_result int

const (
	__codecvt_ok      __codecvt_result = 0
	__codecvt_partial                  = 1
	__codecvt_error                    = 2
	__codecvt_noconv                   = 3
)

var stdin *noarch.File

var stdout *noarch.File

var stderr *noarch.File

type wchar_t int

const (
	P_ALL  int = 0
	P_PID      = 1
	P_PGID     = 2
)

type __locale_t int
type locale_t __locale_t
type vec3_struct struct {
	x float64
	y float64
	z float64
}
type vec3 vec3_struct

const (
	_PC_LINK_MAX           int = 0
	_PC_MAX_CANON              = 1
	_PC_MAX_INPUT              = 2
	_PC_NAME_MAX               = 3
	_PC_PATH_MAX               = 4
	_PC_PIPE_BUF               = 5
	_PC_CHOWN_RESTRICTED       = 6
	_PC_NO_TRUNC               = 7
	_PC_VDISABLE               = 8
	_PC_SYNC_IO                = 9
	_PC_ASYNC_IO               = 10
	_PC_PRIO_IO                = 11
	_PC_SOCK_MAXBUF            = 12
	_PC_FILESIZEBITS           = 13
	_PC_REC_INCR_XFER_SIZE     = 14
	_PC_REC_MAX_XFER_SIZE      = 15
	_PC_REC_MIN_XFER_SIZE      = 16
	_PC_REC_XFER_ALIGN         = 17
	_PC_ALLOC_SIZE_MIN         = 18
	_PC_SYMLINK_MAX            = 19
	_PC_2_SYMLINKS             = 20
)
const (
	_SC_ARG_MAX                      int = 0
	_SC_CHILD_MAX                        = 1
	_SC_CLK_TCK                          = 2
	_SC_NGROUPS_MAX                      = 3
	_SC_OPEN_MAX                         = 4
	_SC_STREAM_MAX                       = 5
	_SC_TZNAME_MAX                       = 6
	_SC_JOB_CONTROL                      = 7
	_SC_SAVED_IDS                        = 8
	_SC_REALTIME_SIGNALS                 = 9
	_SC_PRIORITY_SCHEDULING              = 10
	_SC_TIMERS                           = 11
	_SC_ASYNCHRONOUS_IO                  = 12
	_SC_PRIORITIZED_IO                   = 13
	_SC_SYNCHRONIZED_IO                  = 14
	_SC_FSYNC                            = 15
	_SC_MAPPED_FILES                     = 16
	_SC_MEMLOCK                          = 17
	_SC_MEMLOCK_RANGE                    = 18
	_SC_MEMORY_PROTECTION                = 19
	_SC_MESSAGE_PASSING                  = 20
	_SC_SEMAPHORES                       = 21
	_SC_SHARED_MEMORY_OBJECTS            = 22
	_SC_AIO_LISTIO_MAX                   = 23
	_SC_AIO_MAX                          = 24
	_SC_AIO_PRIO_DELTA_MAX               = 25
	_SC_DELAYTIMER_MAX                   = 26
	_SC_MQ_OPEN_MAX                      = 27
	_SC_MQ_PRIO_MAX                      = 28
	_SC_VERSION                          = 29
	_SC_PAGESIZE                         = 30
	_SC_RTSIG_MAX                        = 31
	_SC_SEM_NSEMS_MAX                    = 32
	_SC_SEM_VALUE_MAX                    = 33
	_SC_SIGQUEUE_MAX                     = 34
	_SC_TIMER_MAX                        = 35
	_SC_BC_BASE_MAX                      = 36
	_SC_BC_DIM_MAX                       = 37
	_SC_BC_SCALE_MAX                     = 38
	_SC_BC_STRING_MAX                    = 39
	_SC_COLL_WEIGHTS_MAX                 = 40
	_SC_EQUIV_CLASS_MAX                  = 41
	_SC_EXPR_NEST_MAX                    = 42
	_SC_LINE_MAX                         = 43
	_SC_RE_DUP_MAX                       = 44
	_SC_CHARCLASS_NAME_MAX               = 45
	_SC_2_VERSION                        = 46
	_SC_2_C_BIND                         = 47
	_SC_2_C_DEV                          = 48
	_SC_2_FORT_DEV                       = 49
	_SC_2_FORT_RUN                       = 50
	_SC_2_SW_DEV                         = 51
	_SC_2_LOCALEDEF                      = 52
	_SC_PII                              = 53
	_SC_PII_XTI                          = 54
	_SC_PII_SOCKET                       = 55
	_SC_PII_INTERNET                     = 56
	_SC_PII_OSI                          = 57
	_SC_POLL                             = 58
	_SC_SELECT                           = 59
	_SC_UIO_MAXIOV                       = 60
	_SC_IOV_MAX                          = 61
	_SC_PII_INTERNET_STREAM              = 62
	_SC_PII_INTERNET_DGRAM               = 63
	_SC_PII_OSI_COTS                     = 64
	_SC_PII_OSI_CLTS                     = 65
	_SC_PII_OSI_M                        = 66
	_SC_T_IOV_MAX                        = 67
	_SC_THREADS                          = 68
	_SC_THREAD_SAFE_FUNCTIONS            = 69
	_SC_GETGR_R_SIZE_MAX                 = 70
	_SC_GETPW_R_SIZE_MAX                 = 71
	_SC_LOGIN_NAME_MAX                   = 72
	_SC_TTY_NAME_MAX                     = 73
	_SC_THREAD_DESTRUCTOR_ITERATIONS     = 74
	_SC_THREAD_KEYS_MAX                  = 75
	_SC_THREAD_STACK_MIN                 = 76
	_SC_THREAD_THREADS_MAX               = 77
	_SC_THREAD_ATTR_STACKADDR            = 78
	_SC_THREAD_ATTR_STACKSIZE            = 79
	_SC_THREAD_PRIORITY_SCHEDULING       = 80
	_SC_THREAD_PRIO_INHERIT              = 81
	_SC_THREAD_PRIO_PROTECT              = 82
	_SC_THREAD_PROCESS_SHARED            = 83
	_SC_NPROCESSORS_CONF                 = 84
	_SC_NPROCESSORS_ONLN                 = 85
	_SC_PHYS_PAGES                       = 86
	_SC_AVPHYS_PAGES                     = 87
	_SC_ATEXIT_MAX                       = 88
	_SC_PASS_MAX                         = 89
	_SC_XOPEN_VERSION                    = 90
	_SC_XOPEN_XCU_VERSION                = 91
	_SC_XOPEN_UNIX                       = 92
	_SC_XOPEN_CRYPT                      = 93
	_SC_XOPEN_ENH_I18N                   = 94
	_SC_XOPEN_SHM                        = 95
	_SC_2_CHAR_TERM                      = 96
	_SC_2_C_VERSION                      = 97
	_SC_2_UPE                            = 98
	_SC_XOPEN_XPG2                       = 99
	_SC_XOPEN_XPG3                       = 100
	_SC_XOPEN_XPG4                       = 101
	_SC_CHAR_BIT                         = 102
	_SC_CHAR_MAX                         = 103
	_SC_CHAR_MIN                         = 104
	_SC_INT_MAX                          = 105
	_SC_INT_MIN                          = 106
	_SC_LONG_BIT                         = 107
	_SC_WORD_BIT                         = 108
	_SC_MB_LEN_MAX                       = 109
	_SC_NZERO                            = 110
	_SC_SSIZE_MAX                        = 111
	_SC_SCHAR_MAX                        = 112
	_SC_SCHAR_MIN                        = 113
	_SC_SHRT_MAX                         = 114
	_SC_SHRT_MIN                         = 115
	_SC_UCHAR_MAX                        = 116
	_SC_UINT_MAX                         = 117
	_SC_ULONG_MAX                        = 118
	_SC_USHRT_MAX                        = 119
	_SC_NL_ARGMAX                        = 120
	_SC_NL_LANGMAX                       = 121
	_SC_NL_MSGMAX                        = 122
	_SC_NL_NMAX                          = 123
	_SC_NL_SETMAX                        = 124
	_SC_NL_TEXTMAX                       = 125
	_SC_XBS5_ILP32_OFF32                 = 126
	_SC_XBS5_ILP32_OFFBIG                = 127
	_SC_XBS5_LP64_OFF64                  = 128
	_SC_XBS5_LPBIG_OFFBIG                = 129
	_SC_XOPEN_LEGACY                     = 130
	_SC_XOPEN_REALTIME                   = 131
	_SC_XOPEN_REALTIME_THREADS           = 132
	_SC_ADVISORY_INFO                    = 133
	_SC_BARRIERS                         = 134
	_SC_BASE                             = 135
	_SC_C_LANG_SUPPORT                   = 136
	_SC_C_LANG_SUPPORT_R                 = 137
	_SC_CLOCK_SELECTION                  = 138
	_SC_CPUTIME                          = 139
	_SC_THREAD_CPUTIME                   = 140
	_SC_DEVICE_IO                        = 141
	_SC_DEVICE_SPECIFIC                  = 142
	_SC_DEVICE_SPECIFIC_R                = 143
	_SC_FD_MGMT                          = 144
	_SC_FIFO                             = 145
	_SC_PIPE                             = 146
	_SC_FILE_ATTRIBUTES                  = 147
	_SC_FILE_LOCKING                     = 148
	_SC_FILE_SYSTEM                      = 149
	_SC_MONOTONIC_CLOCK                  = 150
	_SC_MULTI_PROCESS                    = 151
	_SC_SINGLE_PROCESS                   = 152
	_SC_NETWORKING                       = 153
	_SC_READER_WRITER_LOCKS              = 154
	_SC_SPIN_LOCKS                       = 155
	_SC_REGEXP                           = 156
	_SC_REGEX_VERSION                    = 157
	_SC_SHELL                            = 158
	_SC_SIGNALS                          = 159
	_SC_SPAWN                            = 160
	_SC_SPORADIC_SERVER                  = 161
	_SC_THREAD_SPORADIC_SERVER           = 162
	_SC_SYSTEM_DATABASE                  = 163
	_SC_SYSTEM_DATABASE_R                = 164
	_SC_TIMEOUTS                         = 165
	_SC_TYPED_MEMORY_OBJECTS             = 166
	_SC_USER_GROUPS                      = 167
	_SC_USER_GROUPS_R                    = 168
	_SC_2_PBS                            = 169
	_SC_2_PBS_ACCOUNTING                 = 170
	_SC_2_PBS_LOCATE                     = 171
	_SC_2_PBS_MESSAGE                    = 172
	_SC_2_PBS_TRACK                      = 173
	_SC_SYMLOOP_MAX                      = 174
	_SC_STREAMS                          = 175
	_SC_2_PBS_CHECKPOINT                 = 176
	_SC_V6_ILP32_OFF32                   = 177
	_SC_V6_ILP32_OFFBIG                  = 178
	_SC_V6_LP64_OFF64                    = 179
	_SC_V6_LPBIG_OFFBIG                  = 180
	_SC_HOST_NAME_MAX                    = 181
	_SC_TRACE                            = 182
	_SC_TRACE_EVENT_FILTER               = 183
	_SC_TRACE_INHERIT                    = 184
	_SC_TRACE_LOG                        = 185
	_SC_LEVEL1_ICACHE_SIZE               = 186
	_SC_LEVEL1_ICACHE_ASSOC              = 187
	_SC_LEVEL1_ICACHE_LINESIZE           = 188
	_SC_LEVEL1_DCACHE_SIZE               = 189
	_SC_LEVEL1_DCACHE_ASSOC              = 190
	_SC_LEVEL1_DCACHE_LINESIZE           = 191
	_SC_LEVEL2_CACHE_SIZE                = 192
	_SC_LEVEL2_CACHE_ASSOC               = 193
	_SC_LEVEL2_CACHE_LINESIZE            = 194
	_SC_LEVEL3_CACHE_SIZE                = 195
	_SC_LEVEL3_CACHE_ASSOC               = 196
	_SC_LEVEL3_CACHE_LINESIZE            = 197
	_SC_LEVEL4_CACHE_SIZE                = 198
	_SC_LEVEL4_CACHE_ASSOC               = 199
	_SC_LEVEL4_CACHE_LINESIZE            = 200
	_SC_IPV6                             = _SC_LEVEL1_ICACHE_SIZE + 50
	_SC_RAW_SOCKETS                      = 201
	_SC_V7_ILP32_OFF32                   = 202
	_SC_V7_ILP32_OFFBIG                  = 203
	_SC_V7_LP64_OFF64                    = 204
	_SC_V7_LPBIG_OFFBIG                  = 205
	_SC_SS_REPL_MAX                      = 206
	_SC_TRACE_EVENT_NAME_MAX             = 207
	_SC_TRACE_NAME_MAX                   = 208
	_SC_TRACE_SYS_MAX                    = 209
	_SC_TRACE_USER_EVENT_MAX             = 210
	_SC_XOPEN_STREAMS                    = 211
	_SC_THREAD_ROBUST_PRIO_INHERIT       = 212
	_SC_THREAD_ROBUST_PRIO_PROTECT       = 213
)
const (
	_CS_PATH                            int = 0
	_CS_V6_WIDTH_RESTRICTED_ENVS            = 1
	_CS_GNU_LIBC_VERSION                    = 2
	_CS_GNU_LIBPTHREAD_VERSION              = 3
	_CS_V5_WIDTH_RESTRICTED_ENVS            = 4
	_CS_V7_WIDTH_RESTRICTED_ENVS            = 5
	_CS_LFS_CFLAGS                          = 1000
	_CS_LFS_LDFLAGS                         = 1001
	_CS_LFS_LIBS                            = 1002
	_CS_LFS_LINTFLAGS                       = 1003
	_CS_LFS64_CFLAGS                        = 1004
	_CS_LFS64_LDFLAGS                       = 1005
	_CS_LFS64_LIBS                          = 1006
	_CS_LFS64_LINTFLAGS                     = 1007
	_CS_XBS5_ILP32_OFF32_CFLAGS             = 1100
	_CS_XBS5_ILP32_OFF32_LDFLAGS            = 1101
	_CS_XBS5_ILP32_OFF32_LIBS               = 1102
	_CS_XBS5_ILP32_OFF32_LINTFLAGS          = 1103
	_CS_XBS5_ILP32_OFFBIG_CFLAGS            = 1104
	_CS_XBS5_ILP32_OFFBIG_LDFLAGS           = 1105
	_CS_XBS5_ILP32_OFFBIG_LIBS              = 1106
	_CS_XBS5_ILP32_OFFBIG_LINTFLAGS         = 1107
	_CS_XBS5_LP64_OFF64_CFLAGS              = 1108
	_CS_XBS5_LP64_OFF64_LDFLAGS             = 1109
	_CS_XBS5_LP64_OFF64_LIBS                = 1110
	_CS_XBS5_LP64_OFF64_LINTFLAGS           = 1111
	_CS_XBS5_LPBIG_OFFBIG_CFLAGS            = 1112
	_CS_XBS5_LPBIG_OFFBIG_LDFLAGS           = 1113
	_CS_XBS5_LPBIG_OFFBIG_LIBS              = 1114
	_CS_XBS5_LPBIG_OFFBIG_LINTFLAGS         = 1115
	_CS_POSIX_V6_ILP32_OFF32_CFLAGS         = 1116
	_CS_POSIX_V6_ILP32_OFF32_LDFLAGS        = 1117
	_CS_POSIX_V6_ILP32_OFF32_LIBS           = 1118
	_CS_POSIX_V6_ILP32_OFF32_LINTFLAGS      = 1119
	_CS_POSIX_V6_ILP32_OFFBIG_CFLAGS        = 1120
	_CS_POSIX_V6_ILP32_OFFBIG_LDFLAGS       = 1121
	_CS_POSIX_V6_ILP32_OFFBIG_LIBS          = 1122
	_CS_POSIX_V6_ILP32_OFFBIG_LINTFLAGS     = 1123
	_CS_POSIX_V6_LP64_OFF64_CFLAGS          = 1124
	_CS_POSIX_V6_LP64_OFF64_LDFLAGS         = 1125
	_CS_POSIX_V6_LP64_OFF64_LIBS            = 1126
	_CS_POSIX_V6_LP64_OFF64_LINTFLAGS       = 1127
	_CS_POSIX_V6_LPBIG_OFFBIG_CFLAGS        = 1128
	_CS_POSIX_V6_LPBIG_OFFBIG_LDFLAGS       = 1129
	_CS_POSIX_V6_LPBIG_OFFBIG_LIBS          = 1130
	_CS_POSIX_V6_LPBIG_OFFBIG_LINTFLAGS     = 1131
	_CS_POSIX_V7_ILP32_OFF32_CFLAGS         = 1132
	_CS_POSIX_V7_ILP32_OFF32_LDFLAGS        = 1133
	_CS_POSIX_V7_ILP32_OFF32_LIBS           = 1134
	_CS_POSIX_V7_ILP32_OFF32_LINTFLAGS      = 1135
	_CS_POSIX_V7_ILP32_OFFBIG_CFLAGS        = 1136
	_CS_POSIX_V7_ILP32_OFFBIG_LDFLAGS       = 1137
	_CS_POSIX_V7_ILP32_OFFBIG_LIBS          = 1138
	_CS_POSIX_V7_ILP32_OFFBIG_LINTFLAGS     = 1139
	_CS_POSIX_V7_LP64_OFF64_CFLAGS          = 1140
	_CS_POSIX_V7_LP64_OFF64_LDFLAGS         = 1141
	_CS_POSIX_V7_LP64_OFF64_LIBS            = 1142
	_CS_POSIX_V7_LP64_OFF64_LINTFLAGS       = 1143
	_CS_POSIX_V7_LPBIG_OFFBIG_CFLAGS        = 1144
	_CS_POSIX_V7_LPBIG_OFFBIG_LDFLAGS       = 1145
	_CS_POSIX_V7_LPBIG_OFFBIG_LIBS          = 1146
	_CS_POSIX_V7_LPBIG_OFFBIG_LINTFLAGS     = 1147
	_CS_V6_ENV                              = 1148
	_CS_V7_ENV                              = 1149
)

type FCOMPLEX struct {
	r float32
	i float32
}
type fcomplex FCOMPLEX

// main - transpiled function from  $GOPATH/src/github.com/Konstantin8105/History_frame3DD/src/main.c:65
/*
   This file is part of FRAME3DD:
 Static and dynamic structural analysis of 2D and 3D frames and trusses with
 elastic and geometric stiffness.
 ---------------------------------------------------------------------------
 http://frame3dd.sourceforge.net/
 ---------------------------------------------------------------------------
 Copyright (C) 1992-2014  Henri P. Gavin

    FRAME3DD is free software: you can redistribute it and/or modify
    it under the terms of the GNU General Public License as published by
    the Free Software Foundation, either version 3 of the License, or
    (at your option) any later version.

    FRAME3DD is distributed in the hope that it will be useful,
    but WITHOUT ANY WARRANTY; without even the implied warranty of
    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
    GNU General Public License for more details.

    You should have received a copy of the GNU General Public License
    along with FRAME3DD.  If not, see <http://www.gnu.org/licenses/>.
*/ //
/** @file
Main FRAME3DD program driver
*/ //
/** @mainpage
FRAME3DD: a program for static and dynamic structural analysis of 2D and 3D
frames and trusses with elastic and geometric stiffness.

Also included is a system for parsing Microstran .arc 'Archive' files and
for parsing calculated force and displacement output files (.p1 format) from
Microstran. See @ref mstranp. It is intended that ultimately the .arc format
be an alternative method of inputting data to the FRAME3DD program,
but currently these two parts of the code are distinct.

For more information go to http://frame3dd.sourceforge.net/

The input file format for FRAME is defined in doc/user_manual.html

Henri P. Gavin hpgavin@duke.edu (main FRAME3DD code)
John Pye john.pye@anu.edu.au (Microstran parser and viewer)

For compilation/installation, see README.txt.

*/ //
// compile the Frame3DD analysis into another code, such as a GUI
// compile Frame3DD to run as a stand-alone code through the terminal
// the input  data filename
// the output data filename
// the title of the analysis
// the text of an error message
// mesh data path
// plot file path
// int  file path
// mode data path
// temp data path
// input and output file pointer
// X,Y,Z node coordinates (global)
// node size radius, for finite sizes
// cross section areas, incl. shear
// section inertias
// elastic modulus and shear moduli
// roll of each member, radians
// uniform distributed member loads
// trapizoidal distributed member loads
// member concentrated loads
// member temperature  loads
// prescribed node displacements
// member densities and extra inertia
// mass of a node
// inertia of a node in global coord
// gravitational acceleration in global X
// gravitational acceleration in global Y
// gravitational acceleration in global Z
// >0: pan during animation; 0: don't
// zoom scale for 3D plotting in Gnuplot
// x-increment for internal force data
// equilibrium stiffness matrix
// **Ks=NULL,	// Broyden secant stiffness matrix
// trace of the global stiffness matrix
// global mass matrix
// trace of the global mass matrix
// equivalent end forces from mech loads global
// equivalent end forces from temp loads global
// mechanical load vectors, all load cases
// thermal load vectors, all load cases
// total load vectors for a load case
// total reaction force vector
// incremental reaction force vector
// displacement vector
// incremental displacement vector
//dDdD = 0.0,	// dD' * dD
// equilibrium error in nonlinear anlys
// node-to-node length of each element
// effcve lngth, accounts for node size
// local member node end-forces
// tolerance for modal convergence
// shift-factor for rigid-body-modes
// mass of structural system
// total structural mass and extra mass
// resonant frequencies
// resonant mode-shapes
// root mean square of residual displ. error
// rms equilibrium error and reactions
// frequency used for Guyan condensation
// condensed stiffness and mass matrices
// exaggerate static displ. in mesh data
// exaggerate modal displ. in mesh data
// peak internal forces, moments, and displacments
// in each frame element and each load case
// number of Nodes
// number of frame Elements
// number of Load cases
// number of Degrees of Freedom
// number of restrained nodes
// number of prescribed nodal displ'nts
// number of loaded nodes
// number of members w/ unifm dist loads
// number of members w/ trapz dist loads
// number of members w/ conc point loads
// number of members w/ temp. changes
// number of nodes w/ extra inertia
// number of elemts w/ extra mass
// number of condensed nodes
// begin and end node numbers
// indicates shear deformation
// indicates  geometric nonlinearity
// 1: stiffness analysis, 0: data check
// reaction data, total no. of reactions
// number of desired modes
// 1: Subspace Jacobi, 2: Stodola
// number of modes to calculate
// 1: lumped, 0: consistent mass matrix
// number of iterations
// number of (-ve) diag. terms of L D L'
// the modes to be animated
// number of condensed degrees o freedom
// matrix condensation method
// vector of DoF's to condense
// vector of modes to condense
// 1 if .CSV, 2 if file is Matlab
// 1: debugging screen output, 0: none
// 1: copious screen output, 0: none
// 0: "ok", 1: strain > 0.001
// error code returned by Frame3DD
//   over-ride input file value
//   over-ride input file value
//   over-ride input file value
//   over-ride 3D plotting check
//   over-ride input file value
//   over-ride input file value
//   write stiffness and mass matrix
//   suppress 't' or 'c' in output data
// over-ride input file value
// *scanf return value for err checking
// over-ride input file value
// over-ride input file value
// over-ride input file value
// over-ride input file value
// Input Data file name extension
/*
	parse_options ( argc, argv, IN_file, OUT_file,
			&shear_flag, &geom_flag, &anlyz_flag, &exagg_flag,
			&D3_flag,
			&lump_flag, &modal_flag, &tol_flag, &shift_flag,
			&pan_flag, &write_matrix, &axial_sign, &condense_flag,
			&verbose, &debug);
*/ //
/*  display program name, version and license type */ //
//frame3dd.sf.net\n");
/* open the input data file */ //
/* open input data file */ //
// clear the buffer ??
// wait for the Enter key
/* .CSV or .FMM or other? */ //
//	temp_file_location("frame3dd.3dd",strippedInputFile,FRAME3DD_PATHMAX);
/* strip comments from input data */ //
/* open stripped input file */ //
/*  display analysis title */ //
/* number of nodes	*/ //
/* display nN */ //
/* allocate memory for node data ... */ //
/* rigid radius around each node */ //
/* node coordinates */ //
/* total number of degrees of freedom	*/ //
/* allocate memory for reaction data ... */ //
/* allocate memory for reaction data ... */ //
/* number of frame elements	*/ //
/* display nE */ //
/* not enough elements */ //
/* allocate memory for frame elements ... */ //
/* length of each element		*/ //
/* effective length of each element	*/ //
/* node #1 of each element		*/ //
/* node #2 of each element		*/ //
/* cross section area of each element	*/ //
/* shear area in local y direction 	*/ //
/* shear area in local z direction	*/ //
/* torsional moment of inertia 		*/ //
/* bending moment of inertia about y-axis */ //
/* bending moment of inertia about z-axis */ //
/* frame element Young's modulus	*/ //
/* frame element shear modulus		*/ //
/* element rotation angle about local x axis */ //
/* element mass density			*/ //
/* number of load cases		*/ //
/* display nL */ //
/* not enough load cases */ //
/* too many load cases */ //
/* allocate memory for loads ... */ //
/* uniform load on each member */ //
/* trapezoidal load on each member */ //
/* internal point load each member */ //
/* internal temp change each member*/ //
/* prescribed displacement of each node */ //
/* mechanical load vector	*/ //
/* temperature load vector	*/ //
/* external load vector	*/ //
/* equilibrium error {F} - [K]{D} */ //
/* eqF due to mech loads */ //
/* eqF due to temp loads */ //
/* global stiffness matrix	*/ //
/* end forces for each member	*/ //
/* displacments of each node		*/ //
/* incremental displ. of each node	*/ //
/* reaction forces			*/ //
/* incremental reaction forces		*/ //
/* lumped mass for each frame element	*/ //
/* node mass for each node		*/ //
/* node inertia about global X axis	*/ //
/* node inertia about global Y axis	*/ //
/* node inertia about global Z axis	*/ //
/* vector of condensed degrees of freedom */ //
/* vector of condensed mode numbers	*/ //
// peak axial forces, shears, torques, and moments along each element
// peak displacements and slopes along each element
/* display load data complete */ //
/* display mass data complete */ //
/*  display condensation data complete */ //
/* close the input data file */ //
/* open the output data file for appending */ //
/* unable to append to output data file */ //
/* solve the problem	*/ //
/* begin load case analysis loop */ //
/* display the load case number  */ //
/*  initialize displacements and displ. increment to {0}  */ //
/*  initialize reactions     and react. increment to {0}  */ //
/*  initialize internal element end forces Q = {0}	*/ //
/*  elastic stiffness matrix  [K({D}^(i))], {D}^(0)={0} (i=0) */ //
/* first apply temperature loads only, if there are any ... */ //
/*  solve {F_t} = [K({D=0})] * {D_t} */ //
/* increment {D_t} = {0} + {D_t} temp.-induced displ */ //
/* increment {R_t} = {0} + {R_t} temp.-induced react */ //
/* assemble K = Ke + Kg */ //
/* compute   {Q}={Q_t} ... temp.-induced forces     */ //
/* assemble temp.-stressed stiffness [K({D_t})]     */ //
/* ... then apply mechanical loads only, if there are any ... */ //
/* incremental displ at react'ns = prescribed displ */ //
/*  solve {F_m} = [K({D_t})] * {D_m}	*/ //
/* combine {D} = {D_t} + {D_m}	*/ //
/* combine {R} = {R_t} + {R_m} --- for linear systems */ //
/*  combine {F} = {F_t} + {F_m} */ //
/*  element forces {Q} for displacements {D}	*/ //
/*  check the equilibrium error	*/ //
/*
 * 		if ( geom ) { // initialize Broyden secant stiffness matrix, Ks
 *			Ks  = dmatrix( 1, DoF, 1, DoF );
 *			for (i=1;i<=DoF;i++) {
 *				for(j=i;j<=DoF;j++) {
 *					Ks[i][j]=Ks[j][i]=K[i][j];
 *				}
 *			}
 *		}
 */ //
/* quasi Newton-Raphson iteration for geometric nonlinearity  */ //
/* re-initialize */ //
/*  assemble stiffness matrix [K({D}^(i))]	      */ //
/*  compute equilibrium error, {dF}, at iteration i   */ //
/*  {dF}^(i) = {F} - [K({D}^(i))]*{D}^(i)	      */ //
/*  convergence criteria = || {dF}^(i) ||  /  || F || */ //
/*  Powell-Symmetric-Broyden secant stiffness update  */ //
// PSB_update ( Ks, dF, dD, DoF );  /* not helpful?   */
/*  solve {dF}^(i) = [K({D}^(i))] * {dD}^(i)	      */ //
/*  K is not positive definite	      */ //
/*  increment {D}^(i+1) = {D}^(i) + {dD}^(i)	      */ //
/*  element forces {Q} for displacements {D}^(i)      */ //
/*  display equilibrium error        */ //
/* end quasi Newton-Raphson iteration */ //
/*   strain limit failure ... */ //
/*   strain limit _and_ buckling failure ... */ //
/*  dealocate Broyden secant stiffness matrix, Ks */ //
// if ( geom )	free_dmatrix(Ks, 1, DoF, 1, DoF );
/* write static stiffness matrix */ //
/*  display RMS equilibrium error */ //
// .CSV format output
// .m matlab format output
/*
 *		if ( verbose )
 *		 printf("\n   If the program pauses here for very long,"
 *		 " hit CTRL-C to stop execution, \n"
 *		 "    reduce exagg_static in the Input Data,"
 *		 " and re-run the analysis. \n");
 */ //
/* end load case loop */ //
/*  data check only  */ //
/* display data check only */ //
/* carry out modal analysis */ //
/* Bathe */ //
/*  compute traceK and traceM */ //
/*  modify K and M for reactions    */ //
/* apply reactions to upper triangle */ //
/* write Kd and Md matrices */ //
/* dynamic stff matx */ //
/* dynamic mass matx */ //
/* subspace or stodola methods */ //
/* write modal analysis results */ //
/* matrix condensation of stiffness and mass */ //
/* static condensation only */ //
/*  dynamic condensation  */ //
/* deallocate memory used for each frame analysis variable */ //
/* wait for keyboard entry to close the terminal */ //
/* (void) getchar();	// clear the buffer ??  */ //
/* while( !getchar() ) ;	// wait for the Enter key to be hit  */ //
//
func main() {
	argc := len(os.Args)
	argv := [][]byte{}
	for _, argvSingle := range os.Args {
		argv = append(argv, []byte(argvSingle))
	}
	var IN_file []byte = make([]byte, 128)
	var OUT_file []byte = make([]byte, 128)
	var title []byte = make([]byte, 512)
	var errMsg []byte = make([]byte, 512)
	var meshpath []byte = []byte("EMPTY_MESH\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00")
	var plotpath []byte = []byte("EMPTY_PLOT\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00")
	var infcpath []byte = []byte("EMPTY_INFC\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00")
	var modepath []byte = []byte("EMPTY_MODE\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00")
	var strippedInputFile []byte = []byte("EMPTY_TEMP\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00")
	var fp *noarch.File
	var xyz []vec3
	var rj []float32
	var Ax []float32
	var Asy []float32
	var Asz []float32
	var Jx []float32
	var Iy []float32
	var Iz []float32
	var E []float32
	var G []float32
	var p []float32
	var U [][][]float32
	var W [][][]float32
	var P [][][]float32
	var T [][][]float32
	var Dp [][]float32
	var d []float32
	var EMs []float32
	var NMs []float32
	var NMx []float32
	var NMy []float32
	var NMz []float32
	var gX []float32 = make([]float32, 32)
	var gY []float32 = make([]float32, 32)
	var gZ []float32 = make([]float32, 32)
	var pan float32 = float32(1)
	var scale float32 = float32(1)
	var dx float32 = float32(1)
	var K [][]float64
	var traceK float64
	var M [][]float64
	var traceM float64
	var eqF_mech [][][]float64
	var eqF_temp [][][]float64
	var F_mech [][]float64
	var F_temp [][]float64
	var F []float64
	var R []float64
	var dR []float64
	var D []float64
	var dD []float64
	var dF []float64
	var L []float64
	var Le []float64
	var Q [][]float64
	var tol float64 = 1e-09
	var shift float64
	var struct_mass float64
	var total_mass float64
	var f []float64
	var V [][]float64
	var rms_resid float64 = 1
	var error float64 = 1
	var Cfreq float64
	var Kc [][]float64
	var Mc [][]float64
	var exagg_static float64 = 10
	var exagg_modal float64 = 10
	var pkNx [][]float64
	var pkVy [][]float64
	var pkVz [][]float64
	var pkTx [][]float64
	var pkMy [][]float64
	var pkMz [][]float64
	var pkDx [][]float64
	var pkDy [][]float64
	var pkDz [][]float64
	var pkRx [][]float64
	var pkSy [][]float64
	var pkSz [][]float64
	var nN int
	var nE int
	var nL int
	var lc int
	var DoF int
	var i int
	var j int
	var nR int
	var nD []int = make([]int, 32)
	var nF []int = make([]int, 32)
	var nU []int = make([]int, 32)
	var nW []int = make([]int, 32)
	var nP []int = make([]int, 32)
	var nT []int = make([]int, 32)
	var nI int
	var nX int
	var nC int
	var N1 []int
	var N2 []int
	var shear int
	var geom int
	var anlyz int = 1
	var q []int
	var r []int
	var sumR int
	var nM int
	var Mmethod int
	var nM_calc int
	var lump int = 1
	var iter int
	var ok int = 1
	var anim []int = make([]int, 128)
	var Cdof int
	var Cmethod int
	var c []int
	var m []int
	var filetype int
	var debug int
	var verbose int = 1
	var axial_strain_warning int
	var ExitCode int
	var shear_flag int = -1
	var geom_flag int = -1
	var anlyz_flag int = -1
	var D3_flag int = -1
	var lump_flag int = -1
	var modal_flag int = -1
	var write_matrix int = -1
	var axial_sign int = -1
	var condense_flag int = -1
	var sfrv int
	var exagg_flag float64 = -1
	var tol_flag float64 = -1
	var shift_flag float64 = -1
	var pan_flag float32 = float32(-1)
	var extn []byte = make([]byte, 16)
	noarch.Strcpy(IN_file, []byte("exJ.3dd\x00\x00"))
	noarch.Strcpy(OUT_file, []byte("exJ.3dd.out\x00\x00"))
	if verbose != 0 {
		textColor('w', 'b', 'b', 'x')
		noarch.Fprintf(stdout, []byte("\n FRAME3DD version: %s\n\x00"), []byte("20140514+\x00"))
		noarch.Fprintf(stdout, []byte(" Analysis of 2D and 3D structural frames with elastic and geometric stiffness.\n\x00"))
		noarch.Fprintf(stdout, []byte(" http://frame3dd.sf.net\n\x00"))
		noarch.Fprintf(stdout, []byte(" GPL Copyright (C) 1992-2014, Henri P. Gavin\n\x00"))
		noarch.Fprintf(stdout, []byte(" This is free software with absolutely no warranty.\n\x00"))
		noarch.Fprintf(stdout, []byte(" For details, see the GPL license file, LICENSE.txt\n\x00"))
		color(0)
		noarch.Fprintf(stdout, []byte("\n\x00"))
	}
	if (func() *noarch.File {
		fp = noarch.Fopen(IN_file, []byte("r\x00"))
		return fp
	}()) == nil {
		noarch.Sprintf(errMsg, []byte("\n ERROR: cannot open input data file '%s'\n\x00"), IN_file)
		errorMsg(errMsg)
		display_help()
		if argc == 1 {
			noarch.Fprintf(stderr, []byte(" Press the 'Enter' key to close.\n\x00"))
			_ = noarch.Getchar()
			for noarch.NotInt(noarch.Getchar()) != 0 {
			}
		}
		os.Exit(11)
	}
	noarch.Printf([]byte("strat filetype\x00"))
	filetype = get_file_ext(IN_file, extn)
	noarch.Printf([]byte("ed of filetype\n\x00"))
	output_path([]byte("frame3dd.3dd\x00"), strippedInputFile, 512, nil)
	parse_input(fp, strippedInputFile)
	noarch.Fclose(fp)
	if (func() *noarch.File {
		fp = noarch.Fopen(strippedInputFile, []byte("r\x00"))
		return fp
	}()) == nil {
		noarch.Sprintf(errMsg, []byte("\n ERROR: cannot open stripped input data file '%s'\n\x00"), strippedInputFile)
		errorMsg(errMsg)
		os.Exit(13)
	}
	frame3dd_getline(fp, title, 512)
	if verbose != 0 {
		textColor('w', 'g', 'b', 'x')
		noarch.Fprintf(stdout, []byte("\n\x00"))
		noarch.Fprintf(stdout, []byte(" ** %s ** \n\x00"), title)
		color(0)
		noarch.Fprintf(stdout, []byte("\n\x00"))
	}
	sfrv = noarch.Fscanf(fp, []byte("%d\x00"), (*[100000000]int)(unsafe.Pointer(&nN))[:])
	if sfrv != 1 {
		sferr([]byte("nN value for number of nodes\x00"))
	}
	if verbose != 0 {
		noarch.Fprintf(stdout, []byte(" number of nodes \x00"))
		dots(stdout, 36)
		noarch.Fprintf(stdout, []byte(" nN =%4d \x00"), nN)
	}
	rj = vector(1, int32(nN))
	xyz = make([]vec3, 24*uint32(1+nN)/24)
	read_node_data(fp, nN, xyz, rj)
	if verbose != 0 {
		noarch.Printf([]byte(" ... complete\n\x00"))
	}
	DoF = 6 * nN
	q = ivector(1, int32(DoF))
	r = ivector(1, int32(DoF))
	read_reaction_data(fp, DoF, nN, (*[100000000]int)(unsafe.Pointer(&nR))[:], q, r, (*[100000000]int)(unsafe.Pointer(&sumR))[:], verbose)
	if verbose != 0 {
		noarch.Fprintf(stdout, []byte(" ... complete\n\x00"))
	}
	sfrv = noarch.Fscanf(fp, []byte("%d\x00"), (*[100000000]int)(unsafe.Pointer(&nE))[:])
	if sfrv != 1 {
		sferr([]byte("nE value for number of frame elements\x00"))
	}
	if verbose != 0 {
		noarch.Fprintf(stdout, []byte(" number of frame elements\x00"))
		dots(stdout, 28)
		noarch.Fprintf(stdout, []byte(" nE =%4d \x00"), nE)
	}
	if nN > nE+1 {
		noarch.Fprintf(stderr, []byte("\n  warning: %d nodes and %d members...\x00"), nN, nE)
		noarch.Fprintf(stderr, []byte(" not enough elements to connect all nodes.\n\x00"))
	}
	L = dvector(1, int32(nE))
	Le = dvector(1, int32(nE))
	N1 = ivector(1, int32(nE))
	N2 = ivector(1, int32(nE))
	Ax = vector(1, int32(nE))
	Asy = vector(1, int32(nE))
	Asz = vector(1, int32(nE))
	Jx = vector(1, int32(nE))
	Iy = vector(1, int32(nE))
	Iz = vector(1, int32(nE))
	E = vector(1, int32(nE))
	G = vector(1, int32(nE))
	p = vector(1, int32(nE))
	d = vector(1, int32(nE))
	read_frame_element_data(fp, nN, nE, xyz, rj, L, Le, N1, N2, Ax, Asy, Asz, Jx, Iy, Iz, E, G, p, d)
	if verbose != 0 {
		noarch.Fprintf(stdout, []byte(" ... complete\n\x00"))
	}
	read_run_data(fp, OUT_file, (*[100000000]int)(unsafe.Pointer(&shear))[:], shear_flag, (*[100000000]int)(unsafe.Pointer(&geom))[:], geom_flag, meshpath, plotpath, infcpath, (*[100000000]float64)(unsafe.Pointer(&exagg_static))[:], exagg_flag, (*[100000000]float32)(unsafe.Pointer(&scale))[:], (*[100000000]float32)(unsafe.Pointer(&dx))[:], (*[100000000]int)(unsafe.Pointer(&anlyz))[:], anlyz_flag, debug)
	sfrv = noarch.Fscanf(fp, []byte("%d\x00"), (*[100000000]int)(unsafe.Pointer(&nL))[:])
	if sfrv != 1 {
		sferr([]byte("nL value for number of load cases\x00"))
	}
	if verbose != 0 {
		noarch.Fprintf(stdout, []byte(" number of load cases \x00"))
		dots(stdout, 31)
		noarch.Fprintf(stdout, []byte(" nL = %3d \n\x00"), nL)
	}
	if nL < 1 {
		errorMsg([]byte("\n ERROR: the number of load cases must be at least 1\n\x00"))
		os.Exit(101)
	}
	if nL >= 32 {
		noarch.Sprintf(errMsg, []byte("\n ERROR: maximum of %d load cases allowed\n\x00"), 32-1)
		errorMsg(errMsg)
		os.Exit(102)
	}
	U = D3matrix(1, nL, 1, nE, 1, 4)
	W = D3matrix(1, nL, 1, 10*nE, 1, 13)
	P = D3matrix(1, nL, 1, 10*nE, 1, 5)
	T = D3matrix(1, nL, 1, nE, 1, 8)
	Dp = matrix(1, int32(nL), 1, int32(DoF))
	F_mech = dmatrix(1, int32(nL), 1, int32(DoF))
	F_temp = dmatrix(1, int32(nL), 1, int32(DoF))
	F = dvector(1, int32(DoF))
	dF = dvector(1, int32(DoF))
	eqF_mech = D3dmatrix(1, nL, 1, nE, 1, 12)
	eqF_temp = D3dmatrix(1, nL, 1, nE, 1, 12)
	K = dmatrix(1, int32(DoF), 1, int32(DoF))
	Q = dmatrix(1, int32(nE), 1, 12)
	D = dvector(1, int32(DoF))
	dD = dvector(1, int32(DoF))
	R = dvector(1, int32(DoF))
	dR = dvector(1, int32(DoF))
	EMs = vector(1, int32(nE))
	NMs = vector(1, int32(nN))
	NMx = vector(1, int32(nN))
	NMy = vector(1, int32(nN))
	NMz = vector(1, int32(nN))
	c = ivector(1, int32(DoF))
	m = ivector(1, int32(DoF))
	pkNx = dmatrix(1, int32(nL), 1, int32(nE))
	pkVy = dmatrix(1, int32(nL), 1, int32(nE))
	pkVz = dmatrix(1, int32(nL), 1, int32(nE))
	pkTx = dmatrix(1, int32(nL), 1, int32(nE))
	pkMy = dmatrix(1, int32(nL), 1, int32(nE))
	pkMz = dmatrix(1, int32(nL), 1, int32(nE))
	pkDx = dmatrix(1, int32(nL), 1, int32(nE))
	pkDy = dmatrix(1, int32(nL), 1, int32(nE))
	pkDz = dmatrix(1, int32(nL), 1, int32(nE))
	pkRx = dmatrix(1, int32(nL), 1, int32(nE))
	pkSy = dmatrix(1, int32(nL), 1, int32(nE))
	pkSz = dmatrix(1, int32(nL), 1, int32(nE))
	read_and_assemble_loads(fp, nN, nE, nL, DoF, xyz, L, Le, N1, N2, Ax, Asy, Asz, Iy, Iz, E, G, p, d, gX, gY, gZ, r, shear, nF, nU, nW, nP, nT, nD, Q, F_temp, F_mech, F, U, W, P, T, Dp, eqF_mech, eqF_temp, verbose)
	if verbose != 0 {
		noarch.Fprintf(stdout, []byte("                                                     \x00"))
		noarch.Fprintf(stdout, []byte(" load data ... complete\n\x00"))
	}
	read_mass_data(fp, IN_file, nN, nE, (*[100000000]int)(unsafe.Pointer(&nI))[:], (*[100000000]int)(unsafe.Pointer(&nX))[:], d, EMs, NMs, NMx, NMy, NMz, L, Ax, (*[100000000]float64)(unsafe.Pointer(&total_mass))[:], (*[100000000]float64)(unsafe.Pointer(&struct_mass))[:], (*[100000000]int)(unsafe.Pointer(&nM))[:], (*[100000000]int)(unsafe.Pointer(&Mmethod))[:], modal_flag, (*[100000000]int)(unsafe.Pointer(&lump))[:], lump_flag, (*[100000000]float64)(unsafe.Pointer(&tol))[:], tol_flag, (*[100000000]float64)(unsafe.Pointer(&shift))[:], shift_flag, (*[100000000]float64)(unsafe.Pointer(&exagg_modal))[:], modepath, anim, (*[100000000]float32)(unsafe.Pointer(&pan))[:], pan_flag, verbose, debug)
	if verbose != 0 {
		noarch.Fprintf(stdout, []byte("                                                     \x00"))
		noarch.Fprintf(stdout, []byte(" mass data ... complete\n\x00"))
	}
	read_condensation_data(fp, nN, nM, (*[100000000]int)(unsafe.Pointer(&nC))[:], (*[100000000]int)(unsafe.Pointer(&Cdof))[:], (*[100000000]int)(unsafe.Pointer(&Cmethod))[:], condense_flag, c, m, verbose)
	if nC > 0 && verbose != 0 {
		noarch.Fprintf(stdout, []byte("                                      \x00"))
		noarch.Fprintf(stdout, []byte(" matrix condensation data ... complete\n\x00"))
	}
	noarch.Fclose(fp)
	fp = noarch.Fopen(OUT_file, []byte("a\x00"))
	if fp == nil {
		noarch.Fprintf(stderr, []byte("Unable to append to output data file '%s'!\n\x00"), OUT_file)
		os.Exit(14)
	}
	write_input_data(fp, title, nN, nE, nL, nD, nR, nF, nU, nW, nP, nT, xyz, rj, N1, N2, Ax, Asy, Asz, Jx, Iy, Iz, E, G, p, d, gX, gY, gZ, F_temp, F_mech, Dp, r, U, W, P, T, shear, anlyz, geom)
	if anlyz != 0 {
		rand.Seed(int64(uint32(int32((__time_t((noarch.Time(nil))))))))
		for lc = 1; lc <= nL; lc++ {
			if verbose != 0 {
				noarch.Fprintf(stdout, []byte("\n\x00"))
				textColor('y', 'g', 'b', 'x')
				noarch.Fprintf(stdout, []byte(" Load Case %d of %d ... \x00"), lc, nL)
				noarch.Fprintf(stdout, []byte("                                          \x00"))
				noarch.Fflush(stdout)
				color(0)
				noarch.Fprintf(stdout, []byte("\n\x00"))
			}
			for i = 1; i <= DoF; i++ {
				dR[i] = 0
				R[i] = dR[i]
				dD[i] = R[i]
				D[i] = dD[i]
			}
			for i = 1; i <= nE; i++ {
				for j = 1; j <= 12; j++ {
					Q[i][j] = 0
				}
			}
			assemble_K(K, DoF, nE, xyz, rj, L, Le, N1, N2, Ax, Asy, Asz, Jx, Iy, Iz, E, G, p, shear, geom, Q, debug)
			if nT[lc] > 0 {
				if verbose != 0 {
					noarch.Fprintf(stdout, []byte(" Linear Elastic Analysis ... Temperature Loads\n\x00"))
				}
				solve_system(K, dD, F_temp[lc], dR, DoF, q, r, (*[100000000]int)(unsafe.Pointer(&ok))[:], verbose, (*[100000000]float64)(unsafe.Pointer(&rms_resid))[:])
				for i = 1; i <= DoF; i++ {
					if q[i] != 0 {
						D[i] += dD[i]
					}
				}
				for i = 1; i <= DoF; i++ {
					if r[i] != 0 {
						R[i] += dR[i]
					}
				}
				noarch.Printf([]byte("STEP 1\x00"))
				if geom != 0 {
					noarch.Printf([]byte("GEOM\n\x00"))
					element_end_forces(Q, nE, xyz, L, Le, N1, N2, Ax, Asy, Asz, Jx, Iy, Iz, E, G, p, eqF_temp[lc], eqF_mech[lc], D, shear, geom, (*[100000000]int)(unsafe.Pointer(&axial_strain_warning))[:])
					noarch.Printf([]byte("STEP 2\x00"))
					assemble_K(K, DoF, nE, xyz, rj, L, Le, N1, N2, Ax, Asy, Asz, Jx, Iy, Iz, E, G, p, shear, geom, Q, debug)
				}
				noarch.Printf([]byte("STEP 3\x00"))
			}
			noarch.Printf([]byte("STEP 4\n\x00"))
			if nF[lc] > 0 || nU[lc] > 0 || nW[lc] > 0 || nP[lc] > 0 || nD[lc] > 0 || gX[lc] != 0 || gY[lc] != 0 || gZ[lc] != 0 {
				if verbose != 0 {
					noarch.Fprintf(stdout, []byte(" Linear Elastic Analysis ... Mechanical Loads\n\x00"))
				}
				noarch.Printf([]byte("STEP 5\n\x00"))
				for i = 1; i <= DoF; i++ {
					if r[i] != 0 {
						dD[i] = float64(Dp[lc][i])
					}
				}
				noarch.Printf([]byte("STEP 6\n\x00"))
				solve_system(K, dD, F_mech[lc], dR, DoF, q, r, (*[100000000]int)(unsafe.Pointer(&ok))[:], verbose, (*[100000000]float64)(unsafe.Pointer(&rms_resid))[:])
				noarch.Printf([]byte("STEP 7\n\x00"))
				for i = 1; i <= DoF; i++ {
					if q[i] != 0 {
						D[i] += dD[i]
					} else {
						D[i] = float64(Dp[lc][i])
						dD[i] = 0
					}
				}
				for i = 1; i <= DoF; i++ {
					if r[i] != 0 {
						R[i] += dR[i]
					}
				}
				noarch.Printf([]byte("STEP 8\n\x00"))
			}
			for i = 1; i <= DoF; i++ {
				F[i] = F_temp[lc][i] + F_mech[lc][i]
			}
			element_end_forces(Q, nE, xyz, L, Le, N1, N2, Ax, Asy, Asz, Jx, Iy, Iz, E, G, p, eqF_temp[lc], eqF_mech[lc], D, shear, geom, (*[100000000]int)(unsafe.Pointer(&axial_strain_warning))[:])
			noarch.Printf([]byte("STEP 10\n\x00"))
			error = equilibrium_error(dF, F, K, D, DoF, q, r)
			noarch.Printf([]byte("STEP 11\n\x00"))
			if geom != 0 && verbose != 0 {
				noarch.Fprintf(stdout, []byte("\n Non-Linear Elastic Analysis ...\n\x00"))
			}
			noarch.Printf([]byte("STEP 12\n\x00"))
			if geom != 0 {
				error = 1
				ok = 0
				iter = 0
			}
			for geom != 0 && error > tol && iter < 500 && ok >= 0 {
				iter++
				assemble_K(K, DoF, nE, xyz, rj, L, Le, N1, N2, Ax, Asy, Asz, Jx, Iy, Iz, E, G, p, shear, geom, Q, debug)
				error = equilibrium_error(dF, F, K, D, DoF, q, r)
				solve_system(K, dD, dF, dR, DoF, q, r, (*[100000000]int)(unsafe.Pointer(&ok))[:], verbose, (*[100000000]float64)(unsafe.Pointer(&rms_resid))[:])
				if ok < 0 {
					noarch.Fprintf(stderr, []byte("   The stiffness matrix is not pos-def. \n\x00"))
					noarch.Fprintf(stderr, []byte("   Reduce loads and re-run the analysis.\n\x00"))
					ExitCode = 181
					break
				}
				for i = 1; i <= DoF; i++ {
					if q[i] != 0 {
						D[i] += dD[i]
					}
				}
				element_end_forces(Q, nE, xyz, L, Le, N1, N2, Ax, Asy, Asz, Jx, Iy, Iz, E, G, p, eqF_temp[lc], eqF_mech[lc], D, shear, geom, (*[100000000]int)(unsafe.Pointer(&axial_strain_warning))[:])
				if verbose != 0 {
					noarch.Fprintf(stdout, []byte("   NR iteration %3d ---\x00"), iter)
					noarch.Fprintf(stdout, []byte(" RMS relative equilibrium error = %8.2e \n\x00"), error)
				}
			}
			if axial_strain_warning > 0 && ExitCode == 0 {
				ExitCode = 182
			}
			if axial_strain_warning > 0 && ExitCode == 181 {
				ExitCode = 183
			}
			if geom != 0 {
				compute_reaction_forces(R, F, K, D, DoF, r)
			}
			if write_matrix != 0 {
				save_ut_dmatrix([]byte("Ks\x00"), K, DoF, []byte("w\x00"))
			}
			if verbose != 0 && ok >= 0 {
				evaluate(float32(error), float32(rms_resid), float32(tol), geom)
			}
			write_static_results(fp, nN, nE, nL, lc, DoF, N1, N2, F, D, R, r, Q, rms_resid, ok, axial_sign)
			if filetype == 1 {
				write_static_csv(OUT_file, title, nN, nE, nL, lc, DoF, N1, N2, F, D, R, r, Q, error, ok)
			}
			if filetype == 2 {
				write_static_mfile(OUT_file, title, nN, nE, nL, lc, DoF, N1, N2, F, D, R, r, Q, error, ok)
			}
			write_internal_forces(OUT_file, fp, infcpath, lc, nL, title, dx, xyz, Q, nN, nE, L, N1, N2, Ax, Asy, Asz, Jx, Iy, Iz, E, G, p, d, gX[lc], gY[lc], gZ[lc], nU[lc], U[lc], nW[lc], W[lc], nP[lc], P[lc], D, shear, error)
			static_mesh(IN_file, infcpath, meshpath, plotpath, title, nN, nE, nL, lc, DoF, xyz, L, N1, N2, p, D, exagg_static, D3_flag, anlyz, dx, scale)
		}
	} else {
		if verbose != 0 {
			noarch.Fprintf(stdout, []byte("\n * %s *\n\x00"), title)
			noarch.Fprintf(stdout, []byte("  DATA CHECK ONLY.\n\x00"))
		}
		static_mesh(IN_file, infcpath, meshpath, plotpath, title, nN, nE, nL, lc, DoF, xyz, L, N1, N2, p, D, exagg_static, D3_flag, anlyz, dx, scale)
	}
	if nM > 0 {
		if verbose&anlyz != 0 {
			noarch.Fprintf(stdout, []byte("\n\n Modal Analysis ...\n\x00"))
		}
		nM_calc = func() int {
			if nM+8 < 2*nM {
				return nM + 8
			}
			return 2 * nM
		}()
		M = dmatrix(1, int32(DoF), 1, int32(DoF))
		f = dvector(1, int32(nM_calc))
		V = dmatrix(1, int32(DoF), 1, int32(nM_calc))
		assemble_M(M, DoF, nN, nE, xyz, rj, L, N1, N2, Ax, Jx, Iy, Iz, p, d, EMs, NMs, NMx, NMy, NMz, lump, debug)
		for j = 1; j <= DoF; j++ {
			if noarch.NotInt(r[j]) != 0 {
				traceK += K[j][j]
				traceM += M[j][j]
			}
		}
		for i = 1; i <= DoF; i++ {
			if r[i] != 0 {
				K[i][i] = traceK * 10000
				M[i][i] = traceM
				for j = i + 1; j <= DoF; j++ {
					M[i][j] = 0
					M[j][i] = M[i][j]
					K[i][j] = M[j][i]
					K[j][i] = K[i][j]
				}
			}
		}
		if write_matrix != 0 {
			save_ut_dmatrix([]byte("Kd\x00"), K, DoF, []byte("w\x00"))
			save_ut_dmatrix([]byte("Md\x00"), M, DoF, []byte("w\x00"))
		}
		if anlyz != 0 {
			if Mmethod == 1 {
				subspace(K, M, DoF, nM_calc, f, V, tol, shift, (*[100000000]int)(unsafe.Pointer(&iter))[:], (*[100000000]int)(unsafe.Pointer(&ok))[:], verbose)
			}
			if Mmethod == 2 {
				stodola(K, M, DoF, nM_calc, f, V, tol, shift, (*[100000000]int)(unsafe.Pointer(&iter))[:], (*[100000000]int)(unsafe.Pointer(&ok))[:], verbose)
			}
			for j = 1; j <= nM_calc; j++ {
				f[j] = math.Sqrt(f[j]) / (2 * 3.141592653589793)
			}
			write_modal_results(fp, nN, nE, nI, DoF, M, f, V, total_mass, struct_mass, iter, sumR, nM, shift, lump, tol, ok)
		}
	}
	noarch.Fprintf(fp, []byte("\f\n\x00"))
	noarch.Fclose(fp)
	if nM > 0 && anlyz != 0 {
		modal_mesh(IN_file, meshpath, modepath, plotpath, title, nN, nE, DoF, nM, xyz, L, N1, N2, p, M, f, V, exagg_modal, D3_flag, anlyz)
		animate(IN_file, meshpath, modepath, plotpath, title, anim, nN, nE, DoF, nM, xyz, L, p, N1, N2, f, V, exagg_modal, D3_flag, pan, scale)
	}
	if nC > 0 {
		if verbose != 0 {
			noarch.Fprintf(stdout, []byte("\n Matrix Condensation ...\n\x00"))
		}
		if Cdof > nM && Cmethod == 3 {
			noarch.Fprintf(stderr, []byte("  Cdof > nM ... Cdof = %d  nM = %d \n\x00"), Cdof, nM)
			noarch.Fprintf(stderr, []byte("  The number of condensed degrees of freedom\x00"))
			noarch.Fprintf(stderr, []byte(" may not exceed the number of computed modes\x00"))
			noarch.Fprintf(stderr, []byte(" when using dynamic condensation.\n\x00"))
			os.Exit(94)
		}
		Kc = dmatrix(1, int32(Cdof), 1, int32(Cdof))
		Mc = dmatrix(1, int32(Cdof), 1, int32(Cdof))
		if m[1] > 0 && nM > 0 {
			Cfreq = f[m[1]]
		}
		if Cmethod == 1 && anlyz != 0 {
			static_condensation(K, DoF, c, Cdof, Kc, 0)
			if verbose != 0 {
				noarch.Fprintf(stdout, []byte("   static condensation of K complete\n\x00"))
			}
		}
		if Cmethod == 2 && anlyz != 0 {
			paz_condensation(M, K, DoF, c, Cdof, Mc, Kc, Cfreq, 0)
			if verbose != 0 {
				noarch.Fprintf(stdout, []byte("   Paz condensation of K and M complete\x00"))
				noarch.Fprintf(stdout, []byte(" ... dynamics matched at %f Hz.\n\x00"), Cfreq)
			}
		}
		if Cmethod == 3 && nM > 0 && anlyz != 0 {
			modal_condensation(M, K, DoF, r, c, Cdof, Mc, Kc, V, f, m, 0)
			if verbose != 0 {
				noarch.Fprintf(stdout, []byte("   modal condensation of K and M complete\n\x00"))
			}
		}
		save_dmatrix([]byte("Kc\x00"), Kc, 1, Cdof, 1, Cdof, 0, []byte("w\x00"))
		save_dmatrix([]byte("Mc\x00"), Mc, 1, Cdof, 1, Cdof, 0, []byte("w\x00"))
		free_dmatrix(Kc, 1, int32(Cdof), 1, int32(Cdof))
		free_dmatrix(Mc, 1, int32(Cdof), 1, int32(Cdof))
	}
	deallocate(nN, nE, nL, nF, nU, nW, nP, nT, DoF, nM, xyz, rj, L, Le, N1, N2, q, r, Ax, Asy, Asz, Jx, Iy, Iz, E, G, p, U, W, P, T, Dp, F_mech, F_temp, eqF_mech, eqF_temp, F, dF, K, Q, D, dD, R, dR, d, EMs, NMs, NMx, NMy, NMz, M, f, V, c, m, pkNx, pkVy, pkVz, pkTx, pkMy, pkMz, pkDx, pkDy, pkDz, pkRx, pkSy, pkSz)
	if verbose != 0 {
		noarch.Fprintf(stdout, []byte("\n\x00"))
	}
	if argc == 1 {
		noarch.Fprintf(stderr, []byte(" The Output Data was appended to %s \n\x00"), OUT_file)
		noarch.Fprintf(stderr, []byte(" A Gnuplot script was written to %s \n\x00"), plotpath)
		noarch.Fprintf(stderr, []byte(" Press the 'Enter' key to close.\n\x00"))
	}
	color(0)
	os.Exit((ExitCode))
}

// assemble_K - transpiled function from  $GOPATH/src/github.com/Konstantin8105/History_frame3DD/src/frame3dd.c:87
/*
 This file is part of FRAME3DD:
 Static and dynamic structural analysis of 2D and 3D frames and trusses with
 elastic and geometric stiffness.
 ---------------------------------------------------------------------------
 http://frame3dd.sourceforge.net/
 ---------------------------------------------------------------------------
 Copyright (C) 1992-2014  Henri P. Gavin

    FRAME3DD is free software: you can redistribute it and/or modify
    it under the terms of the GNU General Public License as published by
    the Free Software Foundation, either version 3 of the License, or
    (at your option) any later version.

    FRAME3DD is distributed in the hope that it will be useful,
    but WITHOUT ANY WARRANTY; without even the implied warranty of
    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
    GNU General Public License for more details.

    You should have received a copy of the GNU General Public License
    along with FRAME3DD.  If not, see <http://www.gnu.org/licenses/>.
*/ //
/* NOTE main 'driver' routine is now moved to main.c */ //
/* #define MATRIX_DEBUG */ //
/* forward declarations */ //
/*
 * ASSEMBLE_K  -  assemble global stiffness matrix from individual elements 23feb94
 */ //
/* element stiffness matrix in global coord */ //
/* member-structure DoF index table	*/ //
/* res=0, */ //
/* res = sprintf(stiffness_fn,"k_%03d",i); */ //
//
func assemble_K(K [][]float64, DoF int, nE int, xyz []vec3, r []float32, L []float64, Le []float64, N1 []int, N2 []int, Ax []float32, Asy []float32, Asz []float32, Jx []float32, Iy []float32, Iz []float32, E []float32, G []float32, p []float32, shear int, geom int, Q [][]float64, debug int) {
	var k [][]float64
	var ind [][]int
	var i int
	var j int
	var ii int
	var jj int
	var l int
	var ll int
	var stiffness_fn []byte = make([]byte, 128)
	for i = 1; i <= DoF; i++ {
		for j = 1; j <= DoF; j++ {
			K[i][j] = 0
		}
	}
	k = dmatrix(1, 12, 1, 12)
	ind = imatrix(1, 12, 1, int32(nE))
	for i = 1; i <= nE; i++ {
		ind[1][i] = 6*N1[i] - 5
		ind[7][i] = 6*N2[i] - 5
		ind[2][i] = ind[1][i] + 1
		ind[8][i] = ind[7][i] + 1
		ind[3][i] = ind[1][i] + 2
		ind[9][i] = ind[7][i] + 2
		ind[4][i] = ind[1][i] + 3
		ind[10][i] = ind[7][i] + 3
		ind[5][i] = ind[1][i] + 4
		ind[11][i] = ind[7][i] + 4
		ind[6][i] = ind[1][i] + 5
		ind[12][i] = ind[7][i] + 5
	}
	for i = 1; i <= nE; i++ {
		elastic_K(k, xyz, r, L[i], Le[i], N1[i], N2[i], Ax[i], Asy[i], Asz[i], Jx[i], Iy[i], Iz[i], E[i], G[i], p[i], shear)
		if geom != 0 {
			geometric_K(k, xyz, r, L[i], Le[i], N1[i], N2[i], Ax[i], Asy[i], Asz[i], Jx[i], Iy[i], Iz[i], E[i], G[i], p[i], -Q[i][1], shear)
		}
		if debug != 0 {
			save_dmatrix(stiffness_fn, k, 1, 12, 1, 12, 0, []byte("w\x00"))
		}
		for l = 1; l <= 12; l++ {
			ii = ind[l][i]
			for ll = 1; ll <= 12; ll++ {
				jj = ind[ll][i]
				K[ii][jj] += k[l][ll]
			}
		}
	}
	free_dmatrix(k, 1, 12, 1, 12)
	free_imatrix(ind, 1, 12, 1, int32(nE))
}

// elastic_K - transpiled function from  $GOPATH/src/github.com/Konstantin8105/History_frame3DD/src/frame3dd.c:151
/*
 * ELASTIC_K - space frame elastic stiffness matrix in global coordnates	22oct02
 */ //
/* coord Xformn */ //
/* shear deformatn coefficients	*/ //
/* globalize */ //
/* check and enforce symmetry of elastic element stiffness matrix */ //
//
func elastic_K(k [][]float64, xyz []vec3, r []float32, L float64, Le float64, n1 int, n2 int, Ax float32, Asy float32, Asz float32, J float32, Iy float32, Iz float32, E float32, G float32, p float32, shear int) {
	var t1 float64
	var t2 float64
	var t3 float64
	var t4 float64
	var t5 float64
	var t6 float64
	var t7 float64
	var t8 float64
	var t9 float64
	var Ksy float64
	var Ksz float64
	var i int
	var j int
	coord_trans(xyz, L, n1, n2, (*[100000000]float64)(unsafe.Pointer(&t1))[:], (*[100000000]float64)(unsafe.Pointer(&t2))[:], (*[100000000]float64)(unsafe.Pointer(&t3))[:], (*[100000000]float64)(unsafe.Pointer(&t4))[:], (*[100000000]float64)(unsafe.Pointer(&t5))[:], (*[100000000]float64)(unsafe.Pointer(&t6))[:], (*[100000000]float64)(unsafe.Pointer(&t7))[:], (*[100000000]float64)(unsafe.Pointer(&t8))[:], (*[100000000]float64)(unsafe.Pointer(&t9))[:], p)
	for i = 1; i <= 12; i++ {
		for j = 1; j <= 12; j++ {
			k[i][j] = 0
		}
	}
	if shear != 0 {
		Ksy = 12 * float64(E) * float64(Iz) / (float64(G*Asy) * Le * Le)
		Ksz = 12 * float64(E) * float64(Iy) / (float64(G*Asz) * Le * Le)
	} else {
		Ksz = 0
		Ksy = Ksz
	}
	k[7][7] = float64(E*Ax) / Le
	k[1][1] = k[7][7]
	k[8][8] = 12 * float64(E) * float64(Iz) / (Le * Le * Le * (1 + Ksy))
	k[2][2] = k[8][8]
	k[9][9] = 12 * float64(E) * float64(Iy) / (Le * Le * Le * (1 + Ksz))
	k[3][3] = k[9][9]
	k[10][10] = float64(G*J) / Le
	k[4][4] = k[10][10]
	k[11][11] = (4 + Ksz) * float64(E) * float64(Iy) / (Le * (1 + Ksz))
	k[5][5] = k[11][11]
	k[12][12] = (4 + Ksy) * float64(E) * float64(Iz) / (Le * (1 + Ksy))
	k[6][6] = k[12][12]
	k[3][5] = -6 * float64(E) * float64(Iy) / (Le * Le * (1 + Ksz))
	k[5][3] = k[3][5]
	k[2][6] = 6 * float64(E) * float64(Iz) / (Le * Le * (1 + Ksy))
	k[6][2] = k[2][6]
	k[1][7] = -k[1][1]
	k[7][1] = k[1][7]
	k[6][8] = -k[6][2]
	k[8][6] = k[6][8]
	k[8][12] = k[8][6]
	k[12][8] = k[8][12]
	k[5][9] = -k[5][3]
	k[9][5] = k[5][9]
	k[9][11] = k[9][5]
	k[11][9] = k[9][11]
	k[4][10] = -k[4][4]
	k[10][4] = k[4][10]
	k[3][11] = k[5][3]
	k[11][3] = k[3][11]
	k[2][12] = k[6][2]
	k[12][2] = k[2][12]
	k[2][8] = -k[2][2]
	k[8][2] = k[2][8]
	k[3][9] = -k[3][3]
	k[9][3] = k[3][9]
	k[5][11] = (2 - Ksz) * float64(E) * float64(Iy) / (Le * (1 + Ksz))
	k[11][5] = k[5][11]
	k[6][12] = (2 - Ksy) * float64(E) * float64(Iz) / (Le * (1 + Ksy))
	k[12][6] = k[6][12]
	atma(t1, t2, t3, t4, t5, t6, t7, t8, t9, k, r[n1], r[n2])
	for i = 1; i <= 12; i++ {
		for j = i + 1; j <= 12; j++ {
			if k[i][j] != k[j][i] {
				if math.Abs(k[i][j]/k[j][i]-1) > 1e-06 && (math.Abs(k[i][j]/k[i][i]) > 1e-06 || math.Abs(k[j][i]/k[i][i]) > 1e-06) {
					noarch.Fprintf(stderr, []byte("elastic_K: element stiffness matrix not symetric ...\n\x00"))
					noarch.Fprintf(stderr, []byte(" ... k[%d][%d] = %15.6e \n\x00"), i, j, k[i][j])
					noarch.Fprintf(stderr, []byte(" ... k[%d][%d] = %15.6e   \x00"), j, i, k[j][i])
					noarch.Fprintf(stderr, []byte(" ... relative error = %e \n\x00"), math.Abs(k[i][j]/k[j][i]-1))
					noarch.Fprintf(stderr, []byte(" ... element matrix saved in file 'kt'\n\x00"))
					save_dmatrix([]byte("kt\x00"), k, 1, 12, 1, 12, 0, []byte("w\x00"))
				}
				k[j][i] = 0.5 * (k[i][j] + k[j][i])
				k[i][j] = k[j][i]
			}
		}
	}
}

// geometric_K - transpiled function from  $GOPATH/src/github.com/Konstantin8105/History_frame3DD/src/frame3dd.c:232
/*
 * GEOMETRIC_K - space frame geometric stiffness matrix, global coordnates 20dec07
 */ //
/* coord Xformn */ //
/* shear deformation coefficients	*/ //
// T/L;
// -T/L;
/* globalize */ //
/* check and enforce symmetry of geometric element stiffness matrix */ //
/* add geometric stiffness matrix to elastic stiffness matrix ... */ //
//
func geometric_K(k [][]float64, xyz []vec3, r []float32, L float64, Le float64, n1 int, n2 int, Ax float32, Asy float32, Asz float32, J float32, Iy float32, Iz float32, E float32, G float32, p float32, T float64, shear int) {
	var t1 float64
	var t2 float64
	var t3 float64
	var t4 float64
	var t5 float64
	var t6 float64
	var t7 float64
	var t8 float64
	var t9 float64
	var kg [][]float64
	var Ksy float64
	var Ksz float64
	var Dsy float64
	var Dsz float64
	var i int
	var j int
	coord_trans(xyz, L, n1, n2, (*[100000000]float64)(unsafe.Pointer(&t1))[:], (*[100000000]float64)(unsafe.Pointer(&t2))[:], (*[100000000]float64)(unsafe.Pointer(&t3))[:], (*[100000000]float64)(unsafe.Pointer(&t4))[:], (*[100000000]float64)(unsafe.Pointer(&t5))[:], (*[100000000]float64)(unsafe.Pointer(&t6))[:], (*[100000000]float64)(unsafe.Pointer(&t7))[:], (*[100000000]float64)(unsafe.Pointer(&t8))[:], (*[100000000]float64)(unsafe.Pointer(&t9))[:], p)
	kg = dmatrix(1, 12, 1, 12)
	for i = 1; i <= 12; i++ {
		for j = 1; j <= 12; j++ {
			kg[i][j] = 0
		}
	}
	if shear != 0 {
		Ksy = 12 * float64(E) * float64(Iz) / (float64(G*Asy) * Le * Le)
		Ksz = 12 * float64(E) * float64(Iy) / (float64(G*Asz) * Le * Le)
		Dsy = (1 + Ksy) * (1 + Ksy)
		Dsz = (1 + Ksz) * (1 + Ksz)
	} else {
		Ksz = 0
		Ksy = Ksz
		Dsz = 1
		Dsy = Dsz
	}
	kg[7][7] = 0
	kg[1][1] = kg[7][7]
	kg[8][8] = T / L * (1.2 + 2*Ksy + Ksy*Ksy) / Dsy
	kg[2][2] = kg[8][8]
	kg[9][9] = T / L * (1.2 + 2*Ksz + Ksz*Ksz) / Dsz
	kg[3][3] = kg[9][9]
	kg[10][10] = T / L * float64(J) / float64(Ax)
	kg[4][4] = kg[10][10]
	kg[11][11] = T * L * (2/15 + Ksz/6 + Ksz*Ksz/12) / Dsz
	kg[5][5] = kg[11][11]
	kg[12][12] = T * L * (2/15 + Ksy/6 + Ksy*Ksy/12) / Dsy
	kg[6][6] = kg[12][12]
	kg[7][1] = 0
	kg[1][7] = kg[7][1]
	kg[3][11] = -T / 10 / Dsz
	kg[11][3] = kg[3][11]
	kg[3][5] = kg[11][3]
	kg[5][3] = kg[3][5]
	kg[9][11] = T / 10 / Dsz
	kg[11][9] = kg[9][11]
	kg[5][9] = kg[11][9]
	kg[9][5] = kg[5][9]
	kg[2][12] = T / 10 / Dsy
	kg[12][2] = kg[2][12]
	kg[2][6] = kg[12][2]
	kg[6][2] = kg[2][6]
	kg[8][12] = -T / 10 / Dsy
	kg[12][8] = kg[8][12]
	kg[6][8] = kg[12][8]
	kg[8][6] = kg[6][8]
	kg[10][4] = -kg[4][4]
	kg[4][10] = kg[10][4]
	kg[2][8] = -T / L * (1.2 + 2*Ksy + Ksy*Ksy) / Dsy
	kg[8][2] = kg[2][8]
	kg[3][9] = -T / L * (1.2 + 2*Ksz + Ksz*Ksz) / Dsz
	kg[9][3] = kg[3][9]
	kg[5][11] = -T * L * (1/30 + Ksz/6 + Ksz*Ksz/12) / Dsz
	kg[11][5] = kg[5][11]
	kg[6][12] = -T * L * (1/30 + Ksy/6 + Ksy*Ksy/12) / Dsy
	kg[12][6] = kg[6][12]
	atma(t1, t2, t3, t4, t5, t6, t7, t8, t9, kg, r[n1], r[n2])
	for i = 1; i <= 12; i++ {
		for j = i + 1; j <= 12; j++ {
			if kg[i][j] != kg[j][i] {
				if math.Abs(kg[i][j]/kg[j][i]-1) > 1e-06 && (math.Abs(kg[i][j]/kg[i][i]) > 1e-06 || math.Abs(kg[j][i]/kg[i][i]) > 1e-06) {
					noarch.Fprintf(stderr, []byte("geometric_K element stiffness matrix not symetric ...\n\x00"))
					noarch.Fprintf(stderr, []byte(" ... kg[%d][%d] = %15.6e \n\x00"), i, j, kg[i][j])
					noarch.Fprintf(stderr, []byte(" ... kg[%d][%d] = %15.6e   \x00"), j, i, kg[j][i])
					noarch.Fprintf(stderr, []byte(" ... relative error = %e \n\x00"), math.Abs(kg[i][j]/kg[j][i]-1))
					noarch.Fprintf(stderr, []byte(" ... element matrix saved in file 'kg'\n\x00"))
					save_dmatrix([]byte("kg\x00"), kg, 1, 12, 1, 12, 0, []byte("w\x00"))
				}
				kg[j][i] = 0.5 * (kg[i][j] + kg[j][i])
				kg[i][j] = kg[j][i]
			}
		}
	}
	for i = 1; i <= 12; i++ {
		for j = 1; j <= 12; j++ {
			k[i][j] += kg[i][j]
		}
	}
	free_dmatrix(kg, 1, 12, 1, 12)
}

// solve_system - transpiled function from  $GOPATH/src/github.com/Konstantin8105/History_frame3DD/src/frame3dd.c:352
/*
 * SOLVE_SYSTEM  -  solve {F} =   [K]{D} via L D L' decomposition        27dec01
 * Prescribed displacements are "mechanical loads" not "temperature loads"
 */ //
/* diagonal vector of the L D L' decomp. */ //
/* suppress verbose output		*/ //
/*  L D L' decomposition of K[q,q] into lower triangle of K[q,q] and diag[q] */ //
/*  vectors F and D are unchanged */ //
/* exit(31); */ //
/* LDL'  back-substitution for D[q] and R[r] */ //
/* improve solution for D[q] and R[r] */ //
//
func solve_system(K [][]float64, D []float64, F []float64, R []float64, DoF int, q []int, r []int, ok []int, verbose int, rms_resid []float64) {
	var diag []float64
	verbose = 0
	diag = dvector(1, int32(DoF))
	ldl_dcmp_pm(K, DoF, diag, F, D, R, q, r, 1, 0, ok)
	if ok[0] < 0 {
		noarch.Fprintf(stderr, []byte(" Make sure that all six\x00"))
		noarch.Fprintf(stderr, []byte(" rigid body translations are restrained!\n\x00"))
	} else {
		ldl_dcmp_pm(K, DoF, diag, F, D, R, q, r, 0, 1, ok)
		if verbose != 0 {
			noarch.Fprintf(stdout, []byte("    LDL' RMS residual:\x00"))
		}
		rms_resid[0] = float64((func() int {
			ok[0] = 1
			return ok[0]
		}()))
		for {
			ldl_mprove_pm(K, DoF, diag, F, D, R, q, r, rms_resid, ok)
			if verbose != 0 {
				noarch.Fprintf(stdout, []byte("%9.2e\x00"), rms_resid[0])
			}
			if noarch.NotInt((ok[0])) != 0 {
				break
			}
		}
		if verbose != 0 {
			noarch.Fprintf(stdout, []byte("\n\x00"))
		}
	}
	free_dvector(diag, 1, int32(DoF))
}

// equilibrium_error - transpiled function from  $GOPATH/src/github.com/Konstantin8105/History_frame3DD/src/frame3dd.c:390
/*
 * EQUILIBRIUM_ERROR -  compute {dF_q} =   {F_q} - [K_qq]{D_q} - [K_qr]{D_r}
 * use only the upper-triangle of [K_qq]
 * return ||dF||/||F||
 * 2014-05-16
 */ //
//  sum of squares of dF
//  sum of squares of F
// compute equilibrium error at free coord's (q)
// K_qq in upper triangle only
// convergence criterion
//
func equilibrium_error(dF []float64, F []float64, K [][]float64, D []float64, DoF int, q []int, r []int) float64 {
	var ss_dF float64
	var ss_F float64
	var errF float64
	var i int
	var j int
	for i = 1; i <= DoF; i++ {
		errF = 0
		if q[i] != 0 {
			errF = F[i]
			for j = 1; j <= DoF; j++ {
				if q[j] != 0 {
					if i <= j {
						errF -= K[i][j] * D[j]
					} else {
						errF -= K[j][i] * D[j]
					}
				}
			}
			for j = 1; j <= DoF; j++ {
				if r[j] != 0 {
					errF -= K[i][j] * D[j]
				}
			}
		}
		dF[i] = errF
	}
	for i = 1; i <= DoF; i++ {
		if q[i] != 0 {
			ss_dF += dF[i] * dF[i]
		}
	}
	for i = 1; i <= DoF; i++ {
		if q[i] != 0 {
			ss_F += F[i] * F[i]
		}
	}
	return (math.Sqrt(ss_dF) / math.Sqrt(ss_F))
}

// element_end_forces - transpiled function from  $GOPATH/src/github.com/Konstantin8105/History_frame3DD/src/frame3dd.c:425
/*
 * ELEMENT_END_FORCES  -  evaluate the end forces for all elements
 * 23feb94
 */ //
// equivalent element end forces from temp loads
// equivalent element end forces from mech loads
//
func element_end_forces(Q [][]float64, nE int, xyz []vec3, L []float64, Le []float64, N1 []int, N2 []int, Ax []float32, Asy []float32, Asz []float32, Jx []float32, Iy []float32, Iz []float32, E []float32, G []float32, p []float32, eqF_temp [][]float64, eqF_mech [][]float64, D []float64, shear int, geom int, axial_strain_warning []int) {
	var s []float64
	var axial_strain float64
	var m int
	var j int
	s = dvector(1, 12)
	axial_strain_warning[0] = 0
	for m = 1; m <= nE; m++ {
		frame_element_force(s, xyz, L[m], Le[m], N1[m], N2[m], Ax[m], Asy[m], Asz[m], Jx[m], Iy[m], Iz[m], E[m], G[m], p[m], eqF_temp[m], eqF_mech[m], D, shear, geom, (*[100000000]float64)(unsafe.Pointer(&axial_strain))[:])
		for j = 1; j <= 12; j++ {
			Q[m][j] = s[j]
		}
		if math.Abs(axial_strain) > 0.001 {
			noarch.Fprintf(stderr, []byte(" Warning! Frame element %2d has an average axial strain of %8.6f\n\x00"), m, axial_strain)
			axial_strain_warning[0]++
		}
	}
	free_dvector(s, 1, 12)
}

// frame_element_force - transpiled function from  $GOPATH/src/github.com/Konstantin8105/History_frame3DD/src/frame3dd.c:464
/*
 * FRAME_ELEMENT_FORCE  -  evaluate the end forces in local coord's
 * 12nov02
 */ //
/* coord Xformn	*/ //
// x1, y1, z1, x2, y2, z2,	/* node coordinates	*/
//  Ls,			/* stretched length of element */
/* stretch in the frame element */ //
/* shear deformation coeff's	*/ //
/* axial force for geometric stiffness */ //
/* finite strain ... (not consistent with 2nd order formulation) */ //
/*
 *	delta += ( pow(((d7-d1)*t4 + (d8-d2)*t5 + (d9-d3)*t6),2.0) +
 *		   pow(((d7-d1)*t7 + (d8-d2)*t8 + (d9-d3)*t9),2.0) )/(2.0*L);
 */ //
/* true strain ... (not appropriate for structural materials) */ //
/*
 *	x1 = xyz[n1].x;	y1 = xyz[n1].y;	z1 = xyz[n1].z;
 *	x2 = xyz[n2].x;	y2 = xyz[n2].y;	z2 = xyz[n2].z;
 *
 *  	Ls =	pow((x2+d7-x1-d1),2.0) +
 *		pow((y2+d8-y1-d2),2.0) +
 *		pow((z2+d9-z1-d3),2.0);
 *	Ls = sqrt(Ls) + Le - L;
 *
 *	delta = Le*log(Ls/Le);
 */ //
/* axial element displacement ... */ //
// log(Ls/Le);
// add fixed end forces to internal element forces
// 18oct2012, 14may1204, 15may2014
// add temperature fixed-end-forces to variables f1-f12
// add mechanical load fixed-end-forces to variables f1-f12
// f1 ...  f12 are in the global element coordinate system
// transform f1 ... f12 to local element coordinate system and
// add local fixed end forces (-equivalent loads) to internal loads
// {Q} = [T]{f}
//
func frame_element_force(s []float64, xyz []vec3, L float64, Le float64, n1 int, n2 int, Ax float32, Asy float32, Asz float32, J float32, Iy float32, Iz float32, E float32, G float32, p float32, f_t []float64, f_m []float64, D []float64, shear int, geom int, axial_strain []float64) {
	var t1 float64
	var t2 float64
	var t3 float64
	var t4 float64
	var t5 float64
	var t6 float64
	var t7 float64
	var t8 float64
	var t9 float64
	var d1 float64
	var d2 float64
	var d3 float64
	var d4 float64
	var d5 float64
	var d6 float64
	var d7 float64
	var d8 float64
	var d9 float64
	var d10 float64
	var d11 float64
	var d12 float64
	var delta float64
	var Ksy float64
	var Ksz float64
	var Dsy float64
	var Dsz float64
	var T float64
	var f1 float64
	var f2 float64
	var f3 float64
	var f4 float64
	var f5 float64
	var f6 float64
	var f7 float64
	var f8 float64
	var f9 float64
	var f10 float64
	var f11 float64
	var f12 float64
	coord_trans(xyz, L, n1, n2, (*[100000000]float64)(unsafe.Pointer(&t1))[:], (*[100000000]float64)(unsafe.Pointer(&t2))[:], (*[100000000]float64)(unsafe.Pointer(&t3))[:], (*[100000000]float64)(unsafe.Pointer(&t4))[:], (*[100000000]float64)(unsafe.Pointer(&t5))[:], (*[100000000]float64)(unsafe.Pointer(&t6))[:], (*[100000000]float64)(unsafe.Pointer(&t7))[:], (*[100000000]float64)(unsafe.Pointer(&t8))[:], (*[100000000]float64)(unsafe.Pointer(&t9))[:], p)
	n1 = 6 * (n1 - 1)
	n2 = 6 * (n2 - 1)
	d1 = D[n1+1]
	d2 = D[n1+2]
	d3 = D[n1+3]
	d4 = D[n1+4]
	d5 = D[n1+5]
	d6 = D[n1+6]
	d7 = D[n2+1]
	d8 = D[n2+2]
	d9 = D[n2+3]
	d10 = D[n2+4]
	d11 = D[n2+5]
	d12 = D[n2+6]
	if shear != 0 {
		Ksy = 12 * float64(E) * float64(Iz) / (float64(G*Asy) * Le * Le)
		Ksz = 12 * float64(E) * float64(Iy) / (float64(G*Asz) * Le * Le)
		Dsy = (1 + Ksy) * (1 + Ksy)
		Dsz = (1 + Ksz) * (1 + Ksz)
	} else {
		Ksz = 0
		Ksy = Ksz
		Dsz = 1
		Dsy = Dsz
	}
	delta = (d7-d1)*t1 + (d8-d2)*t2 + (d9-d3)*t3
	axial_strain[0] = delta / Le
	s[1] = -(float64(Ax*E) / Le) * ((d7-d1)*t1 + (d8-d2)*t2 + (d9-d3)*t3)
	if geom != 0 {
		T = -s[1]
	}
	s[2] = -(12*float64(E)*float64(Iz)/(Le*Le*Le*(1+Ksy))+T/L*(1.2+2*Ksy+Ksy*Ksy)/Dsy)*((d7-d1)*t4+(d8-d2)*t5+(d9-d3)*t6) + (6*float64(E)*float64(Iz)/(Le*Le*(1+Ksy))+T/10/Dsy)*((d4+d10)*t7+(d5+d11)*t8+(d6+d12)*t9)
	s[3] = -(12*float64(E)*float64(Iy)/(Le*Le*Le*(1+Ksz))+T/L*(1.2+2*Ksz+Ksz*Ksz)/Dsz)*((d7-d1)*t7+(d8-d2)*t8+(d9-d3)*t9) - (6*float64(E)*float64(Iy)/(Le*Le*(1+Ksz))+T/10/Dsz)*((d4+d10)*t4+(d5+d11)*t5+(d6+d12)*t6)
	s[4] = -(float64(G*J) / Le) * ((d10-d4)*t1 + (d11-d5)*t2 + (d12-d6)*t3)
	s[5] = (6*float64(E)*float64(Iy)/(Le*Le*(1+Ksz))+T/10/Dsz)*((d7-d1)*t7+(d8-d2)*t8+(d9-d3)*t9) + ((4+Ksz)*float64(E)*float64(Iy)/(Le*(1+Ksz))+T*L*(2/15+Ksz/6+Ksz*Ksz/12)/Dsz)*(d4*t4+d5*t5+d6*t6) + ((2-Ksz)*float64(E)*float64(Iy)/(Le*(1+Ksz))-T*L*(1/30+Ksz/6+Ksz*Ksz/12)/Dsz)*(d10*t4+d11*t5+d12*t6)
	s[6] = -(6*float64(E)*float64(Iz)/(Le*Le*(1+Ksy))+T/10/Dsy)*((d7-d1)*t4+(d8-d2)*t5+(d9-d3)*t6) + ((4+Ksy)*float64(E)*float64(Iz)/(Le*(1+Ksy))+T*L*(2/15+Ksy/6+Ksy*Ksy/12)/Dsy)*(d4*t7+d5*t8+d6*t9) + ((2-Ksy)*float64(E)*float64(Iz)/(Le*(1+Ksy))-T*L*(1/30+Ksy/6+Ksy*Ksy/12)/Dsy)*(d10*t7+d11*t8+d12*t9)
	s[7] = -s[1]
	s[8] = -s[2]
	s[9] = -s[3]
	s[10] = -s[4]
	s[11] = (6*float64(E)*float64(Iy)/(Le*Le*(1+Ksz))+T/10/Dsz)*((d7-d1)*t7+(d8-d2)*t8+(d9-d3)*t9) + ((4+Ksz)*float64(E)*float64(Iy)/(Le*(1+Ksz))+T*L*(2/15+Ksz/6+Ksz*Ksz/12)/Dsz)*(d10*t4+d11*t5+d12*t6) + ((2-Ksz)*float64(E)*float64(Iy)/(Le*(1+Ksz))-T*L*(1/30+Ksz/6+Ksz*Ksz/12)/Dsz)*(d4*t4+d5*t5+d6*t6)
	s[12] = -(6*float64(E)*float64(Iz)/(Le*Le*(1+Ksy))+T/10/Dsy)*((d7-d1)*t4+(d8-d2)*t5+(d9-d3)*t6) + ((4+Ksy)*float64(E)*float64(Iz)/(Le*(1+Ksy))+T*L*(2/15+Ksy/6+Ksy*Ksy/12)/Dsy)*(d10*t7+d11*t8+d12*t9) + ((2-Ksy)*float64(E)*float64(Iz)/(Le*(1+Ksy))-T*L*(1/30+Ksy/6+Ksy*Ksy/12)/Dsy)*(d4*t7+d5*t8+d6*t9)
	f1 = f_t[1] + f_m[1]
	f2 = f_t[2] + f_m[2]
	f3 = f_t[3] + f_m[3]
	f4 = f_t[4] + f_m[4]
	f5 = f_t[5] + f_m[5]
	f6 = f_t[6] + f_m[6]
	f7 = f_t[7] + f_m[7]
	f8 = f_t[8] + f_m[8]
	f9 = f_t[9] + f_m[9]
	f10 = f_t[10] + f_m[10]
	f11 = f_t[11] + f_m[11]
	f12 = f_t[12] + f_m[12]
	s[1] -= f1*t1 + f2*t2 + f3*t3
	s[2] -= f1*t4 + f2*t5 + f3*t6
	s[3] -= f1*t7 + f2*t8 + f3*t9
	s[4] -= f4*t1 + f5*t2 + f6*t3
	s[5] -= f4*t4 + f5*t5 + f6*t6
	s[6] -= f4*t7 + f5*t8 + f6*t9
	s[7] -= f7*t1 + f8*t2 + f9*t3
	s[8] -= f7*t4 + f8*t5 + f9*t6
	s[9] -= f7*t7 + f8*t8 + f9*t9
	s[10] -= f10*t1 + f11*t2 + f12*t3
	s[11] -= f10*t4 + f11*t5 + f12*t6
	s[12] -= f10*t7 + f11*t8 + f12*t9
}

// compute_reaction_forces - transpiled function from  $GOPATH/src/github.com/Konstantin8105/History_frame3DD/src/frame3dd.c:658
/*
void add_feF(		// DISUSED CODE
	vec3 *xyz,
	double *L, int *N1, int *N2, float *p,
	double **Q,
int nE, int DoF,
	int verbose
){
	double  t1, t2, t3, t4, t5, t6, t7, t8, t9,	// 3D coord Xformn
	int	m, n1, n2, i1, i2; //, J, x;

	for (m=1; m <= nE; m++) {	// loop over all frame elements

		n1 = N1[m];	n2 = N2[m];

 * reaction calculations removed from Frame3DD on 2014-05-14 ...
 * these calculations are now in solve_system()
 *		// add fixed-end forces to reaction forces
 *		for (i=1; i<=6; i++) {
 *			i1 = 6*(n1-1) + i;
 *			if (r[i1])
 *				F[i1] -= ( f_t[m][i] + f_m[m][i] );
 *		}
 *		for (i=1; i<=6; i++) {
 *			i2 = 6*(n2-1) + i;
 *			if (r[i2])
 *				F[i2] -= ( f_t[m][i+6] + f_m[m][i+6] );
 *		}
 *

		coord_trans ( xyz, L[m], n1, n2,
			&t1, &t2, &t3, &t4, &t5, &t6, &t7, &t8, &t9, p[m] );

		// n1 = 6*(n1-1);	n2 = 6*(n2-1);	// ??


	}
}
*/ //
/*
 * COMPUTE_REACTION_FORCES : R(r) = [K(r,q)]*{D(q)} + [K(r,r)]*{D(r)} - F(r)
 * reaction forces satisfy equilibrium in the solved system
 * only really needed for geometric-nonlinear problems
 * 2012-10-12  , 2014-05-16
 */ //
// coordinate "i" is a reaction coord.
// negative of equiv loads at coord i
// reactions are relaxed through system deformations
//
func compute_reaction_forces(R []float64, F []float64, K [][]float64, D []float64, DoF int, r []int) {
	var i int
	var j int
	for i = 1; i <= DoF; i++ {
		R[i] = 0
		if r[i] != 0 {
			R[i] = -F[i]
			for j = 1; j <= DoF; j++ {
				R[i] += K[i][j] * D[j]
			}
		}
	}
}

// assemble_M - transpiled function from  $GOPATH/src/github.com/Konstantin8105/History_frame3DD/src/frame3dd.c:677
/*
 * ASSEMBLE_M  -  assemble global mass matrix from element mass & inertia  24nov98
 */ //
/* element mass matrix in global coord */ //
/* member-structure DoF index table     */ //
/* res=0,  */ //
/* res = sprintf(mass_fn,"m_%03d",i); */ //
// add extra node mass
//
func assemble_M(M [][]float64, DoF int, nN int, nE int, xyz []vec3, r []float32, L []float64, N1 []int, N2 []int, Ax []float32, Jx []float32, Iy []float32, Iz []float32, p []float32, d []float32, EMs []float32, NMs []float32, NMx []float32, NMy []float32, NMz []float32, lump int, debug int) {
	var m [][]float64
	var ind [][]int
	var i int
	var j int
	var ii int
	var jj int
	var l int
	var ll int
	var mass_fn []byte = make([]byte, 128)
	for i = 1; i <= DoF; i++ {
		for j = 1; j <= DoF; j++ {
			M[i][j] = 0
		}
	}
	m = dmatrix(1, 12, 1, 12)
	ind = imatrix(1, 12, 1, int32(nE))
	for i = 1; i <= nE; i++ {
		ind[1][i] = 6*N1[i] - 5
		ind[7][i] = 6*N2[i] - 5
		ind[2][i] = ind[1][i] + 1
		ind[8][i] = ind[7][i] + 1
		ind[3][i] = ind[1][i] + 2
		ind[9][i] = ind[7][i] + 2
		ind[4][i] = ind[1][i] + 3
		ind[10][i] = ind[7][i] + 3
		ind[5][i] = ind[1][i] + 4
		ind[11][i] = ind[7][i] + 4
		ind[6][i] = ind[1][i] + 5
		ind[12][i] = ind[7][i] + 5
	}
	for i = 1; i <= nE; i++ {
		if lump != 0 {
			lumped_M(m, xyz, L[i], N1[i], N2[i], Ax[i], Jx[i], Iy[i], Iz[i], p[i], d[i], EMs[i])
		} else {
			consistent_M(m, xyz, r, L[i], N1[i], N2[i], Ax[i], Jx[i], Iy[i], Iz[i], p[i], d[i], EMs[i])
		}
		if debug != 0 {
			save_dmatrix(mass_fn, m, 1, 12, 1, 12, 0, []byte("w\x00"))
		}
		for l = 1; l <= 12; l++ {
			ii = ind[l][i]
			for ll = 1; ll <= 12; ll++ {
				jj = ind[ll][i]
				M[ii][jj] += m[l][ll]
			}
		}
	}
	for j = 1; j <= nN; j++ {
		i = 6 * (j - 1)
		M[i+1][i+1] += float64(NMs[j])
		M[i+2][i+2] += float64(NMs[j])
		M[i+3][i+3] += float64(NMs[j])
		M[i+4][i+4] += float64(NMx[j])
		M[i+5][i+5] += float64(NMy[j])
		M[i+6][i+6] += float64(NMz[j])
	}
	for i = 1; i <= DoF; i++ {
		if M[i][i] <= 0 {
			noarch.Fprintf(stderr, []byte("  error: Non pos-def mass matrix\n\x00"))
			noarch.Fprintf(stderr, []byte("  M[%d][%d] = %lf\n\x00"), i, i, M[i][i])
		}
	}
	free_dmatrix(m, 1, 12, 1, 12)
	free_imatrix(ind, 1, 12, 1, int32(nE))
}

// lumped_M - transpiled function from  $GOPATH/src/github.com/Konstantin8105/History_frame3DD/src/frame3dd.c:754
/*
 * LUMPED_M  -  space frame element lumped mass matrix in global coordnates 7apr94
 */ //
/* coord Xformn */ //
/* translational, rotational & polar inertia */ //
/* rotatory inertia of extra mass is neglected */ //
/* assumes simple cross-section	*/ //
//
func lumped_M(m [][]float64, xyz []vec3, L float64, n1 int, n2 int, Ax float32, J float32, Iy float32, Iz float32, p float32, d float32, EMs float32) {
	var t1 float64
	var t2 float64
	var t3 float64
	var t4 float64
	var t5 float64
	var t6 float64
	var t7 float64
	var t8 float64
	var t9 float64
	var t float64
	var ry float64
	var rz float64
	var po float64
	var i int
	var j int
	coord_trans(xyz, L, n1, n2, (*[100000000]float64)(unsafe.Pointer(&t1))[:], (*[100000000]float64)(unsafe.Pointer(&t2))[:], (*[100000000]float64)(unsafe.Pointer(&t3))[:], (*[100000000]float64)(unsafe.Pointer(&t4))[:], (*[100000000]float64)(unsafe.Pointer(&t5))[:], (*[100000000]float64)(unsafe.Pointer(&t6))[:], (*[100000000]float64)(unsafe.Pointer(&t7))[:], (*[100000000]float64)(unsafe.Pointer(&t8))[:], (*[100000000]float64)(unsafe.Pointer(&t9))[:], p)
	t = (float64(d*Ax)*L + float64(EMs)) / 2
	ry = float64(d*Iy) * L / 2
	rz = float64(d*Iz) * L / 2
	po = float64(d) * L * float64(J) / 2
	for i = 1; i <= 12; i++ {
		for j = 1; j <= 12; j++ {
			m[i][j] = 0
		}
	}
	m[9][9] = t
	m[8][8] = m[9][9]
	m[7][7] = m[8][8]
	m[3][3] = m[7][7]
	m[2][2] = m[3][3]
	m[1][1] = m[2][2]
	m[10][10] = po*t1*t1 + ry*t4*t4 + rz*t7*t7
	m[4][4] = m[10][10]
	m[11][11] = po*t2*t2 + ry*t5*t5 + rz*t8*t8
	m[5][5] = m[11][11]
	m[12][12] = po*t3*t3 + ry*t6*t6 + rz*t9*t9
	m[6][6] = m[12][12]
	m[11][10] = po*t1*t2 + ry*t4*t5 + rz*t7*t8
	m[10][11] = m[11][10]
	m[5][4] = m[10][11]
	m[4][5] = m[5][4]
	m[12][10] = po*t1*t3 + ry*t4*t6 + rz*t7*t9
	m[10][12] = m[12][10]
	m[6][4] = m[10][12]
	m[4][6] = m[6][4]
	m[12][11] = po*t2*t3 + ry*t5*t6 + rz*t8*t9
	m[11][12] = m[12][11]
	m[6][5] = m[11][12]
	m[5][6] = m[6][5]
}

// consistent_M - transpiled function from  $GOPATH/src/github.com/Konstantin8105/History_frame3DD/src/frame3dd.c:791
/*
 * CONSISTENT_M  -  space frame consistent mass matrix in global coordnates 2oct97
 *		 does not include shear deformations
 */ //
/* coord Xformn */ //
/* translational, rotational & polar inertia */ //
/* rotatory inertia of extra beam mass is neglected */ //
/* globalize */ //
/* check and enforce symmetry of consistent element mass matrix */ //
//
func consistent_M(m [][]float64, xyz []vec3, r []float32, L float64, n1 int, n2 int, Ax float32, J float32, Iy float32, Iz float32, p float32, d float32, EMs float32) {
	var t1 float64
	var t2 float64
	var t3 float64
	var t4 float64
	var t5 float64
	var t6 float64
	var t7 float64
	var t8 float64
	var t9 float64
	var t float64
	var ry float64
	var rz float64
	var po float64
	var i int
	var j int
	coord_trans(xyz, L, n1, n2, (*[100000000]float64)(unsafe.Pointer(&t1))[:], (*[100000000]float64)(unsafe.Pointer(&t2))[:], (*[100000000]float64)(unsafe.Pointer(&t3))[:], (*[100000000]float64)(unsafe.Pointer(&t4))[:], (*[100000000]float64)(unsafe.Pointer(&t5))[:], (*[100000000]float64)(unsafe.Pointer(&t6))[:], (*[100000000]float64)(unsafe.Pointer(&t7))[:], (*[100000000]float64)(unsafe.Pointer(&t8))[:], (*[100000000]float64)(unsafe.Pointer(&t9))[:], p)
	t = float64(d*Ax) * L
	ry = float64(d * Iy)
	rz = float64(d * Iz)
	po = float64(d*J) * L
	for i = 1; i <= 12; i++ {
		for j = 1; j <= 12; j++ {
			m[i][j] = 0
		}
	}
	m[7][7] = t / 3
	m[1][1] = m[7][7]
	m[8][8] = 13*t/35 + 6*rz/(5*L)
	m[2][2] = m[8][8]
	m[9][9] = 13*t/35 + 6*ry/(5*L)
	m[3][3] = m[9][9]
	m[10][10] = po / 3
	m[4][4] = m[10][10]
	m[11][11] = t*L*L/105 + 2*L*ry/15
	m[5][5] = m[11][11]
	m[12][12] = t*L*L/105 + 2*L*rz/15
	m[6][6] = m[12][12]
	m[3][5] = -11*t*L/210 - ry/10
	m[5][3] = m[3][5]
	m[2][6] = 11*t*L/210 + rz/10
	m[6][2] = m[2][6]
	m[1][7] = t / 6
	m[7][1] = m[1][7]
	m[6][8] = 13*t*L/420 - rz/10
	m[8][6] = m[6][8]
	m[5][9] = -13*t*L/420 + ry/10
	m[9][5] = m[5][9]
	m[4][10] = po / 6
	m[10][4] = m[4][10]
	m[3][11] = 13*t*L/420 - ry/10
	m[11][3] = m[3][11]
	m[2][12] = -13*t*L/420 + rz/10
	m[12][2] = m[2][12]
	m[9][11] = 11*t*L/210 + ry/10
	m[11][9] = m[9][11]
	m[8][12] = -11*t*L/210 - rz/10
	m[12][8] = m[8][12]
	m[2][8] = 9*t/70 - 6*rz/(5*L)
	m[8][2] = m[2][8]
	m[3][9] = 9*t/70 - 6*ry/(5*L)
	m[9][3] = m[3][9]
	m[5][11] = -L*L*t/140 - ry*L/30
	m[11][5] = m[5][11]
	m[6][12] = -L*L*t/140 - rz*L/30
	m[12][6] = m[6][12]
	for i = 1; i <= 3; i++ {
		m[i][i] += 0.5 * float64(EMs)
	}
	for i = 7; i <= 9; i++ {
		m[i][i] += 0.5 * float64(EMs)
	}
	atma(t1, t2, t3, t4, t5, t6, t7, t8, t9, m, r[n1], r[n2])
	for i = 1; i <= 12; i++ {
		for j = i + 1; j <= 12; j++ {
			if m[i][j] != m[j][i] {
				if math.Abs(m[i][j]/m[j][i]-1) > 1e-06 && (math.Abs(m[i][j]/m[i][i]) > 1e-06 || math.Abs(m[j][i]/m[i][i]) > 1e-06) {
					noarch.Fprintf(stderr, []byte("consistent_M: element mass matrix not symetric ...\n\x00"))
					noarch.Fprintf(stderr, []byte(" ... m[%d][%d] = %15.6e \n\x00"), i, j, m[i][j])
					noarch.Fprintf(stderr, []byte(" ... m[%d][%d] = %15.6e   \x00"), j, i, m[j][i])
					noarch.Fprintf(stderr, []byte(" ... relative error = %e \n\x00"), math.Abs(m[i][j]/m[j][i]-1))
					noarch.Fprintf(stderr, []byte(" ... element matrix saved in file 'mc'\n\x00"))
					save_dmatrix([]byte("mc\x00"), m, 1, 12, 1, 12, 0, []byte("w\x00"))
				}
				m[j][i] = 0.5 * (m[i][j] + m[j][i])
				m[i][j] = m[j][i]
			}
		}
	}
}

// static_condensation - transpiled function from  $GOPATH/src/github.com/Konstantin8105/History_frame3DD/src/frame3dd.c:879
/*
 * STATIC_CONDENSATION - of stiffness matrix from NxN to nxn    30aug01
 */ //
/* use only upper triangle of A */ //
/* use only upper triangle of A */ //
/* use only upper triangle of A */ //
//
func static_condensation(A [][]float64, N int, c []int, n int, Ac [][]float64, verbose int) {
	var Arr [][]float64
	var Arc [][]float64
	var i int
	var j int
	var k int
	var ri int
	var rj int
	var ci int
	var cj int
	var ok int
	var r []int
	r = ivector(1, int32(N-n))
	Arr = dmatrix(1, int32(N-n), 1, int32(N-n))
	Arc = dmatrix(1, int32(N-n), 1, int32(n))
	k = 1
	for i = 1; i <= N; i++ {
		ok = 1
		for j = 1; j <= n; j++ {
			if c[j] == i {
				ok = 0
				break
			}
		}
		if ok != 0 {
			r[func() int {
				defer func() {
					k++
				}()
				return k
			}()] = i
		}
	}
	for i = 1; i <= N-n; i++ {
		for j = i; j <= N-n; j++ {
			ri = r[i]
			rj = r[j]
			if ri <= rj {
				Arr[i][j] = A[ri][rj]
				Arr[j][i] = Arr[i][j]
			}
		}
	}
	for i = 1; i <= N-n; i++ {
		for j = 1; j <= n; j++ {
			ri = r[i]
			cj = c[j]
			if ri < cj {
				Arc[i][j] = A[ri][cj]
			} else {
				Arc[i][j] = A[cj][ri]
			}
		}
	}
	xtinvAy(Arc, Arr, Arc, N-n, n, Ac, verbose)
	for i = 1; i <= n; i++ {
		for j = i; j <= n; j++ {
			ci = c[i]
			cj = c[j]
			if ci <= cj {
				Ac[i][j] = A[ci][cj] - Ac[i][j]
				Ac[j][i] = Ac[i][j]
			}
		}
	}
	free_ivector(r, 1, int32(N-n))
	free_dmatrix(Arr, 1, int32(N-n), 1, int32(N-n))
	free_dmatrix(Arc, 1, int32(N-n), 1, int32(n))
}

// paz_condensation - transpiled function from  $GOPATH/src/github.com/Konstantin8105/History_frame3DD/src/frame3dd.c:939
/*
 * PAZ_CONDENSATION -   Paz condensation of mass and stiffness matrices 6jun07
 *          Paz M. Dynamic condensation. AIAA J 1984;22(5):724-727.
 */ //
/* inv(Drr) * Drc	*/ //
/* coordinate transformation matrix	*/ //
/* eigen-value ... omega^2 	*/ //
/* find "remaining" (r) degrees of freedom, not "condensed" (c)	*/ //
/* use only upper triangle of K,M */ //
/* use only upper triangle of K,M */ //
/* inv(Drr) * Drc */ //
/* coordinate transformation matrix	*/ //
/* Kc = T' * K * T	*/ //
/* Mc = T' * M * T	*/ //
//
func paz_condensation(M [][]float64, K [][]float64, N int, c []int, n int, Mc [][]float64, Kc [][]float64, w2 float64, verbose int) {
	var Drr [][]float64
	var Drc [][]float64
	var invDrrDrc [][]float64
	var T [][]float64
	var i int
	var j int
	var k int
	var ri int
	var rj int
	var cj int
	var ok int
	var r []int
	func() {
		if M != nil {
		} else {
			linux.AssertFail([]byte("M!=NULL\x00"), []byte("/home/konstantin/go/src/github.com/Konstantin8105/History_frame3DD/src/frame3dd.c\x00"), 949, []byte("void print_number(int *)"))
		}
	}()
	r = ivector(1, int32(N-n))
	Drr = dmatrix(1, int32(N-n), 1, int32(N-n))
	Drc = dmatrix(1, int32(N-n), 1, int32(n))
	invDrrDrc = dmatrix(1, int32(N-n), 1, int32(n))
	T = dmatrix(1, int32(N), 1, int32(n))
	w2 = 4 * 3.141592653589793 * 3.141592653589793 * w2 * w2
	k = 1
	for i = 1; i <= N; i++ {
		ok = 1
		for j = 1; j <= n; j++ {
			if c[j] == i {
				ok = 0
				break
			}
		}
		if ok != 0 {
			r[func() int {
				defer func() {
					k++
				}()
				return k
			}()] = i
		}
	}
	for i = 1; i <= N-n; i++ {
		for j = 1; j <= N-n; j++ {
			ri = r[i]
			rj = r[j]
			if ri <= rj {
				Drr[i][j] = K[ri][rj] - w2*M[ri][rj]
				Drr[j][i] = Drr[i][j]
			} else {
				Drr[i][j] = K[rj][ri] - w2*M[rj][ri]
				Drr[j][i] = Drr[i][j]
			}
		}
	}
	for i = 1; i <= N-n; i++ {
		for j = 1; j <= n; j++ {
			ri = r[i]
			cj = c[j]
			if ri < cj {
				Drc[i][j] = K[ri][cj] - w2*M[ri][cj]
			} else {
				Drc[i][j] = K[cj][ri] - w2*M[cj][ri]
			}
		}
	}
	invAB(Drr, Drc, N-n, n, invDrrDrc, (*[100000000]int)(unsafe.Pointer(&ok))[:], verbose)
	for i = 1; i <= n; i++ {
		for j = 1; j <= n; j++ {
			T[c[i]][j] = 0
		}
		T[c[i]][i] = 1
	}
	for i = 1; i <= N-n; i++ {
		for j = 1; j <= n; j++ {
			T[r[i]][j] = -invDrrDrc[i][j]
		}
	}
	xtAx(K, T, Kc, N, n)
	xtAx(M, T, Mc, N, n)
	free_ivector(r, 1, int32(N-n))
	free_dmatrix(Drr, 1, int32(N-n), 1, int32(N-n))
	free_dmatrix(Drc, 1, int32(N-n), 1, int32(n))
	free_dmatrix(invDrrDrc, 1, int32(N-n), 1, int32(N-n))
	free_dmatrix(T, 1, int32(N-n), 1, int32(n))
}

// modal_condensation - transpiled function from  $GOPATH/src/github.com/Konstantin8105/History_frame3DD/src/frame3dd.c:1019
/*
 * MODAL_CONDENSATION -
 *      dynamic condensation of mass and stiffness matrices    8oct01
 *  	matches the response at a set of frequencies and modes
 * WARNING: Kc and Mc may be ill-conditioned, and xyzsibly non-positive def.
 */ //
/* temporary storage for matrix mult. */ //
/* first n modal vectors at primary DoF's */ //
/* compute inv(P)' * I * inv(P)	*/ //
/* compute inv(P)' * W^2 * inv(P) */ //
//
func modal_condensation(M [][]float64, K [][]float64, N int, R []int, p []int, n int, Mc [][]float64, Kc [][]float64, V [][]float64, f []float64, m []int, verbose int) {
	var P [][]float64
	var invP [][]float64
	var traceM float64
	var traceMc float64
	var Aij float64
	var i int
	var j int
	var k int
	P = dmatrix(1, int32(n), 1, int32(n))
	invP = dmatrix(1, int32(n), 1, int32(n))
	for i = 1; i <= n; i++ {
		for j = 1; j <= n; j++ {
			P[i][j] = V[p[i]][m[j]]
		}
	}
	pseudo_inv(P, invP, n, n, 1e-09, verbose)
	for i = 1; i <= N; i++ {
		if noarch.NotInt(R[i]) != 0 {
			traceM += M[i][i]
		}
	}
	for i = 1; i <= n; i++ {
		for j = 1; j <= n; j++ {
			Aij = 0
			for k = 1; k <= n; k++ {
				Aij += invP[k][i] * invP[k][j]
			}
			Mc[i][j] = Aij
		}
	}
	for i = 1; i <= n; i++ {
		traceMc += Mc[i][i]
	}
	for i = 1; i <= n; i++ {
		for j = 1; j <= n; j++ {
			Aij = 0
			for k = 1; k <= n; k++ {
				Aij += invP[k][i] * 4 * 3.141592653589793 * 3.141592653589793 * f[m[k]] * f[m[k]] * invP[k][j]
			}
			Kc[i][j] = Aij
		}
	}
	for i = 1; i <= n; i++ {
		for j = 1; j <= n; j++ {
			Mc[i][j] *= traceM / traceMc
		}
	}
	for i = 1; i <= n; i++ {
		for j = 1; j <= n; j++ {
			Kc[i][j] *= traceM / traceMc
		}
	}
	free_dmatrix(P, 1, int32(n), 1, int32(n))
	free_dmatrix(invP, 1, int32(n), 1, int32(n))
}

// deallocate - transpiled function from  $GOPATH/src/github.com/Konstantin8105/History_frame3DD/src/frame3dd.c:1076
/*
 * DEALLOCATE  -  release allocated memory					9sep08
 */ //
// printf("..B..element connectivity\n"); /* debug */
// printf("..C..section properties \n"); /* debug */
// printf("..D.. U W P T Dp\n"); /* debug */
// printf("..E..F_mech & F_temp\n"); /* debug */
// printf("..F.. eqF_mech & eqF_temp\n"); /* debug */
// printf("..G.. F & dF\n"); /* debug */
// printf("..H.. K & Q\n"); /* debug */
// printf("..I.. D  dD R dR \n"); /* debug */
// printf("..J.. extra mass\n"); /* debug */
// printf("..K.. peak stats\n"); /* debug */
// printf("..L.. M f V\n"); /* debug */
//
func deallocate(nN int, nE int, nL int, nF []int, nU []int, nW []int, nP []int, nT []int, DoF int, nM int, xyz []vec3, rj []float32, L []float64, Le []float64, N1 []int, N2 []int, q []int, r []int, Ax []float32, Asy []float32, Asz []float32, J []float32, Iy []float32, Iz []float32, E []float32, G []float32, p []float32, U [][][]float32, W [][][]float32, P [][][]float32, T [][][]float32, Dp [][]float32, F_mech [][]float64, F_temp [][]float64, eqF_mech [][][]float64, eqF_temp [][][]float64, F []float64, dF []float64, K [][]float64, Q [][]float64, D []float64, dD []float64, R []float64, dR []float64, d []float32, EMs []float32, NMs []float32, NMx []float32, NMy []float32, NMz []float32, M [][]float64, f []float64, V [][]float64, c []int, m []int, pkNx [][]float64, pkVy [][]float64, pkVz [][]float64, pkTx [][]float64, pkMy [][]float64, pkMz [][]float64, pkDx [][]float64, pkDy [][]float64, pkDz [][]float64, pkRx [][]float64, pkSy [][]float64, pkSz [][]float64) {
	_ = xyz
	free_vector(rj, 1, int32(nN))
	free_dvector(L, 1, int32(nE))
	free_dvector(Le, 1, int32(nE))
	free_ivector(N1, 1, int32(nE))
	free_ivector(N2, 1, int32(nE))
	free_ivector(q, 1, int32(DoF))
	free_ivector(r, 1, int32(DoF))
	free_vector(Ax, 1, int32(nE))
	free_vector(Asy, 1, int32(nE))
	free_vector(Asz, 1, int32(nE))
	free_vector(J, 1, int32(nE))
	free_vector(Iy, 1, int32(nE))
	free_vector(Iz, 1, int32(nE))
	free_vector(E, 1, int32(nE))
	free_vector(G, 1, int32(nE))
	free_vector(p, 1, int32(nE))
	free_D3matrix(U, 1, nL, 1, nE, 1, 4)
	free_D3matrix(W, 1, nL, 1, 10*nE, 1, 13)
	free_D3matrix(P, 1, nL, 1, 10*nE, 1, 5)
	free_D3matrix(T, 1, nL, 1, nE, 1, 8)
	free_matrix(Dp, 1, int32(nL), 1, int32(DoF))
	free_dmatrix(F_mech, 1, int32(nL), 1, int32(DoF))
	free_dmatrix(F_temp, 1, int32(nL), 1, int32(DoF))
	free_D3dmatrix(eqF_mech, 1, nL, 1, nE, 1, 12)
	free_D3dmatrix(eqF_temp, 1, nL, 1, nE, 1, 12)
	free_dvector(F, 1, int32(DoF))
	free_dvector(dF, 1, int32(DoF))
	free_dmatrix(K, 1, int32(DoF), 1, int32(DoF))
	free_dmatrix(Q, 1, int32(nE), 1, 12)
	free_dvector(D, 1, int32(DoF))
	free_dvector(dD, 1, int32(DoF))
	free_dvector(R, 1, int32(DoF))
	free_dvector(dR, 1, int32(DoF))
	free_vector(d, 1, int32(nE))
	free_vector(EMs, 1, int32(nE))
	free_vector(NMs, 1, int32(nN))
	free_vector(NMx, 1, int32(nN))
	free_vector(NMy, 1, int32(nN))
	free_vector(NMz, 1, int32(nN))
	free_ivector(c, 1, int32(DoF))
	free_ivector(m, 1, int32(DoF))
	free_dmatrix(pkNx, 1, int32(nL), 1, int32(nE))
	free_dmatrix(pkVy, 1, int32(nL), 1, int32(nE))
	free_dmatrix(pkVz, 1, int32(nL), 1, int32(nE))
	free_dmatrix(pkTx, 1, int32(nL), 1, int32(nE))
	free_dmatrix(pkMy, 1, int32(nL), 1, int32(nE))
	free_dmatrix(pkMz, 1, int32(nL), 1, int32(nE))
	free_dmatrix(pkDx, 1, int32(nL), 1, int32(nE))
	free_dmatrix(pkDy, 1, int32(nL), 1, int32(nE))
	free_dmatrix(pkDz, 1, int32(nL), 1, int32(nE))
	free_dmatrix(pkRx, 1, int32(nL), 1, int32(nE))
	free_dmatrix(pkSy, 1, int32(nL), 1, int32(nE))
	free_dmatrix(pkSz, 1, int32(nL), 1, int32(nE))
	if nM > 0 {
		free_dmatrix(M, 1, int32(DoF), 1, int32(DoF))
		free_dvector(f, 1, int32(nM))
		free_dmatrix(V, 1, int32(DoF), 1, int32(DoF))
	}
}

const (
	_ISupper  int = (1 << 0) << 8
	_ISlower      = (1 << 1) << 8
	_ISalpha      = (1 << 2) << 8
	_ISdigit      = (1 << 3) << 8
	_ISxdigit     = (1 << 4) << 8
	_ISspace      = (1 << 5) << 8
	_ISprint      = (1 << 6) << 8
	_ISgraph      = (1 << 7) << 8
	_ISblank      = (1 << 8) >> 8
	_IScntrl      = (1 << 9) >> 8
	_ISpunct      = (1 << 10) >> 8
	_ISalnum      = (1 << 11) >> 8
)

// parse_options - transpiled function from  $GOPATH/src/github.com/Konstantin8105/History_frame3DD/src/frame3dd_io.c:58
/*
 This file is part of FRAME3DD:
 Static and dynamic structural analysis of 2D and 3D frames and trusses with
 elastic and geometric stiffness.
 ---------------------------------------------------------------------------
 http://frame3dd.sourceforge.net/
 ---------------------------------------------------------------------------
 Copyright (C) 1992-2015  Henri P. Gavin

 FRAME3DD is free software: you can redistribute it and/or modify
 it under the terms of the GNU General Public License as published by
 the Free Software Foundation, either version 3 of the License, or
 (at your option) any later version.

 FRAME3DD is distributed in the hope that it will be useful,
 but WITHOUT ANY WARRANTY; without even the implied warranty of
 MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 GNU General Public License for more details.

 You should have received a copy of the GNU General Public License
 along with FRAME3DD.  If not, see <http://www.gnu.org/licenses/>.
*/ //
/**
@file
Input/output routines for FRAME.

@note The file format for FRAME is defined in doc/user_manual.html.
*/ //
/* #define MASSDATA_DEBUG */ //
/* forward decls */ //
/**< pointer to the file from which to read */ //
/**< pointer to the string to which to write */ //
/**< the longest anticipated line length  */ //
/*
 * PARSE_OPTIONS -  parse command line options
 * command line options over-ride values in the input data file
 * 04 Mar 2009, 22 Sep 2009
 */ //
/* char	option; */ //
/* char	errMsg[MAXL]; */ //
/* int	sfrv=0;		#<{(| *scanf return value	|)}># */ //
/* default values */ //
/*  */ //
/* strcpy(  IN_file , "\0" ); */ //
/* strcpy( OUT_file , "\0" ); */ //
/*  */ //
/* set up file names for the the input data and the output data */ //
/*  */ //
/* switch ( argc ) { */ //
/*  case 1: { */ //
/* 	fprintf(stderr,"\n Frame3DD version: %s\n", VERSION); */ //
/* 	fprintf(stderr," Analysis of 2D and 3D structural frames with elastic and geometric stiffness.\n"); */ //
/* 	fprintf(stderr," http://frame3dd.sourceforge.net\n\n"); */ //
/* 	fprintf (stderr," Please enter the  input data file name: "); */ //
/* 	sfrv=scanf("%s", IN_file ); */ //
/* 	if (sfrv != 1) sferr("IN_file"); */ //
/* 	fprintf (stderr," Please enter the output data file name: "); */ //
/* 	sfrv=scanf("%s", OUT_file ); */ //
/* 	if (sfrv != 1) sferr("OUT_file"); */ //
/* 	return; */ //
/*  } */ //
/*  case 3: { */ //
/* 	if ( argv[1][0] != '-' ) { */ //
/* 		strcpy(  IN_file , argv[1] ); */ //
/* 		strcpy( OUT_file , argv[2] ); */ //
/* 		return; */ //
/* 	} */ //
/*  } */ //
/* } */ //
/*  */ //
/* // remaining unused flags ... b j k n u y  */ //
/*  */ //
/* while ((option=getopt(argc,argv, "i:o:acdhqvwxzs:e:f:g:l:m:p:r:t:")) != -1){ */ //
/* 	switch ( option ) { */ //
/* 		case 'i':		#<{(| input data file name |)}># */ //
/* 			strcpy(IN_file,optarg); */ //
/* 			break; */ //
/* 		case 'o':		#<{(| output data file name |)}># */ //
/* 			strcpy(OUT_file,optarg); */ //
/* 			break; */ //
/* 		case 'h':		#<{(| help	|)}># */ //
/* 			display_help(); */ //
/* 			exit(0); */ //
/* 		case 'v':		#<{(| version |)}># */ //
/* 			display_version(); */ //
/* 			exit(0); */ //
/* 		case 'a':		#<{(| about |)}># */ //
/* 			display_version_about(); */ //
/* 			exit(0); */ //
/* 		case 'q':		#<{(| quiet |)}># */ //
/* 			*verbose = 0; */ //
/* 			break; */ //
/* 		case 'c':		#<{(| data check only |)}># */ //
/* 			*anlyz_flag = 0; */ //
/* 			break; */ //
/* 		case 'd':		#<{(| debug |)}># */ //
/* 			*debug = 1; */ //
/* 			break; */ //
/* 		case 'w':		#<{(| write stiffness and mass |)}># */ //
/* 			*write_matrix = 1; */ //
/* 			break; */ //
/* 		case 'x':		#<{(| write sign of axial forces |)}># */ //
/* 			*axial_sign = 0; */ //
/* 			break; */ //
/* 		case 's':		#<{(| shear deformation |)}># */ //
/* 			if (strcmp(optarg,"Off")==0) */ //
/* 				*shear_flag = 0; */ //
/* 			else if (strcmp(optarg,"On")==0) */ //
/* 				*shear_flag = 1; */ //
/* 			else { */ //
/* 			 errorMsg("\n frame3dd command-line error: argument to -s option should be either On or Off\n"); */ //
/* 			 exit(3); */ //
/* 			} */ //
/* 			break; */ //
/* 		case 'g':		#<{(| geometric stiffness |)}># */ //
/* 			if (strcmp(optarg,"Off")==0) */ //
/* 				*geom_flag = 0; */ //
/* 			else if (strcmp(optarg,"On")==0) */ //
/* 				*geom_flag = 1; */ //
/* 			else { */ //
/* 			 errorMsg("\n frame3dd command-line error: argument to -g option should be either On or Off\n"); */ //
/* 			 exit(4); */ //
/* 			} */ //
/* 			break; */ //
/* 		case 'e':		#<{(| static mesh exagg. factor |)}># */ //
/* 			*exagg_flag = atof(optarg); */ //
/* 			break; */ //
/* 		case 'z':		#<{(| force 3D plotting |)}># */ //
/* 			*D3_flag = 1; */ //
/* 			break; */ //
/* 		case 'l':		#<{(| lumped or consistent mass |)}># */ //
/* 			if (strcmp(optarg,"Off")==0) */ //
/* 				*lump_flag = 0; */ //
/* 			else if (strcmp(optarg,"On")==0) */ //
/* 				*lump_flag = 1; */ //
/* 			else { */ //
/* 			 errorMsg("\n frame3dd command-line error: argument to -l option should be either On or Off\n"); */ //
/* 			 exit(5); */ //
/* 			} */ //
/* 			break; */ //
/* 		case 'm':		#<{(| modal analysis method |)}># */ //
/* 			if (strcmp(optarg,"J")==0) */ //
/* 				*modal_flag = 1; */ //
/* 			else if (strcmp(optarg,"S")==0) */ //
/* 				*modal_flag = 2; */ //
/* 			else { */ //
/* 			 errorMsg("\n frame3dd command-line error: argument to -m option should be either J or S\n"); */ //
/* 			 exit(6); */ //
/* 			} */ //
/* 			break; */ //
/* 		case 't':		#<{(| modal analysis tolerence |)}># */ //
/* 			*tol_flag = atof(optarg); */ //
/* 			if (*tol_flag == 0.0) { */ //
/* 			 errorMsg("\n frame3dd command-line error: argument to -t option should be a number.\n"); */ //
/* 			 exit(7); */ //
/* 			} */ //
/* 			break; */ //
/* 		case 'f':		#<{(| modal analysis freq. shift |)}># */ //
/* 			*shift_flag = atof(optarg); */ //
/* 			if (*shift_flag == 0.0) { */ //
/* 			 errorMsg("\n frame3dd command-line error: argument to -f option should be a number.\n"); */ //
/* 			 exit(8); */ //
/* 			} */ //
/* 			break; */ //
/* 		case 'p':		#<{(| pan rate	|)}># */ //
/* 			*pan_flag = atof(optarg); */ //
/* 			if (*pan_flag < 0.0) { */ //
/* 			 errorMsg("\n frame3dd command-line error: argument to -p option should be a positive number.\n"); */ //
/* 			 exit(9); */ //
/* 			} */ //
/* 			break; */ //
/* 		case 'r':		#<{(| matrix condensation method |)}># */ //
/* 			*condense_flag = atoi(optarg); */ //
/* 			if (*condense_flag < 0 || *condense_flag > 3) { */ //
/* 			 errorMsg("\n frame3dd command-line error: argument to -r option should be 0, 1, or 2.\n"); */ //
/* 			 exit(10); */ //
/* 			} */ //
/* 			break; */ //
/* 		case '?': */ //
/* 			sprintf(errMsg,"  Missing argument or Unknown option: -%c\n\n", option ); */ //
/* 			errorMsg(errMsg); */ //
/* 			display_help(); */ //
/* 			exit(2); */ //
/* 	} */ //
/* } */ //
/* if ( strcmp(IN_file,"\0") == 0 ) { */ //
/* 	fprintf (stderr," Please enter the  input data file name: "); */ //
/* 	sfrv=scanf("%s", IN_file ); */ //
/* 	if (sfrv != 1) sferr("IN_file"); */ //
/* 	fprintf (stderr," Please enter the output data file name: "); */ //
/* 	sfrv=scanf("%s", OUT_file ); */ //
/* 	if (sfrv != 1) sferr("OUT_file"); */ //
/* } */ //
//
func parse_options(argc int, argv [][]byte, IN_file []byte, OUT_file []byte, shear_flag []int, geom_flag []int, anlyz_flag []int, exagg_flag []float64, D3_flag []int, lump_flag []int, modal_flag []int, tol_flag []float64, shift_flag []float64, pan_flag []float32, write_matrix []int, axial_sign []int, condense_flag []int, verbose []int, debug []int) {
	modal_flag[0] = -1
	lump_flag[0] = modal_flag[0]
	anlyz_flag[0] = lump_flag[0]
	geom_flag[0] = anlyz_flag[0]
	shear_flag[0] = geom_flag[0]
	shift_flag[0] = -1
	tol_flag[0] = shift_flag[0]
	exagg_flag[0] = tol_flag[0]
	D3_flag[0] = 0
	pan_flag[0] = float32(-1)
	condense_flag[0] = -1
	write_matrix[0] = 0
	axial_sign[0] = 1
	debug[0] = 0
	verbose[0] = 1
	noarch.Strcpy(IN_file, []byte("exA.3dd\x00\x00"))
	noarch.Strcpy(OUT_file, []byte("exA.3dd.out\x00\x00"))
	if noarch.Strcmp(IN_file, []byte("\x00\x00")) != 0 && noarch.Strcmp(OUT_file, []byte("\x00\x00")) == 0 {
		noarch.Strcpy(OUT_file, IN_file)
		noarch.Strcat(OUT_file, []byte(".out\x00"))
	}
}

// display_help - transpiled function from  $GOPATH/src/github.com/Konstantin8105/History_frame3DD/src/frame3dd_io.c:257
/*
 * DISPLAY_HELP -  display help information to stderr
 * 04 Mar 2009, 22 Sep 2009
 */ //
//frame3dd.sourceforge.net\n\n");
/* fprintf(stderr,"  Usage: frame3dd -i<input> -o<output> [-hvcqz] [-s<On|Off>] [-g<On|Off>] [-e<value>] [-l<On|Off>] [-f<value>] [-m J|S] [-t<value>] [-p<value>] \n");
 */ //
//
func display_help() {
	textColor('g', 'x', 'x', 'x')
	noarch.Fprintf(stderr, []byte("\n Frame3DD version: %s\n\x00"), []byte("20140514+\x00"))
	noarch.Fprintf(stderr, []byte(" Analysis of 2D and 3D structural frames with elastic and geometric stiffness.\n\x00"))
	noarch.Fprintf(stderr, []byte(" http://frame3dd.sourceforge.net\n\n\x00"))
	noarch.Fprintf(stderr, []byte("  Frame3DD may be run with interactive prompting for file names by typing ...\n\x00"))
	noarch.Fprintf(stderr, []byte("       frame3dd \n\n\x00"))
	noarch.Fprintf(stderr, []byte("  Frame3DD may be run without command-line options by typing ...\n\x00"))
	noarch.Fprintf(stderr, []byte("       frame3dd <InFile> <OutFile> \n\n\x00"))
	noarch.Fprintf(stderr, []byte("  Frame3DD may be run with command-line options by typing ...\n\x00"))
	noarch.Fprintf(stderr, []byte("       frame3dd -i <InFile> -o <OutFile> [OPTIONS] \n\n\x00"))
	noarch.Fprintf(stderr, []byte(" ... where [OPTIONS] over-rides values in the input data file and includes\n\x00"))
	noarch.Fprintf(stderr, []byte("     one or more of the following:\n\n\x00"))
	noarch.Fprintf(stderr, []byte(" -------------------------------------------------------------------------\n\x00"))
	noarch.Fprintf(stderr, []byte("  -i  <InFile>  the  input data file name --- described in the manual\n\x00"))
	noarch.Fprintf(stderr, []byte("  -o <OutFile>  the output data file name\n\x00"))
	noarch.Fprintf(stderr, []byte("  -h            print this help message and exit\n\x00"))
	noarch.Fprintf(stderr, []byte("  -v            display program version, website, brief help info and exit\n\x00"))
	noarch.Fprintf(stderr, []byte("  -a            display program version, website and exit\n\x00"))
	noarch.Fprintf(stderr, []byte("  -c            data check only - the output data reviews the input data\n\x00"))
	noarch.Fprintf(stderr, []byte("  -w            write stiffness and mass matrices to files named Ks Kd Md\n\x00"))
	noarch.Fprintf(stderr, []byte("  -x            suppress writing of 't' or 'c' for sign of axial forces\n\x00"))
	noarch.Fprintf(stderr, []byte("  -q            suppress screen output except for warning messages\n\x00"))
	noarch.Fprintf(stderr, []byte("  -s  On|Off    On: include shear deformation or Off: neglect ...\n\x00"))
	noarch.Fprintf(stderr, []byte("  -g  On|Off    On: include geometric stiffness or Off: neglect ...\n\x00"))
	noarch.Fprintf(stderr, []byte("  -e <value>    static deformation exaggeration factor for Gnuplot output\n\x00"))
	noarch.Fprintf(stderr, []byte("  -z            force X-Y-Z plotting\n\x00"))
	noarch.Fprintf(stderr, []byte("  -l  On|Off    On: lumped mass matrix or Off: consistent mass matrix\n\x00"))
	noarch.Fprintf(stderr, []byte("  -f <value>    modal frequency shift for unrestrained structures\n\x00"))
	noarch.Fprintf(stderr, []byte("  -m   J|S      modal analysis method: J=Jacobi-Subspace or S=Stodola\n\x00"))
	noarch.Fprintf(stderr, []byte("  -t <value>    convergence tolerance for modal analysis\n\x00"))
	noarch.Fprintf(stderr, []byte("  -p <value>    pan rate for mode shape animation\n\x00"))
	noarch.Fprintf(stderr, []byte("  -r <value>    matrix condensation method: 0, 1, 2, or 3 \n\x00"))
	noarch.Fprintf(stderr, []byte(" -------------------------------------------------------------------------\n\x00"))
	color(0)
}

// display_usage - transpiled function from  $GOPATH/src/github.com/Konstantin8105/History_frame3DD/src/frame3dd_io.c:306
/*
 * DISPLAY_USAGE -  display usage information to stderr
 * 04 Mar 2009
 */ //
//frame3dd.sourceforge.net\n\n");
/* fprintf(stderr,"  Usage: frame3dd -i<input> -o<output> [-hvcqz] [-s<On|Off>] [-g<On|Off>] [-e<value>] [-l<On|Off>] [-f<value>] [-m J|S] [-t<value>] [-p<value>] \n");
 */ //
//
func display_usage() {
	noarch.Fprintf(stderr, []byte("\n Frame3DD version: %s\n\x00"), []byte("20140514+\x00"))
	noarch.Fprintf(stderr, []byte(" Analysis of 2D and 3D structural frames with elastic and geometric stiffness.\n\x00"))
	noarch.Fprintf(stderr, []byte(" http://frame3dd.sourceforge.net\n\n\x00"))
	noarch.Fprintf(stderr, []byte("  Usage: frame3dd -i <input> -o <output> [OPTIONS] \n\n\x00"))
	noarch.Fprintf(stderr, []byte("  Type ...   frame3dd -h   ... for additional help information.\n\n\x00"))
}

// display_version - transpiled function from  $GOPATH/src/github.com/Konstantin8105/History_frame3DD/src/frame3dd_io.c:323
/*
 * DISPLAY_VERSION_HELP -  display version, website, and brief help info. to stderr
 * 04 Mar 2009
 */ //
//frame3dd.sourceforge.net\n\n");
//
func display_version() {
	noarch.Fprintf(stderr, []byte("\n Frame3DD version: %s\n\x00"), []byte("20140514+\x00"))
	noarch.Fprintf(stderr, []byte(" Analysis of 2D and 3D structural frames with elastic and geometric stiffness.\n\x00"))
	noarch.Fprintf(stderr, []byte(" http://frame3dd.sourceforge.net\n\n\x00"))
	noarch.Fprintf(stderr, []byte("  Usage: frame3dd -i <input> -o <output> [OPTIONS] \n\n\x00"))
	noarch.Fprintf(stderr, []byte("  Type ...   frame3dd -h   ... for additional help information.\n\n\x00"))
}

// display_version_about - transpiled function from  $GOPATH/src/github.com/Konstantin8105/History_frame3DD/src/frame3dd_io.c:341
/*
 * DISPLAY_VERSION_ABOUT-  display version and website to stderr for
 * running as a background process
 * 22 Sep 2009
 * Contributed by Barry Sanford, barry.sanford@trimjoist.com
 */ //
//frame3dd.sourceforge.net\n");
//www.fsf.org/copyleft/gpl.html\n");
//
func display_version_about() {
	noarch.Fprintf(stderr, []byte(" Frame3DD version: %s\n\x00"), []byte("20140514+\x00"))
	noarch.Fprintf(stderr, []byte(" Analysis of 2D and 3D structural frames with elastic and geometric stiffness\n\x00"))
	noarch.Fprintf(stderr, []byte(" http://frame3dd.sourceforge.net\n\x00"))
	noarch.Fprintf(stderr, []byte(" GPL Copyright (C) 1992-2015, Henri P. Gavin \n\x00"))
	noarch.Fprintf(stderr, []byte(" Frame3DD is distributed in the hope that it will be useful\x00"))
	noarch.Fprintf(stderr, []byte(" but with no warranty.\n\x00"))
	noarch.Fprintf(stderr, []byte(" For details see the GNU Public Licence:\x00"))
	noarch.Fprintf(stderr, []byte(" http://www.fsf.org/copyleft/gpl.html\n\x00"))
}

// read_node_data - transpiled function from  $GOPATH/src/github.com/Konstantin8105/History_frame3DD/src/frame3dd_io.c:358
/*
 * READ_NODE_DATA  -  read node location data
 * 04 Jan 2009
 */ //
/* *scanf return value	*/ //
/* read node coordinates	*/ //
/* fprintf(stderr,"\nj = %d, pos = (%lf, %lf, %lf), r = %f", j, xyz[j].x, xyz[j].y, xyz[j].z, r[j]); */ //
//
func read_node_data(fp *noarch.File, nN int, xyz []vec3, r []float32) {
	var i int
	var j int
	var sfrv int
	var errMsg []byte = make([]byte, 512)
	for i = 1; i <= nN; i++ {
		sfrv = noarch.Fscanf(fp, []byte("%d\x00"), (*[100000000]int)(unsafe.Pointer(&j))[:])
		if sfrv != 1 {
			sferr([]byte("node number in node data\x00"))
		}
		if j <= 0 || j > nN {
			noarch.Sprintf(errMsg, []byte("\nERROR: in node coordinate data, node number out of range\n(node number %d is <= 0 or > %d)\n\x00"), j, nN)
			errorMsg(errMsg)
			os.Exit(41)
		}
		sfrv = noarch.Fscanf(fp, []byte("%lf %lf %lf %f\x00"), (*[100000000]float64)(unsafe.Pointer(&xyz[j].x))[:], (*[100000000]float64)(unsafe.Pointer(&xyz[j].y))[:], (*[100000000]float64)(unsafe.Pointer(&xyz[j].z))[:], (*[100000000]float32)(unsafe.Pointer(&r[j]))[:])
		if sfrv != 4 {
			sferr([]byte("node coordinates in node data\x00"))
		}
		r[j] = float32(math.Abs(float64(r[j])))
	}
}

// read_frame_element_data - transpiled function from  $GOPATH/src/github.com/Konstantin8105/History_frame3DD/src/frame3dd_io.c:385
/*
 * READ_FRAME_ELEMENT_DATA  -  read frame element property data
 * 04 Jan 2009
 */ //
/* vector of elements per node */ //
/* *scanf return value */ //
/* read frame element properties */ //
/* convert from degrees to radians */ //
/* calculate frame element lengths */ //
//
func read_frame_element_data(fp *noarch.File, nN int, nE int, xyz []vec3, r []float32, L []float64, Le []float64, N1 []int, N2 []int, Ax []float32, Asy []float32, Asz []float32, Jx []float32, Iy []float32, Iz []float32, E []float32, G []float32, p []float32, d []float32) {
	var n1 int
	var n2 int
	var i int
	var n int
	var b int
	var epn []int
	var epn0 int
	var sfrv int
	var errMsg []byte = make([]byte, 512)
	epn = ivector(1, int32(nN))
	for n = 1; n <= nN; n++ {
		epn[n] = 0
	}
	for i = 1; i <= nE; i++ {
		sfrv = noarch.Fscanf(fp, []byte("%d\x00"), (*[100000000]int)(unsafe.Pointer(&b))[:])
		if sfrv != 1 {
			sferr([]byte("frame element number in element data\x00"))
		}
		if b <= 0 || b > nE {
			noarch.Sprintf(errMsg, []byte("\n  error in frame element property data: Element number out of range  \n Frame element number: %d  \n\x00"), b)
			errorMsg(errMsg)
			os.Exit(51)
		}
		sfrv = noarch.Fscanf(fp, []byte("%d %d\x00"), (*[100000000]int)(unsafe.Pointer(&N1[b]))[:], (*[100000000]int)(unsafe.Pointer(&N2[b]))[:])
		epn[N1[b]]++
		epn[N2[b]]++
		if sfrv != 2 {
			sferr([]byte("node numbers in frame element data\x00"))
		}
		if N1[b] <= 0 || N1[b] > nN || N2[b] <= 0 || N2[b] > nN {
			noarch.Sprintf(errMsg, []byte("\n  error in frame element property data: node number out of range  \n Frame element number: %d \n\x00"), b)
			errorMsg(errMsg)
			os.Exit(52)
		}
		sfrv = noarch.Fscanf(fp, []byte("%f %f %f\x00"), (*[100000000]float32)(unsafe.Pointer(&Ax[b]))[:], (*[100000000]float32)(unsafe.Pointer(&Asy[b]))[:], (*[100000000]float32)(unsafe.Pointer(&Asz[b]))[:])
		if sfrv != 3 {
			sferr([]byte("section areas in frame element data\x00"))
		}
		sfrv = noarch.Fscanf(fp, []byte("%f %f %f\x00"), (*[100000000]float32)(unsafe.Pointer(&Jx[b]))[:], (*[100000000]float32)(unsafe.Pointer(&Iy[b]))[:], (*[100000000]float32)(unsafe.Pointer(&Iz[b]))[:])
		if sfrv != 3 {
			sferr([]byte("section inertias in frame element data\x00"))
		}
		sfrv = noarch.Fscanf(fp, []byte("%f %f\x00"), (*[100000000]float32)(unsafe.Pointer(&E[b]))[:], (*[100000000]float32)(unsafe.Pointer(&G[b]))[:])
		if sfrv != 2 {
			sferr([]byte("material moduli in frame element data\x00"))
		}
		sfrv = noarch.Fscanf(fp, []byte("%f\x00"), (*[100000000]float32)(unsafe.Pointer(&p[b]))[:])
		if sfrv != 1 {
			sferr([]byte("roll angle in frame element data\x00"))
		}
		p[b] = float32(float64(p[b]) * 3.141592653589793 / 180)
		sfrv = noarch.Fscanf(fp, []byte("%f\x00"), (*[100000000]float32)(unsafe.Pointer(&d[b]))[:])
		if sfrv != 1 {
			sferr([]byte("mass density in frame element data\x00"))
		}
		if Ax[b] < 0 || Asy[b] < 0 || Asz[b] < 0 || Jx[b] < 0 || Iy[b] < 0 || Iz[b] < 0 {
			noarch.Sprintf(errMsg, []byte("\n  error in frame element property data: section property < 0 \n  Frame element number: %d  \n\x00"), b)
			errorMsg(errMsg)
			os.Exit(53)
		}
		if Ax[b] == 0 {
			noarch.Sprintf(errMsg, []byte("\n  error in frame element property data: cross section area is zero   \n  Frame element number: %d  \n\x00"), b)
			errorMsg(errMsg)
			os.Exit(54)
		}
		if (Asy[b] == 0 || Asz[b] == 0) && G[b] == 0 {
			noarch.Sprintf(errMsg, []byte("\n  error in frame element property data: a shear area and shear modulus are zero   \n  Frame element number: %d  \n\x00"), b)
			errorMsg(errMsg)
			os.Exit(55)
		}
		if Jx[b] == 0 {
			noarch.Sprintf(errMsg, []byte("\n  error in frame element property data: torsional moment of inertia is zero   \n  Frame element number: %d  \n\x00"), b)
			errorMsg(errMsg)
			os.Exit(56)
		}
		if Iy[b] == 0 || Iz[b] == 0 {
			noarch.Sprintf(errMsg, []byte("\n  error: cross section bending moment of inertia is zero   \n  Frame element number : %d  \n\x00"), b)
			errorMsg(errMsg)
			os.Exit(57)
		}
		if E[b] <= 0 || G[b] <= 0 {
			noarch.Sprintf(errMsg, []byte("\n  error : material elastic modulus E or G is not positive   \n  Frame element number: %d  \n\x00"), b)
			errorMsg(errMsg)
			os.Exit(58)
		}
		if d[b] <= 0 {
			noarch.Sprintf(errMsg, []byte("\n  error : mass density d is not positive   \n  Frame element number: %d  \n\x00"), b)
			errorMsg(errMsg)
			os.Exit(59)
		}
	}
	for b = 1; b <= nE; b++ {
		n1 = N1[b]
		n2 = N2[b]
		L[b] = (xyz[n2].x-xyz[n1].x)*(xyz[n2].x-xyz[n1].x) + (xyz[n2].y-xyz[n1].y)*(xyz[n2].y-xyz[n1].y) + (xyz[n2].z-xyz[n1].z)*(xyz[n2].z-xyz[n1].z)
		L[b] = math.Sqrt(L[b])
		Le[b] = L[b] - float64(r[n1]) - float64(r[n2])
		if n1 == n2 || L[b] == 0 {
			noarch.Sprintf(errMsg, []byte(" Frame elements must start and stop at different nodes\n  frame element %d  N1= %d N2= %d L= %e\n   Perhaps frame element number %d has not been specified.\n  or perhaps the Input Data file is missing expected data.\n\x00"), b, n1, n2, L[b], i)
			errorMsg(errMsg)
			os.Exit(60)
		}
		if Le[b] <= 0 {
			noarch.Sprintf(errMsg, []byte(" Node  radii are too large.\n  frame element %d  N1= %d N2= %d L= %e \n  r1= %e r2= %e Le= %e \n\x00"), b, n1, n2, L[b], float64(r[n1]), float64(r[n2]), Le[b])
			errorMsg(errMsg)
			os.Exit(61)
		}
	}
	for n = 1; n <= nN; n++ {
		if epn[n] == 0 {
			noarch.Sprintf(errMsg, []byte("node or frame element property data:\n     node number %3d is unconnected. \n\x00"), n)
			sferr(errMsg)
			epn0++
		}
	}
	free_ivector(epn, 1, int32(nN))
	if epn0 > 0 {
		os.Exit(42)
	}
}

// read_run_data - transpiled function from  $GOPATH/src/github.com/Konstantin8105/History_frame3DD/src/frame3dd_io.c:519
/*
 * READ_RUN_DATA  -  read information for analysis
 * 29 Dec 2008
 */ //
/* output data file name */ //
/* *scanf return value */ //
/* the length of the base_file */ //
/* find the last '.' in base_file */ //
/* end base_file at the last '.' */ //
/* find the last '/' or '\' in base_file */ //
/* over-ride values from input data file with command-line options */ //
//
func read_run_data(fp *noarch.File, OUT_file []byte, shear []int, shear_flag int, geom []int, geom_flag int, meshpath []byte, plotpath []byte, infcpath []byte, exagg_static []float64, exagg_flag float64, scale []float32, dx []float32, anlyz []int, anlyz_flag int, debug int) {
	var full_len int
	var len int
	var i int
	var base_file []byte = []byte("EMPTY_BASE\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00")
	var mesh_file []byte = []byte("EMPTY_MESH\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00")
	var sfrv int
	noarch.Strcpy(base_file, OUT_file)
	for int(base_file[func() int {
		defer func() {
			len++
		}()
		return len
	}()]) != int('\x00') {
	}
	full_len = len
	for int(base_file[func() int {
		defer func() {
			len--
		}()
		return len
	}()]) != int('.') && len > 0 {
	}
	if len == 0 {
		len = full_len
	}
	base_file[func() int {
		len++
		return len
	}()] = '\x00'
	noarch.Strcpy(plotpath, base_file)
	noarch.Strcat(plotpath, []byte(".plt\x00"))
	noarch.Strcpy(infcpath, base_file)
	noarch.Strcat(infcpath, []byte(".if\x00"))
	for int(base_file[len]) != int('/') && int(base_file[len]) != int('\\') && len > 0 {
		len--
	}
	i = 0
	for int(base_file[len]) != int('\x00') {
		mesh_file[func() int {
			defer func() {
				i++
			}()
			return i
		}()] = base_file[func() int {
			defer func() {
				len++
			}()
			return len
		}()]
	}
	mesh_file[i] = '\x00'
	noarch.Strcat(mesh_file, []byte("-msh\x00"))
	output_path(mesh_file, meshpath, 512, nil)
	if debug != 0 {
		noarch.Fprintf(stderr, []byte("OUT_FILE  = %s \n\x00"), OUT_file)
		noarch.Fprintf(stderr, []byte("BASE_FILE = %s \n\x00"), base_file)
		noarch.Fprintf(stderr, []byte("PLOTPATH  = %s \n\x00"), plotpath)
		noarch.Fprintf(stderr, []byte("MESH_FILE = %s \n\x00"), mesh_file)
		noarch.Fprintf(stderr, []byte("MESHPATH  = %s \n\x00"), meshpath)
	}
	sfrv = noarch.Fscanf(fp, []byte("%d %d %lf %f %f\x00"), shear, geom, exagg_static, scale, dx)
	if sfrv != 5 {
		sferr([]byte("shear, geom, exagg_static, scale, or dx variables\x00"))
	}
	if shear[0] != 0 && shear[0] != 1 {
		errorMsg([]byte(" Rember to specify shear deformations with a 0 or a 1 \n after the frame element property info.\n\x00"))
		os.Exit(71)
	}
	if geom[0] != 0 && geom[0] != 1 {
		errorMsg([]byte(" Rember to specify geometric stiffness with a 0 or a 1 \n after the frame element property info.\n\x00"))
		os.Exit(72)
	}
	if exagg_static[0] < 0 {
		errorMsg([]byte(" Remember to specify an exageration factor greater than zero.\n\x00"))
		os.Exit(73)
	}
	if float64(dx[0]) <= 0 && dx[0] != float32(-1) {
		errorMsg([]byte(" Remember to specify a frame element increment greater than zero.\n\x00"))
		os.Exit(74)
	}
	if shear_flag != -1 {
		shear[0] = shear_flag
	}
	if geom_flag != -1 {
		geom[0] = geom_flag
	}
	if exagg_flag != -1 {
		exagg_static[0] = exagg_flag
	}
	if float64(anlyz_flag) != -1 {
		anlyz[0] = anlyz_flag
	}
}

// frame3dd_getline - transpiled function from  $GOPATH/src/github.com/Konstantin8105/History_frame3DD/src/frame3dd_io.c:612
/*
 * FRAME3DD_GETLINE -  get line into a character string. from K&R        03feb94
 */ //
/*      if (c == '\n')  s[i++] = c;	*/ //
//
func frame3dd_getline(fp *noarch.File, s []byte, lim int) int {
	var c int
	var i int
	for func() int {
		lim--
		return lim
	}() > 0 && (func() int {
		c = noarch.Fgetc(fp)
		return c
	}()) != -1 && c != int('\n') {
		s[func() int {
			defer func() {
				i++
			}()
			return i
		}()] = byte(c)
	}
	s[i] = '\x00'
	return i
}

var sep byte = '/'

// temp_dir - transpiled function from  $GOPATH/src/github.com/Konstantin8105/History_frame3DD/src/frame3dd_io.c:640
/* platform-dependent path sperator character ... */ //
/*
 * TEMP_DIR
 * return platform-specific temp file location --
 * John Pye, Feb 2009
 */ //
/* Linux, Unix, OS X	*/ //
//
func temp_dir() []byte {
	var tmp []byte = []byte("/tmp\x00")
	return tmp
}

// output_path - transpiled function from  $GOPATH/src/github.com/Konstantin8105/History_frame3DD/src/frame3dd_io.c:665
/*
 * OUTPUT_PATH
 * return path for output files using either current directory, or FRAME3DD_OUTDIR
 * if specified. --
 * John Pye, Feb 2009.
 */ //
/*			deprecated code, January 15 2010 ...
	if ( fname[0]==sep ) {	in Win32 absolute path starts with C:\ not \ ??
		// absolute output path specified
//		res = snprintf(fullpath,len,"%s",fname);
		res = sprintf(fullpath,"%s",fname);
	} else {
*/ //
//		fprintf(stderr,"Generating output path for file '%s'\n",fname);
//		res = snprintf(fullpath,len,"%s%c%s",outdir,sep,fname);
/*			closing bracket for deprecated code "if"
}
*/ //
//	printf("Output file path generated: %s\n",fullpath); /* debug */
//
func output_path(fname []byte, fullpath []byte, len int, default_outdir []byte) {
	var res int
	func() {
		if fname != nil {
		} else {
			linux.AssertFail([]byte("fname!=NULL\x00"), []byte("/home/konstantin/go/src/github.com/Konstantin8105/History_frame3DD/src/frame3dd_io.c\x00"), 667, []byte("void print_number(int *)"))
		}
	}()
	var outdir []byte
	outdir = noarch.Getenv([]byte("FRAME3DD_OUTDIR\x00"))
	if outdir == nil {
		if default_outdir == nil {
			outdir = temp_dir()
		} else {
			outdir = default_outdir
		}
	}
	res = noarch.Sprintf(fullpath, []byte("%s%c%s\x00"), outdir, int(sep), fname)
	if res > len {
		errorMsg([]byte("ERROR: unable to construct output filename: overflow.\n\x00"))
		os.Exit(16)
	}
}

// parse_input - transpiled function from  $GOPATH/src/github.com/Konstantin8105/History_frame3DD/src/frame3dd_io.c:706
/*
 * PARSE_INPUT
 * strip comments from the input file, and write a stripped input file
 * 07 May 2003
 */ //
/* stripped input file pointer	*/ //
//
func parse_input(fp *noarch.File, tpath []byte) {
	var fpc *noarch.File
	var line []byte = make([]byte, 256)
	var errMsg []byte = make([]byte, 512)
	if (func() *noarch.File {
		fpc = noarch.Fopen(tpath, []byte("w\x00"))
		return fpc
	}()) == nil {
		noarch.Sprintf(errMsg, []byte("\n  error: cannot open parsed input data file: '%s' \n\x00"), tpath)
		errorMsg(errMsg)
		os.Exit(12)
	}
	for {
		getline_no_comment(fp, line, 256)
		noarch.Fprintf(fpc, []byte("%s \n\x00"), line)
		if !(int(line[0]) != int('_') && int(line[0]) != -1+256) {
			break
		}
	}
	noarch.Fclose(fpc)
}

// getline_no_comment - transpiled function from  $GOPATH/src/github.com/Konstantin8105/History_frame3DD/src/frame3dd_io.c:736
/*
 * GETLINE_NO_COMMENT
 * get a line into a character string. from K&R
 * get the line only up to one of the following characters:  \n  %  #  ?
 * ignore all comma (,) characters
 * ignore all double quote (") characters
 * ignore all semi-colon (;) characters
 * 09 Feb 2009
 */ //
/**< pointer to the file from which to read */ //
/**< pointer to the string to which to write */ //
/**< the longest anticipated line length  */ //
/*      if (c == '\n')  s[i++] = c;     */ //
/* read the rest of the line, otherwise do nothing */ //
//
func getline_no_comment(fp *noarch.File, s []byte, lim int) {
	var c int
	var i int
	for func() int {
		lim--
		return lim
	}() > 0 && (func() int {
		c = noarch.Fgetc(fp)
		return c
	}()) != -1 && c != int('\n') && c != int('%') && c != int('#') && c != int('?') {
		if c != int(',') && c != int('"') && c != int(';') {
			s[func() int {
				defer func() {
					i++
				}()
				return i
			}()] = byte(c)
		} else {
			s[func() int {
				defer func() {
					i++
				}()
				return i
			}()] = ' '
		}
	}
	s[i] = '\x00'
	if c != int('\n') {
		for func() int {
			lim--
			return lim
		}() > 0 && (func() int {
			c = noarch.Fgetc(fp)
			return c
		}()) != -1 && c != int('\n') {
		}
	}
	if c == -1 {
		s[0] = byte(-1 + 256)
	}
}

// read_reaction_data - transpiled function from  $GOPATH/src/github.com/Konstantin8105/History_frame3DD/src/frame3dd_io.c:766
/*
 * READ_REACTION_DATA - Read fixed node displacement boundary conditions
 * 29 Dec 2009
 */ //
/* *scanf return value */ //
/* read restrained degrees of freedom */ //
/*	exit(84); */ //
//
func read_reaction_data(fp *noarch.File, DoF int, nN int, nR []int, q []int, r []int, sumR []int, verbose int) {
	var i int
	var j int
	var l int
	var sfrv int
	var errMsg []byte = make([]byte, 512)
	for i = 1; i <= DoF; i++ {
		r[i] = 0
	}
	sfrv = noarch.Fscanf(fp, []byte("%d\x00"), nR)
	if sfrv != 1 {
		sferr([]byte("number of reactions in reaction data\x00"))
	}
	if verbose != 0 {
		noarch.Fprintf(stdout, []byte(" number of nodes with reactions \x00"))
		dots(stdout, 21)
		noarch.Fprintf(stdout, []byte(" nR =%4d \x00"), nR[0])
	}
	if nR[0] < 0 || nR[0] > DoF/6 {
		noarch.Fprintf(stderr, []byte(" number of nodes with reactions \x00"))
		dots(stderr, 21)
		noarch.Fprintf(stderr, []byte(" nR = %3d \x00"), nR[0])
		noarch.Sprintf(errMsg, []byte("\n  error: valid ranges for nR is 0 ... %d \n\x00"), DoF/6)
		errorMsg(errMsg)
		os.Exit(80)
	}
	for i = 1; i <= nR[0]; i++ {
		sfrv = noarch.Fscanf(fp, []byte("%d\x00"), (*[100000000]int)(unsafe.Pointer(&j))[:])
		if sfrv != 1 {
			sferr([]byte("node number in reaction data\x00"))
		}
		for l = 5; l >= 0; l-- {
			sfrv = noarch.Fscanf(fp, []byte("%d\x00"), (*[100000000]int)(unsafe.Pointer(&r[6*j-l]))[:])
			if sfrv != 1 {
				sferr([]byte("reaction value in reaction data\x00"))
			}
			if j > nN {
				noarch.Sprintf(errMsg, []byte("\n  error in reaction data: node number %d is greater than the number of nodes, %d \n\x00"), j, nN)
				errorMsg(errMsg)
				os.Exit(81)
			}
			if r[6*j-l] != 0 && r[6*j-l] != 1 {
				noarch.Sprintf(errMsg, []byte("\n  error in reaction data: Reaction data must be 0 or 1\n   Data for node %d, DoF %d is %d\n\x00"), j, 6-l, r[6*j-l])
				errorMsg(errMsg)
				os.Exit(82)
			}
		}
		sumR[0] = 0
		for l = 5; l >= 0; l-- {
			sumR[0] += r[6*j-l]
		}
		if sumR[0] == 0 {
			noarch.Sprintf(errMsg, []byte("\n  error: node %3d has no reactions\n   Remove node %3d from the list of reactions\n   and set nR to %3d \n\x00"), j, j, nR[0]-1)
			errorMsg(errMsg)
			os.Exit(83)
		}
	}
	sumR[0] = 0
	for i = 1; i <= DoF; i++ {
		sumR[0] += r[i]
	}
	if sumR[0] < 4 {
		noarch.Sprintf(errMsg, []byte("\n  Warning:  un-restrained structure   %d imposed reactions.\n  At least 4 reactions are required to support static loads.\n\x00"), sumR[0])
		errorMsg(errMsg)
	}
	if sumR[0] >= DoF {
		noarch.Sprintf(errMsg, []byte("\n  error in reaction data:  Fully restrained structure\n   %d imposed reactions >= %d degrees of freedom\n\x00"), sumR[0], DoF)
		errorMsg(errMsg)
		os.Exit(85)
	}
	for i = 1; i <= DoF; i++ {
		if r[i] != 0 {
			q[i] = 0
		} else {
			q[i] = 1
		}
	}
}

// read_and_assemble_loads - transpiled function from  $GOPATH/src/github.com/Konstantin8105/History_frame3DD/src/frame3dd_io.c:846
/*
 * READ_AND_ASSEMBLE_LOADS
 * Read load information data, assemble load vectors in global coordinates
 * Returns vector of equivalent loadal forces F_temp and F_mech and
 * a matrix of equivalent element end forces eqF_temp and eqF_mech from
 * distributed internal and temperature loadings.
 * eqF_temp and eqF_mech are computed for the global coordinate system
 * 2008-09-09, 2015-05-15
 */ //
// equivalent mech loads, global coord
// equivalent temp loads, global coord
/* section dimensions in local coords */ //
/* equivalent element end forces from distributed and thermal loads */ //
/* shear deformatn coefficients	*/ //
/* point load locations */ //
/* 3D coord Xfrm coeffs */ //
/* *scanf return value */ //
/* initialize load data vectors and matrices to zero */ //
/* begin load-case loop */ //
/*  display the load case number */ //
/* gravity loads applied uniformly to all frame elements ------- */ //
/* debugging ... check eqF data
printf("n=%d ", n);
for (l=1;l<=12;l++) {
	if (eqF_mech[lc][n][l] != 0)
	   printf(" eqF %d = %9.2e ", l, eqF_mech[lc][n][l] );
}
printf("\n");
*/ //
/* end gravity loads */ //
/* node point loads -------------------------------------------- */ //
/* ! global structural coordinates ! */ //
/* end node point loads  */ //
/* uniformly distributed loads --------------------------------- */ //
/* ! local element coordinates ! */ //
/* debugging ... check end force values
 * printf("n=%d Vy=%9.2e Vz=%9.2e My=%9.2e Mz=%9.2e\n",
 *				n, Vy1,Vz1, My1,Mz1 );
 */ //
/* debugging ... check coordinate transform coefficients
printf("t1=%5.2f t2=%5.2f t3=%5.2f \n", t1, t2, t3 );
printf("t4=%5.2f t5=%5.2f t6=%5.2f \n", t4, t5, t6 );
printf("t7=%5.2f t8=%5.2f t9=%5.2f \n", t7, t8, t9 );
*/ //
/* {F} = [T]'{Q} */ //
/* debugging ... check eqF values
printf("n=%d ", n);
for (l=1;l<=12;l++) {
	if (eqF_mech[lc][n][l] != 0)
	   printf(" eqF %d = %9.2e ", l, eqF_mech[lc][n][l] );
}
printf("\n");
*/ //
/* end uniformly distributed loads */ //
/* trapezoidally distributed loads ----------------------------- */ //
/* ! local element coordinates ! */ //
/* error checking */ //
/* x-axis trapezoidal loads (along the frame element length) */ //
/* y-axis trapezoidal loads (across the frame element length) */ //
/* z-axis trapezoidal loads (across the frame element length) */ //
/* debugging ... check internal force values
printf("n=%d\n Nx1=%9.3f\n Nx2=%9.3f\n Vy1=%9.3f\n Vy2=%9.3f\n Vz1=%9.3f\n Vz2=%9.3f\n My1=%9.3f\n My2=%9.3f\n Mz1=%9.3f\n Mz2=%9.3f\n",
		n, Nx1,Nx2,Vy1,Vy2,Vz1,Vz2, My1,My2,Mz1,Mz2 );
*/ //
/* debugging ... check coordinate transformation coefficients
printf("t1=%5.2f t2=%5.2f t3=%5.2f \n", t1, t2, t3 );
printf("t4=%5.2f t5=%5.2f t6=%5.2f \n", t4, t5, t6 );
printf("t7=%5.2f t8=%5.2f t9=%5.2f \n", t7, t8, t9 );
*/ //
/* {F} = [T]'{Q} */ //
/* debugging ... check eqF data
for (l=1;l<=13;l++) printf(" %9.2e ", W[lc][i][l] );
printf("\n");
printf("n=%d ", n);
for (l=1;l<=12;l++) {
	if (eqF_mech[lc][n][l] != 0)
	   printf(" eqF %d = %9.3f ", l, eqF_mech[lc][n][l] );
}
printf("\n");
*/ //
/* end trapezoidally distributed loads */ //
/* internal element point loads -------------------------------- */ //
/* ! local element coordinates ! */ //
/* {F} = [T]'{Q} */ //
/* end element point loads */ //
/* thermal loads ----------------------------------------------- */ //
/* ! local element coordinates ! */ //
/* {F} = [T]'{Q} */ //
/* end thermal loads	*/ //
/* debugging ...  check eqF's prior to asembly
  for (n=1; n<=nE; n++) {
	printf("n=%d ", n);
	for (l=1;l<=12;l++) {
		if (eqF_mech[lc][n][l] != 0)
		   printf(" eqF %d = %9.2e ", l, eqF_mech[lc][n][l] );
	}
	printf("\n");
  }
*/ //
// assemble all element equivalent loads into
// separate load vectors for mechanical and thermal loading
/* prescribed displacements ------------------------------------ */ //
/* end load-case loop */ //
//
func read_and_assemble_loads(fp *noarch.File, nN int, nE int, nL int, DoF int, xyz []vec3, L []float64, Le []float64, J1 []int, J2 []int, Ax []float32, Asy []float32, Asz []float32, Iy []float32, Iz []float32, E []float32, G []float32, p []float32, d []float32, gX []float32, gY []float32, gZ []float32, r []int, shear int, nF []int, nU []int, nW []int, nP []int, nT []int, nD []int, Q [][]float64, F_temp [][]float64, F_mech [][]float64, Fo []float64, U [][][]float32, W [][][]float32, P [][][]float32, T [][][]float32, Dp [][]float32, eqF_mech [][][]float64, eqF_temp [][][]float64, verbose int) {
	var hy float32
	var hz float32
	var x1 float32
	var x2 float32
	var w1 float32
	var w2 float32
	var Ln float64
	var R1o float64
	var R2o float64
	var f01 float64
	var f02 float64
	var Nx1 float64
	var Vy1 float64
	var Vz1 float64
	var Mx1 float64
	var My1 float64
	var Mz1 float64
	var Nx2 float64
	var Vy2 float64
	var Vz2 float64
	var Mx2 float64
	var My2 float64
	var Mz2 float64
	var Ksy float64
	var Ksz float64
	var a float64
	var b float64
	var t1 float64
	var t2 float64
	var t3 float64
	var t4 float64
	var t5 float64
	var t6 float64
	var t7 float64
	var t8 float64
	var t9 float64
	var i int
	var j int
	var l int
	var lc int
	var n int
	var n1 int
	var n2 int
	var sfrv int
	var errMsg []byte = make([]byte, 512)
	for j = 1; j <= DoF; j++ {
		Fo[j] = 0
	}
	for j = 1; j <= DoF; j++ {
		for lc = 1; lc <= nL; lc++ {
			F_mech[lc][j] = 0
			F_temp[lc][j] = F_mech[lc][j]
		}
	}
	for i = 1; i <= 12; i++ {
		for n = 1; n <= nE; n++ {
			for lc = 1; lc <= nL; lc++ {
				eqF_temp[lc][n][i] = 0
				eqF_mech[lc][n][i] = eqF_temp[lc][n][i]
			}
		}
	}
	for i = 1; i <= DoF; i++ {
		for lc = 1; lc <= nL; lc++ {
			Dp[lc][i] = float32(0)
		}
	}
	for i = 1; i <= nE; i++ {
		for j = 1; j <= 12; j++ {
			Q[i][j] = 0
		}
	}
	for lc = 1; lc <= nL; lc++ {
		if verbose != 0 {
			textColor('y', 'g', 'b', 'x')
			noarch.Fprintf(stdout, []byte(" load case %d of %d: \x00"), lc, nL)
			noarch.Fprintf(stdout, []byte("                                            \x00"))
			noarch.Fflush(stdout)
			color(0)
			noarch.Fprintf(stdout, []byte("\n\x00"))
		}
		sfrv = noarch.Fscanf(fp, []byte("%f %f %f\x00"), (*[100000000]float32)(unsafe.Pointer(&gX[lc]))[:], (*[100000000]float32)(unsafe.Pointer(&gY[lc]))[:], (*[100000000]float32)(unsafe.Pointer(&gZ[lc]))[:])
		if sfrv != 3 {
			sferr([]byte("gX gY gZ values in load data\x00"))
		}
		for n = 1; n <= nE; n++ {
			n1 = J1[n]
			n2 = J2[n]
			coord_trans(xyz, L[n], n1, n2, (*[100000000]float64)(unsafe.Pointer(&t1))[:], (*[100000000]float64)(unsafe.Pointer(&t2))[:], (*[100000000]float64)(unsafe.Pointer(&t3))[:], (*[100000000]float64)(unsafe.Pointer(&t4))[:], (*[100000000]float64)(unsafe.Pointer(&t5))[:], (*[100000000]float64)(unsafe.Pointer(&t6))[:], (*[100000000]float64)(unsafe.Pointer(&t7))[:], (*[100000000]float64)(unsafe.Pointer(&t8))[:], (*[100000000]float64)(unsafe.Pointer(&t9))[:], p[n])
			eqF_mech[lc][n][1] = float64(d[n]*Ax[n]) * L[n] * float64(gX[lc]) / 2
			eqF_mech[lc][n][2] = float64(d[n]*Ax[n]) * L[n] * float64(gY[lc]) / 2
			eqF_mech[lc][n][3] = float64(d[n]*Ax[n]) * L[n] * float64(gZ[lc]) / 2
			eqF_mech[lc][n][4] = float64(d[n]*Ax[n]) * L[n] * L[n] / 12 * ((-t4*t8+t5*t7)*float64(gY[lc]) + (-t4*t9+t6*t7)*float64(gZ[lc]))
			eqF_mech[lc][n][5] = float64(d[n]*Ax[n]) * L[n] * L[n] / 12 * ((-t5*t7+t4*t8)*float64(gX[lc]) + (-t5*t9+t6*t8)*float64(gZ[lc]))
			eqF_mech[lc][n][6] = float64(d[n]*Ax[n]) * L[n] * L[n] / 12 * ((-t6*t7+t4*t9)*float64(gX[lc]) + (-t6*t8+t5*t9)*float64(gY[lc]))
			eqF_mech[lc][n][7] = float64(d[n]*Ax[n]) * L[n] * float64(gX[lc]) / 2
			eqF_mech[lc][n][8] = float64(d[n]*Ax[n]) * L[n] * float64(gY[lc]) / 2
			eqF_mech[lc][n][9] = float64(d[n]*Ax[n]) * L[n] * float64(gZ[lc]) / 2
			eqF_mech[lc][n][10] = float64(d[n]*Ax[n]) * L[n] * L[n] / 12 * ((t4*t8-t5*t7)*float64(gY[lc]) + (t4*t9-t6*t7)*float64(gZ[lc]))
			eqF_mech[lc][n][11] = float64(d[n]*Ax[n]) * L[n] * L[n] / 12 * ((t5*t7-t4*t8)*float64(gX[lc]) + (t5*t9-t6*t8)*float64(gZ[lc]))
			eqF_mech[lc][n][12] = float64(d[n]*Ax[n]) * L[n] * L[n] / 12 * ((t6*t7-t4*t9)*float64(gX[lc]) + (t6*t8-t5*t9)*float64(gY[lc]))
		}
		sfrv = noarch.Fscanf(fp, []byte("%d\x00"), (*[100000000]int)(unsafe.Pointer(&nF[lc]))[:])
		if sfrv != 1 {
			sferr([]byte("nF value in load data\x00"))
		}
		if verbose != 0 {
			noarch.Fprintf(stdout, []byte("  number of loaded nodes \x00"))
			dots(stdout, 28)
			noarch.Fprintf(stdout, []byte(" nF = %3d\n\x00"), nF[lc])
		}
		for i = 1; i <= nF[lc]; i++ {
			sfrv = noarch.Fscanf(fp, []byte("%d\x00"), (*[100000000]int)(unsafe.Pointer(&j))[:])
			if sfrv != 1 {
				sferr([]byte("node value in point load data\x00"))
			}
			if j < 1 || j > nN {
				noarch.Sprintf(errMsg, []byte("\n  error in node load data: node number out of range ... Node : %d\n   Perhaps you did not specify %d node loads \n  or perhaps the Input Data file is missing expected data.\n\x00"), j, nF[lc])
				errorMsg(errMsg)
				os.Exit(121)
			}
			for l = 5; l >= 0; l-- {
				sfrv = noarch.Fscanf(fp, []byte("%lf\x00"), (*[100000000]float64)(unsafe.Pointer(&F_mech[lc][6*j-l]))[:])
				if sfrv != 1 {
					sferr([]byte("force value in point load data\x00"))
				}
			}
			if F_mech[lc][6*j-5] == 0 && F_mech[lc][6*j-4] == 0 && F_mech[lc][6*j-3] == 0 && F_mech[lc][6*j-2] == 0 && F_mech[lc][6*j-1] == 0 && F_mech[lc][6*j] == 0 {
				noarch.Fprintf(stderr, []byte("\n   Warning: All node loads applied at node %d  are zero\n\x00"), j)
			}
		}
		sfrv = noarch.Fscanf(fp, []byte("%d\x00"), (*[100000000]int)(unsafe.Pointer(&nU[lc]))[:])
		if sfrv != 1 {
			sferr([]byte("nU value in uniform load data\x00"))
		}
		if verbose != 0 {
			noarch.Fprintf(stdout, []byte("  number of uniformly distributed loads \x00"))
			dots(stdout, 13)
			noarch.Fprintf(stdout, []byte(" nU = %3d\n\x00"), nU[lc])
		}
		if nU[lc] < 0 || nU[lc] > nE {
			noarch.Fprintf(stderr, []byte("  number of uniformly distributed loads \x00"))
			dots(stderr, 13)
			noarch.Fprintf(stderr, []byte(" nU = %3d\n\x00"), nU[lc])
			noarch.Sprintf(errMsg, []byte("\n  error: valid ranges for nU is 0 ... %d \n\x00"), nE)
			errorMsg(errMsg)
			os.Exit(131)
		}
		for i = 1; i <= nU[lc]; i++ {
			sfrv = noarch.Fscanf(fp, []byte("%d\x00"), (*[100000000]int)(unsafe.Pointer(&n))[:])
			if sfrv != 1 {
				sferr([]byte("frame element number in uniform load data\x00"))
			}
			if n < 1 || n > nE {
				noarch.Sprintf(errMsg, []byte("\n  error in uniform distributed loads: element number %d is out of range\n\x00"), n)
				errorMsg(errMsg)
				os.Exit(132)
			}
			U[lc][i][1] = float32(float64(n))
			for l = 2; l <= 4; l++ {
				sfrv = noarch.Fscanf(fp, []byte("%f\x00"), (*[100000000]float32)(unsafe.Pointer(&U[lc][i][l]))[:])
				if sfrv != 1 {
					sferr([]byte("load value in uniform load data\x00"))
				}
			}
			if U[lc][i][2] == 0 && U[lc][i][3] == 0 && U[lc][i][4] == 0 {
				noarch.Fprintf(stderr, []byte("\n   Warning: All distributed loads applied to frame element %d  are zero\n\x00"), n)
			}
			Nx2 = float64(U[lc][i][2]) * Le[n] / 2
			Nx1 = Nx2
			Vy2 = float64(U[lc][i][3]) * Le[n] / 2
			Vy1 = Vy2
			Vz2 = float64(U[lc][i][4]) * Le[n] / 2
			Vz1 = Vz2
			Mx2 = 0
			Mx1 = Mx2
			My1 = float64(-U[lc][i][4]) * Le[n] * Le[n] / 12
			My2 = -My1
			Mz1 = float64(U[lc][i][3]) * Le[n] * Le[n] / 12
			Mz2 = -Mz1
			n1 = J1[n]
			n2 = J2[n]
			coord_trans(xyz, L[n], n1, n2, (*[100000000]float64)(unsafe.Pointer(&t1))[:], (*[100000000]float64)(unsafe.Pointer(&t2))[:], (*[100000000]float64)(unsafe.Pointer(&t3))[:], (*[100000000]float64)(unsafe.Pointer(&t4))[:], (*[100000000]float64)(unsafe.Pointer(&t5))[:], (*[100000000]float64)(unsafe.Pointer(&t6))[:], (*[100000000]float64)(unsafe.Pointer(&t7))[:], (*[100000000]float64)(unsafe.Pointer(&t8))[:], (*[100000000]float64)(unsafe.Pointer(&t9))[:], p[n])
			eqF_mech[lc][n][1] += Nx1*t1 + Vy1*t4 + Vz1*t7
			eqF_mech[lc][n][2] += Nx1*t2 + Vy1*t5 + Vz1*t8
			eqF_mech[lc][n][3] += Nx1*t3 + Vy1*t6 + Vz1*t9
			eqF_mech[lc][n][4] += Mx1*t1 + My1*t4 + Mz1*t7
			eqF_mech[lc][n][5] += Mx1*t2 + My1*t5 + Mz1*t8
			eqF_mech[lc][n][6] += Mx1*t3 + My1*t6 + Mz1*t9
			eqF_mech[lc][n][7] += Nx2*t1 + Vy2*t4 + Vz2*t7
			eqF_mech[lc][n][8] += Nx2*t2 + Vy2*t5 + Vz2*t8
			eqF_mech[lc][n][9] += Nx2*t3 + Vy2*t6 + Vz2*t9
			eqF_mech[lc][n][10] += Mx2*t1 + My2*t4 + Mz2*t7
			eqF_mech[lc][n][11] += Mx2*t2 + My2*t5 + Mz2*t8
			eqF_mech[lc][n][12] += Mx2*t3 + My2*t6 + Mz2*t9
		}
		sfrv = noarch.Fscanf(fp, []byte("%d\x00"), (*[100000000]int)(unsafe.Pointer(&nW[lc]))[:])
		if sfrv != 1 {
			sferr([]byte("nW value in load data\x00"))
		}
		if verbose != 0 {
			noarch.Fprintf(stdout, []byte("  number of trapezoidally distributed loads \x00"))
			dots(stdout, 9)
			noarch.Fprintf(stdout, []byte(" nW = %3d\n\x00"), nW[lc])
		}
		if nW[lc] < 0 || nW[lc] > 10*nE {
			noarch.Sprintf(errMsg, []byte("\n  error: valid ranges for nW is 0 ... %d \n\x00"), 10*nE)
			errorMsg(errMsg)
			os.Exit(140)
		}
		for i = 1; i <= nW[lc]; i++ {
			sfrv = noarch.Fscanf(fp, []byte("%d\x00"), (*[100000000]int)(unsafe.Pointer(&n))[:])
			if sfrv != 1 {
				sferr([]byte("frame element number in trapezoidal load data\x00"))
			}
			if n < 1 || n > nE {
				noarch.Sprintf(errMsg, []byte("\n  error in trapezoidally-distributed loads: element number %d is out of range\n\x00"), n)
				errorMsg(errMsg)
				os.Exit(141)
			}
			W[lc][i][1] = float32(float64(n))
			for l = 2; l <= 13; l++ {
				sfrv = noarch.Fscanf(fp, []byte("%f\x00"), (*[100000000]float32)(unsafe.Pointer(&W[lc][i][l]))[:])
				if sfrv != 1 {
					sferr([]byte("value in trapezoidal load data\x00"))
				}
			}
			Ln = L[n]
			if W[lc][i][4] == 0 && W[lc][i][5] == 0 && W[lc][i][8] == 0 && W[lc][i][9] == 0 && W[lc][i][12] == 0 && W[lc][i][13] == 0 {
				noarch.Fprintf(stderr, []byte("\n   Warning: All trapezoidal loads applied to frame element %d  are zero\n\x00"), n)
				noarch.Fprintf(stderr, []byte("     load case: %d , element %d , load %d\n \x00"), lc, n, i)
			}
			if W[lc][i][2] < 0 {
				noarch.Sprintf(errMsg, []byte("\n   error in x-axis trapezoidal loads, load case: %d , element %d , load %d\n  starting location = %f < 0\n\x00"), lc, n, i, float64(W[lc][i][2]))
				errorMsg(errMsg)
				os.Exit(142)
			}
			if W[lc][i][2] > W[lc][i][3] {
				noarch.Sprintf(errMsg, []byte("\n   error in x-axis trapezoidal loads, load case: %d , element %d , load %d\n  starting location = %f > ending location = %f \n\x00"), lc, n, i, float64(W[lc][i][2]), float64(W[lc][i][3]))
				errorMsg(errMsg)
				os.Exit(143)
			}
			if float64(W[lc][i][3]) > Ln {
				noarch.Sprintf(errMsg, []byte("\n   error in x-axis trapezoidal loads, load case: %d , element %d , load %d\n ending location = %f > L (%f) \n\x00"), lc, n, i, float64(W[lc][i][3]), Ln)
				errorMsg(errMsg)
				os.Exit(144)
			}
			if W[lc][i][6] < 0 {
				noarch.Sprintf(errMsg, []byte("\n   error in y-axis trapezoidal loads, load case: %d , element %d , load %d\n starting location = %f < 0\n\x00"), lc, n, i, float64(W[lc][i][6]))
				errorMsg(errMsg)
				os.Exit(142)
			}
			if W[lc][i][6] > W[lc][i][7] {
				noarch.Sprintf(errMsg, []byte("\n   error in y-axis trapezoidal loads, load case: %d , element %d , load %d\n starting location = %f > ending location = %f \n\x00"), lc, n, i, float64(W[lc][i][6]), float64(W[lc][i][7]))
				errorMsg(errMsg)
				os.Exit(143)
			}
			if float64(W[lc][i][7]) > Ln {
				noarch.Sprintf(errMsg, []byte("\n   error in y-axis trapezoidal loads, load case: %d , element %d , load %d\n ending location = %f > L (%f) \n\x00"), lc, n, i, float64(W[lc][i][7]), Ln)
				errorMsg(errMsg)
				os.Exit(144)
			}
			if W[lc][i][10] < 0 {
				noarch.Sprintf(errMsg, []byte("\n   error in z-axis trapezoidal loads, load case: %d , element %d , load %d\n starting location = %f < 0\n\x00"), lc, n, i, float64(W[lc][i][10]))
				errorMsg(errMsg)
				os.Exit(142)
			}
			if W[lc][i][10] > W[lc][i][11] {
				noarch.Sprintf(errMsg, []byte("\n   error in z-axis trapezoidal loads, load case: %d , element %d , load %d\n starting location = %f > ending location = %f \n\x00"), lc, n, i, float64(W[lc][i][10]), float64(W[lc][i][11]))
				errorMsg(errMsg)
				os.Exit(143)
			}
			if float64(W[lc][i][11]) > Ln {
				noarch.Sprintf(errMsg, []byte("\n   error in z-axis trapezoidal loads, load case: %d , element %d , load %d\n ending location = %f > L (%f) \n\x00"), lc, n, i, float64(W[lc][i][11]), Ln)
				errorMsg(errMsg)
				os.Exit(144)
			}
			if shear != 0 {
				Ksy = 12 * float64(E[n]) * float64(Iz[n]) / (float64(G[n]*Asy[n]) * Le[n] * Le[n])
				Ksz = 12 * float64(E[n]) * float64(Iy[n]) / (float64(G[n]*Asz[n]) * Le[n] * Le[n])
			} else {
				Ksz = 0
				Ksy = Ksz
			}
			x1 = W[lc][i][2]
			x2 = W[lc][i][3]
			w1 = W[lc][i][4]
			w2 = W[lc][i][5]
			Nx1 = (3*float64(w1+w2)*Ln*float64(x2-x1) - (2*float64(w2)+float64(w1))*float64(x2)*float64(x2) + float64((w2-w1)*x2*x1) + (2*float64(w1)+float64(w2))*float64(x1)*float64(x1)) / (6 * Ln)
			Nx2 = (-(2*float64(w1)+float64(w2))*float64(x1)*float64(x1) + (2*float64(w2)+float64(w1))*float64(x2)*float64(x2) - float64((w2-w1)*x1*x2)) / (6 * Ln)
			x1 = W[lc][i][6]
			x2 = W[lc][i][7]
			w1 = W[lc][i][8]
			w2 = W[lc][i][9]
			R1o = ((2*float64(w1)+float64(w2))*float64(x1)*float64(x1) - (float64(w1)+2*float64(w2))*float64(x2)*float64(x2) + 3*float64(w1+w2)*Ln*float64(x2-x1) - float64((w1-w2)*x1*x2)) / (6 * Ln)
			R2o = ((float64(w1)+2*float64(w2))*float64(x2)*float64(x2) + float64((w1-w2)*x1*x2) - (2*float64(w1)+float64(w2))*float64(x1)*float64(x1)) / (2 * Ln)
			f01 = (3*(float64(w2)+4*float64(w1))*float64(x1)*float64(x1)*float64(x1)*float64(x1) - 3*(float64(w1)+4*float64(w2))*float64(x2)*float64(x2)*float64(x2)*float64(x2) - 15*(float64(w2)+3*float64(w1))*Ln*float64(x1)*float64(x1)*float64(x1) + 0*(float64(w1)+0*float64(w2))*Ln*float64(x2)*float64(x2)*float64(x2) - 3*float64(w1-w2)*float64(x1)*float64(x2)*float64(x1*x1+x2*x2) + 20*(float64(w2)+2*float64(w1))*Ln*Ln*float64(x1)*float64(x1) - 0*(float64(w1)+0*float64(w2))*Ln*Ln*float64(x2)*float64(x2) + 15*float64(w1-w2)*Ln*float64(x1)*float64(x2)*float64(x1+x2) - 3*float64(w1-w2)*float64(x1)*float64(x1)*float64(x2)*float64(x2) - 1*float64(w1-w2)*Ln*Ln*float64(x1)*float64(x2)) / 360
			f02 = (3*(float64(w2)+4*float64(w1))*float64(x1)*float64(x1)*float64(x1)*float64(x1) - 3*(float64(w1)+4*float64(w2))*float64(x2)*float64(x2)*float64(x2)*float64(x2) - 3*float64(w1-w2)*float64(x1)*float64(x2)*float64(x1*x1+x2*x2) - 10*(float64(w2)+2*float64(w1))*Ln*Ln*float64(x1)*float64(x1) + 0*(float64(w1)+0*float64(w2))*Ln*Ln*float64(x2)*float64(x2) - 3*float64(w1-w2)*float64(x1)*float64(x1)*float64(x2)*float64(x2) + 1*float64(w1-w2)*Ln*Ln*float64(x1)*float64(x2)) / 360
			Mz1 = -(4*f01 + 2*f02 + Ksy*(f01-f02)) / (Ln * Ln * (1 + Ksy))
			Mz2 = -(2*f01 + 4*f02 - Ksy*(f01-f02)) / (Ln * Ln * (1 + Ksy))
			Vy1 = R1o + Mz1/Ln + Mz2/Ln
			Vy2 = R2o - Mz1/Ln - Mz2/Ln
			x1 = W[lc][i][10]
			x2 = W[lc][i][11]
			w1 = W[lc][i][12]
			w2 = W[lc][i][13]
			R1o = ((2*float64(w1)+float64(w2))*float64(x1)*float64(x1) - (float64(w1)+2*float64(w2))*float64(x2)*float64(x2) + 3*float64(w1+w2)*Ln*float64(x2-x1) - float64((w1-w2)*x1*x2)) / (6 * Ln)
			R2o = ((float64(w1)+2*float64(w2))*float64(x2)*float64(x2) + float64((w1-w2)*x1*x2) - (2*float64(w1)+float64(w2))*float64(x1)*float64(x1)) / (2 * Ln)
			f01 = (3*(float64(w2)+4*float64(w1))*float64(x1)*float64(x1)*float64(x1)*float64(x1) - 3*(float64(w1)+4*float64(w2))*float64(x2)*float64(x2)*float64(x2)*float64(x2) - 15*(float64(w2)+3*float64(w1))*Ln*float64(x1)*float64(x1)*float64(x1) + 0*(float64(w1)+0*float64(w2))*Ln*float64(x2)*float64(x2)*float64(x2) - 3*float64(w1-w2)*float64(x1)*float64(x2)*float64(x1*x1+x2*x2) + 20*(float64(w2)+2*float64(w1))*Ln*Ln*float64(x1)*float64(x1) - 0*(float64(w1)+0*float64(w2))*Ln*Ln*float64(x2)*float64(x2) + 15*float64(w1-w2)*Ln*float64(x1)*float64(x2)*float64(x1+x2) - 3*float64(w1-w2)*float64(x1)*float64(x1)*float64(x2)*float64(x2) - 1*float64(w1-w2)*Ln*Ln*float64(x1)*float64(x2)) / 360
			f02 = (3*(float64(w2)+4*float64(w1))*float64(x1)*float64(x1)*float64(x1)*float64(x1) - 3*(float64(w1)+4*float64(w2))*float64(x2)*float64(x2)*float64(x2)*float64(x2) - 3*float64(w1-w2)*float64(x1)*float64(x2)*float64(x1*x1+x2*x2) - 10*(float64(w2)+2*float64(w1))*Ln*Ln*float64(x1)*float64(x1) + 0*(float64(w1)+0*float64(w2))*Ln*Ln*float64(x2)*float64(x2) - 3*float64(w1-w2)*float64(x1)*float64(x1)*float64(x2)*float64(x2) + 1*float64(w1-w2)*Ln*Ln*float64(x1)*float64(x2)) / 360
			My1 = (4*f01 + 2*f02 + Ksz*(f01-f02)) / (Ln * Ln * (1 + Ksz))
			My2 = (2*f01 + 4*f02 - Ksz*(f01-f02)) / (Ln * Ln * (1 + Ksz))
			Vz1 = R1o - My1/Ln - My2/Ln
			Vz2 = R2o + My1/Ln + My2/Ln
			n1 = J1[n]
			n2 = J2[n]
			coord_trans(xyz, Ln, n1, n2, (*[100000000]float64)(unsafe.Pointer(&t1))[:], (*[100000000]float64)(unsafe.Pointer(&t2))[:], (*[100000000]float64)(unsafe.Pointer(&t3))[:], (*[100000000]float64)(unsafe.Pointer(&t4))[:], (*[100000000]float64)(unsafe.Pointer(&t5))[:], (*[100000000]float64)(unsafe.Pointer(&t6))[:], (*[100000000]float64)(unsafe.Pointer(&t7))[:], (*[100000000]float64)(unsafe.Pointer(&t8))[:], (*[100000000]float64)(unsafe.Pointer(&t9))[:], p[n])
			eqF_mech[lc][n][1] += Nx1*t1 + Vy1*t4 + Vz1*t7
			eqF_mech[lc][n][2] += Nx1*t2 + Vy1*t5 + Vz1*t8
			eqF_mech[lc][n][3] += Nx1*t3 + Vy1*t6 + Vz1*t9
			eqF_mech[lc][n][4] += Mx1*t1 + My1*t4 + Mz1*t7
			eqF_mech[lc][n][5] += Mx1*t2 + My1*t5 + Mz1*t8
			eqF_mech[lc][n][6] += Mx1*t3 + My1*t6 + Mz1*t9
			eqF_mech[lc][n][7] += Nx2*t1 + Vy2*t4 + Vz2*t7
			eqF_mech[lc][n][8] += Nx2*t2 + Vy2*t5 + Vz2*t8
			eqF_mech[lc][n][9] += Nx2*t3 + Vy2*t6 + Vz2*t9
			eqF_mech[lc][n][10] += Mx2*t1 + My2*t4 + Mz2*t7
			eqF_mech[lc][n][11] += Mx2*t2 + My2*t5 + Mz2*t8
			eqF_mech[lc][n][12] += Mx2*t3 + My2*t6 + Mz2*t9
		}
		sfrv = noarch.Fscanf(fp, []byte("%d\x00"), (*[100000000]int)(unsafe.Pointer(&nP[lc]))[:])
		if sfrv != 1 {
			sferr([]byte("nP value load data\x00"))
		}
		if verbose != 0 {
			noarch.Fprintf(stdout, []byte("  number of concentrated frame element point loads \x00"))
			dots(stdout, 2)
			noarch.Fprintf(stdout, []byte(" nP = %3d\n\x00"), nP[lc])
		}
		if nP[lc] < 0 || nP[lc] > 10*nE {
			noarch.Fprintf(stderr, []byte("  number of concentrated frame element point loads \x00"))
			dots(stderr, 3)
			noarch.Fprintf(stderr, []byte(" nP = %3d\n\x00"), nP[lc])
			noarch.Sprintf(errMsg, []byte("\n  error: valid ranges for nP is 0 ... %d \n\x00"), 10*nE)
			errorMsg(errMsg)
			os.Exit(150)
		}
		for i = 1; i <= nP[lc]; i++ {
			sfrv = noarch.Fscanf(fp, []byte("%d\x00"), (*[100000000]int)(unsafe.Pointer(&n))[:])
			if sfrv != 1 {
				sferr([]byte("frame element number value point load data\x00"))
			}
			if n < 1 || n > nE {
				noarch.Sprintf(errMsg, []byte("\n   error in internal point loads: frame element number %d is out of range\n\x00"), n)
				errorMsg(errMsg)
				os.Exit(151)
			}
			P[lc][i][1] = float32(float64(n))
			for l = 2; l <= 5; l++ {
				sfrv = noarch.Fscanf(fp, []byte("%f\x00"), (*[100000000]float32)(unsafe.Pointer(&P[lc][i][l]))[:])
				if sfrv != 1 {
					sferr([]byte("value in point load data\x00"))
				}
			}
			a = float64(P[lc][i][5])
			b = L[n] - a
			if a < 0 || L[n] < a || b < 0 || L[n] < b {
				noarch.Sprintf(errMsg, []byte("\n  error in point load data: Point load coord. out of range\n   Frame element number: %d  L: %lf  load coord.: %lf\n\x00"), n, L[n], float64(P[lc][i][5]))
				errorMsg(errMsg)
				os.Exit(152)
			}
			if shear != 0 {
				Ksy = 12 * float64(E[n]) * float64(Iz[n]) / (float64(G[n]*Asy[n]) * Le[n] * Le[n])
				Ksz = 12 * float64(E[n]) * float64(Iy[n]) / (float64(G[n]*Asz[n]) * Le[n] * Le[n])
			} else {
				Ksz = 0
				Ksy = Ksz
			}
			Ln = L[n]
			Nx1 = float64(P[lc][i][2]) * a / Ln
			Nx2 = float64(P[lc][i][2]) * b / Ln
			Vy1 = 1/(1+Ksz)*float64(P[lc][i][3])*b*b*(3*a+b)/(Ln*Ln*Ln) + Ksz/(1+Ksz)*float64(P[lc][i][3])*b/Ln
			Vy2 = 1/(1+Ksz)*float64(P[lc][i][3])*a*a*(3*b+a)/(Ln*Ln*Ln) + Ksz/(1+Ksz)*float64(P[lc][i][3])*a/Ln
			Vz1 = 1/(1+Ksy)*float64(P[lc][i][4])*b*b*(3*a+b)/(Ln*Ln*Ln) + Ksy/(1+Ksy)*float64(P[lc][i][4])*b/Ln
			Vz2 = 1/(1+Ksy)*float64(P[lc][i][4])*a*a*(3*b+a)/(Ln*Ln*Ln) + Ksy/(1+Ksy)*float64(P[lc][i][4])*a/Ln
			Mx2 = 0
			Mx1 = Mx2
			My1 = -(1/(1+Ksy))*float64(P[lc][i][4])*a*b*b/(Ln*Ln) - Ksy/(1+Ksy)*float64(P[lc][i][4])*a*b/(2*Ln)
			My2 = 1/(1+Ksy)*float64(P[lc][i][4])*a*a*b/(Ln*Ln) + Ksy/(1+Ksy)*float64(P[lc][i][4])*a*b/(2*Ln)
			Mz1 = 1/(1+Ksz)*float64(P[lc][i][3])*a*b*b/(Ln*Ln) + Ksz/(1+Ksz)*float64(P[lc][i][3])*a*b/(2*Ln)
			Mz2 = -(1/(1+Ksz))*float64(P[lc][i][3])*a*a*b/(Ln*Ln) - Ksz/(1+Ksz)*float64(P[lc][i][3])*a*b/(2*Ln)
			n1 = J1[n]
			n2 = J2[n]
			coord_trans(xyz, Ln, n1, n2, (*[100000000]float64)(unsafe.Pointer(&t1))[:], (*[100000000]float64)(unsafe.Pointer(&t2))[:], (*[100000000]float64)(unsafe.Pointer(&t3))[:], (*[100000000]float64)(unsafe.Pointer(&t4))[:], (*[100000000]float64)(unsafe.Pointer(&t5))[:], (*[100000000]float64)(unsafe.Pointer(&t6))[:], (*[100000000]float64)(unsafe.Pointer(&t7))[:], (*[100000000]float64)(unsafe.Pointer(&t8))[:], (*[100000000]float64)(unsafe.Pointer(&t9))[:], p[n])
			eqF_mech[lc][n][1] += Nx1*t1 + Vy1*t4 + Vz1*t7
			eqF_mech[lc][n][2] += Nx1*t2 + Vy1*t5 + Vz1*t8
			eqF_mech[lc][n][3] += Nx1*t3 + Vy1*t6 + Vz1*t9
			eqF_mech[lc][n][4] += Mx1*t1 + My1*t4 + Mz1*t7
			eqF_mech[lc][n][5] += Mx1*t2 + My1*t5 + Mz1*t8
			eqF_mech[lc][n][6] += Mx1*t3 + My1*t6 + Mz1*t9
			eqF_mech[lc][n][7] += Nx2*t1 + Vy2*t4 + Vz2*t7
			eqF_mech[lc][n][8] += Nx2*t2 + Vy2*t5 + Vz2*t8
			eqF_mech[lc][n][9] += Nx2*t3 + Vy2*t6 + Vz2*t9
			eqF_mech[lc][n][10] += Mx2*t1 + My2*t4 + Mz2*t7
			eqF_mech[lc][n][11] += Mx2*t2 + My2*t5 + Mz2*t8
			eqF_mech[lc][n][12] += Mx2*t3 + My2*t6 + Mz2*t9
		}
		sfrv = noarch.Fscanf(fp, []byte("%d\x00"), (*[100000000]int)(unsafe.Pointer(&nT[lc]))[:])
		if sfrv != 1 {
			sferr([]byte("nT value in load data\x00"))
		}
		if verbose != 0 {
			noarch.Fprintf(stdout, []byte("  number of temperature changes \x00"))
			dots(stdout, 21)
			noarch.Fprintf(stdout, []byte(" nT = %3d\n\x00"), nT[lc])
		}
		if nT[lc] < 0 || nT[lc] > nE {
			noarch.Fprintf(stderr, []byte("  number of temperature changes \x00"))
			dots(stderr, 21)
			noarch.Fprintf(stderr, []byte(" nT = %3d\n\x00"), nT[lc])
			noarch.Sprintf(errMsg, []byte("\n  error: valid ranges for nT is 0 ... %d \n\x00"), nE)
			errorMsg(errMsg)
			os.Exit(160)
		}
		for i = 1; i <= nT[lc]; i++ {
			sfrv = noarch.Fscanf(fp, []byte("%d\x00"), (*[100000000]int)(unsafe.Pointer(&n))[:])
			if sfrv != 1 {
				sferr([]byte("frame element number in temperature load data\x00"))
			}
			if n < 1 || n > nE {
				noarch.Sprintf(errMsg, []byte("\n  error in temperature loads: frame element number %d is out of range\n\x00"), n)
				errorMsg(errMsg)
				os.Exit(161)
			}
			T[lc][i][1] = float32(float64(n))
			for l = 2; l <= 8; l++ {
				sfrv = noarch.Fscanf(fp, []byte("%f\x00"), (*[100000000]float32)(unsafe.Pointer(&T[lc][i][l]))[:])
				if sfrv != 1 {
					sferr([]byte("value in temperature load data\x00"))
				}
			}
			a = float64(T[lc][i][2])
			hy = T[lc][i][3]
			hz = T[lc][i][4]
			if hy < 0 || hz < 0 {
				noarch.Sprintf(errMsg, []byte("\n  error in thermal load data: section dimension < 0\n   Frame element number: %d  hy: %f  hz: %f\n\x00"), n, float64(hy), float64(hz))
				errorMsg(errMsg)
				os.Exit(162)
			}
			Nx2 = a * (1 / 4) * float64(T[lc][i][5]+T[lc][i][6]+T[lc][i][7]+T[lc][i][8]) * float64(E[n]) * float64(Ax[n])
			Nx1 = -Nx2
			Vz2 = 0
			Vz1 = Vz2
			Vy2 = Vz1
			Vy1 = Vy2
			Mx2 = 0
			Mx1 = Mx2
			My1 = a / float64(hz) * float64(T[lc][i][8]-T[lc][i][7]) * float64(E[n]) * float64(Iy[n])
			My2 = -My1
			Mz1 = a / float64(hy) * float64(T[lc][i][5]-T[lc][i][6]) * float64(E[n]) * float64(Iz[n])
			Mz2 = -Mz1
			n1 = J1[n]
			n2 = J2[n]
			coord_trans(xyz, L[n], n1, n2, (*[100000000]float64)(unsafe.Pointer(&t1))[:], (*[100000000]float64)(unsafe.Pointer(&t2))[:], (*[100000000]float64)(unsafe.Pointer(&t3))[:], (*[100000000]float64)(unsafe.Pointer(&t4))[:], (*[100000000]float64)(unsafe.Pointer(&t5))[:], (*[100000000]float64)(unsafe.Pointer(&t6))[:], (*[100000000]float64)(unsafe.Pointer(&t7))[:], (*[100000000]float64)(unsafe.Pointer(&t8))[:], (*[100000000]float64)(unsafe.Pointer(&t9))[:], p[n])
			eqF_temp[lc][n][1] += Nx1*t1 + Vy1*t4 + Vz1*t7
			eqF_temp[lc][n][2] += Nx1*t2 + Vy1*t5 + Vz1*t8
			eqF_temp[lc][n][3] += Nx1*t3 + Vy1*t6 + Vz1*t9
			eqF_temp[lc][n][4] += Mx1*t1 + My1*t4 + Mz1*t7
			eqF_temp[lc][n][5] += Mx1*t2 + My1*t5 + Mz1*t8
			eqF_temp[lc][n][6] += Mx1*t3 + My1*t6 + Mz1*t9
			eqF_temp[lc][n][7] += Nx2*t1 + Vy2*t4 + Vz2*t7
			eqF_temp[lc][n][8] += Nx2*t2 + Vy2*t5 + Vz2*t8
			eqF_temp[lc][n][9] += Nx2*t3 + Vy2*t6 + Vz2*t9
			eqF_temp[lc][n][10] += Mx2*t1 + My2*t4 + Mz2*t7
			eqF_temp[lc][n][11] += Mx2*t2 + My2*t5 + Mz2*t8
			eqF_temp[lc][n][12] += Mx2*t3 + My2*t6 + Mz2*t9
		}
		for n = 1; n <= nE; n++ {
			n1 = J1[n]
			n2 = J2[n]
			for i = 1; i <= 6; i++ {
				F_mech[lc][6*n1-6+i] += eqF_mech[lc][n][i]
			}
			for i = 7; i <= 12; i++ {
				F_mech[lc][6*n2-12+i] += eqF_mech[lc][n][i]
			}
			for i = 1; i <= 6; i++ {
				F_temp[lc][6*n1-6+i] += eqF_temp[lc][n][i]
			}
			for i = 7; i <= 12; i++ {
				F_temp[lc][6*n2-12+i] += eqF_temp[lc][n][i]
			}
		}
		sfrv = noarch.Fscanf(fp, []byte("%d\x00"), (*[100000000]int)(unsafe.Pointer(&nD[lc]))[:])
		if sfrv != 1 {
			sferr([]byte("nD value in load data\x00"))
		}
		if verbose != 0 {
			noarch.Fprintf(stdout, []byte("  number of prescribed displacements \x00"))
			dots(stdout, 16)
			noarch.Fprintf(stdout, []byte(" nD = %3d\n\x00"), nD[lc])
		}
		for i = 1; i <= nD[lc]; i++ {
			sfrv = noarch.Fscanf(fp, []byte("%d\x00"), (*[100000000]int)(unsafe.Pointer(&j))[:])
			if sfrv != 1 {
				sferr([]byte("node number value in prescribed displacement data\x00"))
			}
			for l = 5; l >= 0; l-- {
				sfrv = noarch.Fscanf(fp, []byte("%f\x00"), (*[100000000]float32)(unsafe.Pointer(&Dp[lc][6*j-l]))[:])
				if sfrv != 1 {
					sferr([]byte("prescribed displacement value\x00"))
				}
				if r[6*j-l] == 0 && float64(Dp[lc][6*j-l]) != 0 {
					noarch.Sprintf(errMsg, []byte(" Initial displacements can be prescribed only at restrained coordinates\n  node: %d  dof: %d  r: %d\n\x00"), j, 6-l, r[6*j-l])
					errorMsg(errMsg)
					os.Exit(171)
				}
			}
		}
	}
}

// read_mass_data - transpiled function from  $GOPATH/src/github.com/Konstantin8105/History_frame3DD/src/frame3dd_io.c:1465
/*
 * READ_MASS_DATA  -  read element densities and extra inertial mass data	16aug01
 */ //
/*	double	ms = 0.0; */ //
/* *scanf return value	*/ //
/* number of nodes with extra inertias */ //
/* number of frame elements with extra beam mass */ //
/* calculate the total mass and the structural mass */ //
/* check inertia data	*/ //
/*	for (m=1;m<=nE;m++) ms += EMs[m]; // consistent mass doesn't agree  */ //
/*	if ( ms > 0.0 )	    *lump = 1;    // with concentrated masses, EMs  */ //
/* the length of the base_file */ //
/* find the last '.' in base_file */ //
/* end base_file at the last '.' */ //
/* find the last '/' or '\' in base_file */ //
//
func read_mass_data(fp *noarch.File, OUT_file []byte, nN int, nE int, nI []int, nX []int, d []float32, EMs []float32, NMs []float32, NMx []float32, NMy []float32, NMz []float32, L []float64, Ax []float32, total_mass []float64, struct_mass []float64, nM []int, Mmethod []int, modal_flag int, lump []int, lump_flag int, tol []float64, tol_flag float64, shift []float64, shift_flag float64, exagg_modal []float64, modepath []byte, anim []int, pan []float32, pan_flag float32, verbose int, debug int) {
	var i int
	var j int
	var jnt int
	var m int
	var b int
	var nA int
	var full_len int
	var len int
	var sfrv int
	var base_file []byte = []byte("EMPTY_BASE\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00")
	var mode_file []byte = []byte("EMPTY_MODE\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00")
	var errMsg []byte = make([]byte, 512)
	struct_mass[0] = 0
	total_mass[0] = struct_mass[0]
	sfrv = noarch.Fscanf(fp, []byte("%d\x00"), nM)
	if sfrv != 1 {
		sferr([]byte("nM value in mass data\x00"))
	}
	if verbose != 0 {
		noarch.Fprintf(stdout, []byte(" number of dynamic modes \x00"))
		dots(stdout, 28)
		noarch.Fprintf(stdout, []byte(" nM = %3d\n\x00"), nM[0])
	}
	if nM[0] < 1 || sfrv != 1 {
		nM[0] = 0
		return
	}
	sfrv = noarch.Fscanf(fp, []byte("%d\x00"), Mmethod)
	if sfrv != 1 {
		sferr([]byte("Mmethod value in mass data\x00"))
	}
	if modal_flag != -1 {
		Mmethod[0] = modal_flag
	}
	if verbose != 0 {
		noarch.Fprintf(stdout, []byte(" modal analysis method \x00"))
		dots(stdout, 30)
		noarch.Fprintf(stdout, []byte(" %3d \x00"), Mmethod[0])
		if Mmethod[0] == 1 {
			noarch.Fprintf(stdout, []byte(" (Subspace-Jacobi)\n\x00"))
		}
		if Mmethod[0] == 2 {
			noarch.Fprintf(stdout, []byte(" (Stodola)\n\x00"))
		}
	}
	sfrv = noarch.Fscanf(fp, []byte("%d\x00"), lump)
	if sfrv != 1 {
		sferr([]byte("lump value in mass data\x00"))
	}
	sfrv = noarch.Fscanf(fp, []byte("%lf\x00"), tol)
	if sfrv != 1 {
		sferr([]byte("tol value in mass data\x00"))
	}
	sfrv = noarch.Fscanf(fp, []byte("%lf\x00"), shift)
	if sfrv != 1 {
		sferr([]byte("shift value in mass data\x00"))
	}
	sfrv = noarch.Fscanf(fp, []byte("%lf\x00"), exagg_modal)
	if sfrv != 1 {
		sferr([]byte("exagg_modal value in mass data\x00"))
	}
	if lump_flag != -1 {
		lump[0] = lump_flag
	}
	if tol_flag != -1 {
		tol[0] = tol_flag
	}
	if shift_flag != -1 {
		shift[0] = shift_flag
	}
	sfrv = noarch.Fscanf(fp, []byte("%d\x00"), nI)
	if sfrv != 1 {
		sferr([]byte("nI value in mass data\x00"))
	}
	if verbose != 0 {
		noarch.Fprintf(stdout, []byte(" number of nodes with extra lumped inertia \x00"))
		dots(stdout, 10)
		noarch.Fprintf(stdout, []byte(" nI = %3d\n\x00"), nI[0])
	}
	for j = 1; j <= nI[0]; j++ {
		sfrv = noarch.Fscanf(fp, []byte("%d\x00"), (*[100000000]int)(unsafe.Pointer(&jnt))[:])
		if sfrv != 1 {
			sferr([]byte("node value in extra node mass data\x00"))
		}
		if jnt < 1 || jnt > nN {
			noarch.Sprintf(errMsg, []byte("\n  error in node mass data: node number out of range    Node : %d  \n   Perhaps you did not specify %d extra masses \n   or perhaps the Input Data file is missing expected data.\n\x00"), jnt, nI[0])
			errorMsg(errMsg)
			os.Exit(86)
		}
		sfrv = noarch.Fscanf(fp, []byte("%f %f %f %f\x00"), (*[100000000]float32)(unsafe.Pointer(&NMs[jnt]))[:], (*[100000000]float32)(unsafe.Pointer(&NMx[jnt]))[:], (*[100000000]float32)(unsafe.Pointer(&NMy[jnt]))[:], (*[100000000]float32)(unsafe.Pointer(&NMz[jnt]))[:])
		if sfrv != 4 {
			sferr([]byte("node inertia in extra mass data\x00"))
		}
		total_mass[0] += float64(NMs[jnt])
		if NMs[jnt] == 0 && NMx[jnt] == 0 && NMy[jnt] == 0 && NMz[jnt] == 0 {
			noarch.Fprintf(stderr, []byte("\n  Warning: All extra node inertia at node %d  are zero\n\x00"), jnt)
		}
	}
	sfrv = noarch.Fscanf(fp, []byte("%d\x00"), nX)
	if sfrv != 1 {
		sferr([]byte("nX value in mass data\x00"))
	}
	if verbose != 0 {
		noarch.Fprintf(stdout, []byte(" number of frame elements with extra mass \x00"))
		dots(stdout, 11)
		noarch.Fprintf(stdout, []byte(" nX = %3d\n\x00"), nX[0])
		if sfrv != 1 {
			sferr([]byte("element value in extra element mass data\x00"))
		}
	}
	for m = 1; m <= nX[0]; m++ {
		sfrv = noarch.Fscanf(fp, []byte("%d\x00"), (*[100000000]int)(unsafe.Pointer(&b))[:])
		if sfrv != 1 {
			sferr([]byte("element number in extra element mass data\x00"))
		}
		if b < 1 || b > nE {
			noarch.Sprintf(errMsg, []byte("\n  error in element mass data: element number out of range   Element: %d  \n   Perhaps you did not specify %d extra masses \n   or perhaps the Input Data file is missing expected data.\n\x00"), b, nX[0])
			errorMsg(errMsg)
			os.Exit(87)
		}
		sfrv = noarch.Fscanf(fp, []byte("%f\x00"), (*[100000000]float32)(unsafe.Pointer(&EMs[b]))[:])
		if sfrv != 1 {
			sferr([]byte("extra element mass value in mass data\x00"))
		}
	}
	for b = 1; b <= nE; b++ {
		total_mass[0] += float64(d[b]*Ax[b])*L[b] + float64(EMs[b])
		struct_mass[0] += float64(d[b]*Ax[b]) * L[b]
	}
	for m = 1; m <= nE; m++ {
		if float64(d[m]) < 0 || float64(EMs[m]) < 0 || float64(d[m]+EMs[m]) <= 0 {
			noarch.Sprintf(errMsg, []byte("\n  error: Non-positive mass or density\n  d[%d]= %f  EMs[%d]= %f\n\x00"), m, float64(d[m]), m, float64(EMs[m]))
			errorMsg(errMsg)
			os.Exit(88)
		}
	}
	if verbose != 0 {
		noarch.Fprintf(stdout, []byte(" structural mass \x00"))
		dots(stdout, 36)
		noarch.Fprintf(stdout, []byte("  %12.4e\n\x00"), struct_mass[0])
		noarch.Fprintf(stdout, []byte(" total mass \x00"))
		dots(stdout, 41)
		noarch.Fprintf(stdout, []byte("  %12.4e\n\x00"), total_mass[0])
	}
	sfrv = noarch.Fscanf(fp, []byte("%d\x00"), (*[100000000]int)(unsafe.Pointer(&nA))[:])
	if sfrv != 1 {
		sferr([]byte("nA value in mode animation data\x00"))
	}
	if verbose != 0 {
		noarch.Fprintf(stdout, []byte(" number of modes to be animated \x00"))
		dots(stdout, 21)
		noarch.Fprintf(stdout, []byte(" nA = %3d\n\x00"), nA)
	}
	if nA > 20 {
		noarch.Fprintf(stderr, []byte(" nA = %d, only 100 or fewer modes may be animated\n\x00"), nA)
	}
	for m = 0; m < 20; m++ {
		anim[m] = 0
	}
	for m = 1; m <= nA; m++ {
		sfrv = noarch.Fscanf(fp, []byte("%d\x00"), (*[100000000]int)(unsafe.Pointer(&anim[m]))[:])
		if sfrv != 1 {
			sferr([]byte("mode number in mode animation data\x00"))
		}
	}
	sfrv = noarch.Fscanf(fp, []byte("%f\x00"), pan)
	if sfrv != 1 {
		sferr([]byte("pan value in mode animation data\x00"))
	}
	if float64(pan_flag) != -1 {
		pan[0] = pan_flag
	}
	if verbose != 0 {
		noarch.Fprintf(stdout, []byte(" pan rate \x00"))
		dots(stdout, 43)
		noarch.Fprintf(stdout, []byte(" %8.3f\n\x00"), float64(pan[0]))
	}
	noarch.Strcpy(base_file, OUT_file)
	for int(base_file[func() int {
		defer func() {
			len++
		}()
		return len
	}()]) != int('\x00') {
	}
	full_len = len
	for int(base_file[func() int {
		defer func() {
			len--
		}()
		return len
	}()]) != int('.') && len > 0 {
	}
	if len == 0 {
		len = full_len
	}
	base_file[func() int {
		len++
		return len
	}()] = '\x00'
	for int(base_file[len]) != int('/') && int(base_file[len]) != int('\\') && len > 0 {
		len--
	}
	i = 0
	for int(base_file[len]) != int('\x00') {
		mode_file[func() int {
			defer func() {
				i++
			}()
			return i
		}()] = base_file[func() int {
			defer func() {
				len++
			}()
			return len
		}()]
	}
	mode_file[i] = '\x00'
	noarch.Strcat(mode_file, []byte("-m\x00"))
	output_path(mode_file, modepath, 512, nil)
}

// read_condensation_data - transpiled function from  $GOPATH/src/github.com/Konstantin8105/History_frame3DD/src/frame3dd_io.c:1668
/*
 * READ_CONDENSE   -  read matrix condensation information 	        30aug01
 */ //
/* *scanf return value */ //
/* default */ //
/* error check */ //
//
func read_condensation_data(fp *noarch.File, nN int, nM int, nC []int, Cdof []int, Cmethod []int, condense_flag int, c []int, m []int, verbose int) {
	var i int
	var j int
	var k int
	var cm [][]int
	var sfrv int
	var errMsg []byte = make([]byte, 512)
	Cdof[0] = 0
	nC[0] = Cdof[0]
	Cmethod[0] = nC[0]
	if (func() int {
		sfrv = noarch.Fscanf(fp, []byte("%d\x00"), Cmethod)
		return sfrv
	}()) != 1 {
		Cdof[0] = 0
		nC[0] = Cdof[0]
		Cmethod[0] = nC[0]
		if verbose != 0 {
			noarch.Fprintf(stdout, []byte(" missing matrix condensation data \n\x00"))
		}
		return
	}
	if condense_flag != -1 {
		Cmethod[0] = condense_flag
	}
	if Cmethod[0] <= 0 {
		if verbose != 0 {
			noarch.Fprintf(stdout, []byte(" Cmethod = %d : no matrix condensation \n\x00"), Cmethod[0])
		}
		Cdof[0] = 0
		nC[0] = Cdof[0]
		Cmethod[0] = nC[0]
		return
	}
	if Cmethod[0] > 3 {
		Cmethod[0] = 1
	}
	if verbose != 0 {
		noarch.Fprintf(stdout, []byte(" condensation method \x00"))
		dots(stdout, 32)
		noarch.Fprintf(stdout, []byte(" %d \x00"), Cmethod[0])
		if Cmethod[0] == 1 {
			noarch.Fprintf(stdout, []byte(" (static only) \n\x00"))
		}
		if Cmethod[0] == 2 {
			noarch.Fprintf(stdout, []byte(" (Guyan) \n\x00"))
		}
		if Cmethod[0] == 3 {
			noarch.Fprintf(stdout, []byte(" (dynamic) \n\x00"))
		}
	}
	if (func() int {
		sfrv = noarch.Fscanf(fp, []byte("%d\x00"), nC)
		return sfrv
	}()) != 1 {
		Cdof[0] = 0
		nC[0] = Cdof[0]
		Cmethod[0] = nC[0]
		if verbose != 0 {
			noarch.Fprintf(stdout, []byte(" missing matrix condensation data \n\x00"))
		}
		return
	}
	if verbose != 0 {
		noarch.Fprintf(stdout, []byte(" number of nodes with condensed DoF's \x00"))
		dots(stdout, 15)
		noarch.Fprintf(stdout, []byte(" nC = %3d\n\x00"), nC[0])
	}
	if nC[0] > nN {
		noarch.Sprintf(errMsg, []byte("\n  error in matrix condensation data: \n error: nC > nN ... nC=%d; nN=%d;\n The number of nodes with condensed DoF's may not exceed the total number of nodes.\n\x00"), nC[0], nN)
		errorMsg(errMsg)
		os.Exit(90)
	}
	cm = imatrix(1, int32(nC[0]), 1, 7)
	for i = 1; i <= nC[0]; i++ {
		sfrv = noarch.Fscanf(fp, []byte("%d %d %d %d %d %d %d\x00"), (*[100000000]int)(unsafe.Pointer(&cm[i][1]))[:], (*[100000000]int)(unsafe.Pointer(&cm[i][2]))[:], (*[100000000]int)(unsafe.Pointer(&cm[i][3]))[:], (*[100000000]int)(unsafe.Pointer(&cm[i][4]))[:], (*[100000000]int)(unsafe.Pointer(&cm[i][5]))[:], (*[100000000]int)(unsafe.Pointer(&cm[i][6]))[:], (*[100000000]int)(unsafe.Pointer(&cm[i][7]))[:])
		if sfrv != 7 {
			sferr([]byte("DoF numbers in condensation data\x00"))
		}
		if cm[i][1] < 1 || cm[i][1] > nN {
			noarch.Sprintf(errMsg, []byte("\n  error in matrix condensation data: \n  condensed node number out of range\n  cj[%d] = %d  ... nN = %d  \n\x00"), i, cm[i][1], nN)
			errorMsg(errMsg)
			os.Exit(91)
		}
	}
	for i = 1; i <= nC[0]; i++ {
		for j = 2; j <= 7; j++ {
			if cm[i][j] != 0 {
				Cdof[0]++
			}
		}
	}
	k = 1
	for i = 1; i <= nC[0]; i++ {
		for j = 2; j <= 7; j++ {
			if cm[i][j] != 0 {
				c[k] = 6*(cm[i][1]-1) + j - 1
				k++
			}
		}
	}
	for i = 1; i <= Cdof[0]; i++ {
		sfrv = noarch.Fscanf(fp, []byte("%d\x00"), (*[100000000]int)(unsafe.Pointer(&m[i]))[:])
		if sfrv != 1 && Cmethod[0] == 3 {
			sferr([]byte("mode number in condensation data\x00"))
			noarch.Sprintf(errMsg, []byte("condensed mode %d = %d\x00"), i, m[i])
			errorMsg(errMsg)
		}
		if (m[i] < 0 || m[i] > nM) && Cmethod[0] == 3 {
			noarch.Sprintf(errMsg, []byte("\n  error in matrix condensation data: \n  m[%d] = %d \n The condensed mode number must be between   1 and %d (modes).\n\x00"), i, m[i], nM)
			errorMsg(errMsg)
			os.Exit(92)
		}
	}
	free_imatrix(cm, 1, int32(nC[0]), 1, 7)
}

// write_input_data - transpiled function from  $GOPATH/src/github.com/Konstantin8105/History_frame3DD/src/frame3dd_io.c:1773
/*
 * WRITE_INPUT_DATA  -  save input data					07nov02
 */ //
/* modern time variable type	*/ //
//frame3dd.sf.net/\n");
//www.fsf.org/copyleft/gpl.html\n");
/* start load case loop */ //
/* end load case loop	*/ //
//
func write_input_data(fp *noarch.File, title []byte, nN int, nE int, nL int, nD []int, nR int, nF []int, nU []int, nW []int, nP []int, nT []int, xyz []vec3, r []float32, J1 []int, J2 []int, Ax []float32, Asy []float32, Asz []float32, Jx []float32, Iy []float32, Iz []float32, E []float32, G []float32, p []float32, d []float32, gX []float32, gY []float32, gZ []float32, Ft [][]float64, Fm [][]float64, Dp [][]float32, R []int, U [][][]float32, W [][][]float32, P [][][]float32, T [][][]float32, shear int, anlyz int, geom int) {
	var i int
	var j int
	var n int
	var lc int
	var now noarch.TimeT
	_ = noarch.Time((*[100000000]noarch.TimeT)(unsafe.Pointer(&now))[:])
	for i = 1; i <= 80; i++ {
		noarch.Fprintf(fp, []byte("_\x00"))
	}
	noarch.Fprintf(fp, []byte("\nFrame3DD version: %s \x00"), []byte("20140514+\x00"))
	noarch.Fprintf(fp, []byte("              http://frame3dd.sf.net/\n\x00"))
	noarch.Fprintf(fp, []byte("GPL Copyright (C) 1992-2015, Henri P. Gavin \n\x00"))
	noarch.Fprintf(fp, []byte("Frame3DD is distributed in the hope that it will be useful\x00"))
	noarch.Fprintf(fp, []byte(" but with no warranty.\n\x00"))
	noarch.Fprintf(fp, []byte("For details see the GNU Public Licence:\x00"))
	noarch.Fprintf(fp, []byte(" http://www.fsf.org/copyleft/gpl.html\n\x00"))
	for i = 1; i <= 80; i++ {
		noarch.Fprintf(fp, []byte("_\x00"))
	}
	noarch.Fprintf(fp, []byte("\n\n\x00"))
	noarch.Fprintf(fp, []byte("%s\n\x00"), title)
	noarch.Fprintf(fp, []byte("%s\x00"), noarch.Ctime((*[100000000]noarch.TimeT)(unsafe.Pointer(&now))[:]))
	for i = 1; i <= 80; i++ {
		noarch.Fprintf(fp, []byte("_\x00"))
	}
	noarch.Fprintf(fp, []byte("\n\x00"))
	noarch.Fprintf(fp, []byte("In 2D problems the Y-axis is vertical.  \x00"))
	noarch.Fprintf(fp, []byte("In 3D problems the Z-axis is vertical.\n\x00"))
	for i = 1; i <= 80; i++ {
		noarch.Fprintf(fp, []byte("_\x00"))
	}
	noarch.Fprintf(fp, []byte("\n\x00"))
	noarch.Fprintf(fp, []byte("%5d NODES          \x00"), nN)
	noarch.Fprintf(fp, []byte("%5d FIXED NODES    \x00"), nR)
	noarch.Fprintf(fp, []byte("%5d FRAME ELEMENTS \x00"), nE)
	noarch.Fprintf(fp, []byte("%3d LOAD CASES   \n\x00"), nL)
	for i = 1; i <= 80; i++ {
		noarch.Fprintf(fp, []byte("_\x00"))
	}
	noarch.Fprintf(fp, []byte("\n\x00"))
	noarch.Fprintf(fp, []byte("N O D E   D A T A       \x00"))
	noarch.Fprintf(fp, []byte("                                    R E S T R A I N T S\n\x00"))
	noarch.Fprintf(fp, []byte("  Node       X              Y              Z\x00"))
	noarch.Fprintf(fp, []byte("         radius  Fx Fy Fz Mx My Mz\n\x00"))
	for i = 1; i <= nN; i++ {
		j = 6 * (i - 1)
		noarch.Fprintf(fp, []byte("%5d %14.6f %14.6f %14.6f %8.3f  %2d %2d %2d %2d %2d %2d\n\x00"), i, xyz[i].x, xyz[i].y, xyz[i].z, float64(r[i]), R[j+1], R[j+2], R[j+3], R[j+4], R[j+5], R[j+6])
	}
	noarch.Fprintf(fp, []byte("F R A M E   E L E M E N T   D A T A\t\t\t\t\t(local)\n\x00"))
	noarch.Fprintf(fp, []byte("  Elmnt  J1    J2     Ax   Asy   Asz    \x00"))
	noarch.Fprintf(fp, []byte("Jxx     Iyy     Izz       E       G roll  density\n\x00"))
	for i = 1; i <= nE; i++ {
		noarch.Fprintf(fp, []byte("%5d %5d %5d %6.1f %5.1f %5.1f\x00"), i, J1[i], J2[i], float64(Ax[i]), float64(Asy[i]), float64(Asz[i]))
		noarch.Fprintf(fp, []byte(" %6.1f %7.1f %7.1f %8.1f %7.1f %3.0f %8.2e\n\x00"), float64(Jx[i]), float64(Iy[i]), float64(Iz[i]), float64(E[i]), float64(G[i]), float64(p[i])*180/3.141592653589793, float64(d[i]))
	}
	if shear != 0 {
		noarch.Fprintf(fp, []byte("  Include shear deformations.\n\x00"))
	} else {
		noarch.Fprintf(fp, []byte("  Neglect shear deformations.\n\x00"))
	}
	if geom != 0 {
		noarch.Fprintf(fp, []byte("  Include geometric stiffness.\n\x00"))
	} else {
		noarch.Fprintf(fp, []byte("  Neglect geometric stiffness.\n\x00"))
	}
	for lc = 1; lc <= nL; lc++ {
		noarch.Fprintf(fp, []byte("\nL O A D   C A S E   %d   O F   %d  ... \n\n\x00"), lc, nL)
		noarch.Fprintf(fp, []byte("   Gravity X = \x00"))
		if gX[lc] == 0 {
			noarch.Fprintf(fp, []byte(" 0.0 \x00"))
		} else {
			noarch.Fprintf(fp, []byte(" %.3f \x00"), float64(gX[lc]))
		}
		noarch.Fprintf(fp, []byte("   Gravity Y = \x00"))
		if gY[lc] == 0 {
			noarch.Fprintf(fp, []byte(" 0.0 \x00"))
		} else {
			noarch.Fprintf(fp, []byte(" %.3f \x00"), float64(gY[lc]))
		}
		noarch.Fprintf(fp, []byte("   Gravity Z = \x00"))
		if gZ[lc] == 0 {
			noarch.Fprintf(fp, []byte(" 0.0 \x00"))
		} else {
			noarch.Fprintf(fp, []byte(" %.3f \x00"), float64(gZ[lc]))
		}
		noarch.Fprintf(fp, []byte("\n\x00"))
		noarch.Fprintf(fp, []byte(" %3d concentrated loads\n\x00"), nF[lc])
		noarch.Fprintf(fp, []byte(" %3d uniformly distributed loads\n\x00"), nU[lc])
		noarch.Fprintf(fp, []byte(" %3d trapezoidally distributed loads\n\x00"), nW[lc])
		noarch.Fprintf(fp, []byte(" %3d concentrated point loads\n\x00"), nP[lc])
		noarch.Fprintf(fp, []byte(" %3d temperature loads\n\x00"), nT[lc])
		noarch.Fprintf(fp, []byte(" %3d prescribed displacements\n\x00"), nD[lc])
		if nF[lc] > 0 || nU[lc] > 0 || nW[lc] > 0 || nP[lc] > 0 || nT[lc] > 0 {
			noarch.Fprintf(fp, []byte(" N O D A L   L O A D S\x00"))
			noarch.Fprintf(fp, []byte("  +  E Q U I V A L E N T   N O D A L   L O A D S  (global)\n\x00"))
			noarch.Fprintf(fp, []byte("  Node        Fx          Fy          Fz\x00"))
			noarch.Fprintf(fp, []byte("          Mxx         Myy         Mzz\n\x00"))
			for j = 1; j <= nN; j++ {
				i = 6 * (j - 1)
				if Fm[lc][i+1] != 0 || Fm[lc][i+2] != 0 || Fm[lc][i+3] != 0 || Fm[lc][i+4] != 0 || Fm[lc][i+5] != 0 || Fm[lc][i+6] != 0 {
					noarch.Fprintf(fp, []byte(" %5d\x00"), j)
					for i = 5; i >= 0; i-- {
						noarch.Fprintf(fp, []byte(" %11.3f\x00"), Fm[lc][6*j-i])
					}
					noarch.Fprintf(fp, []byte("\n\x00"))
				}
			}
		}
		if nU[lc] > 0 {
			noarch.Fprintf(fp, []byte(" U N I F O R M   L O A D S\x00"))
			noarch.Fprintf(fp, []byte("\t\t\t\t\t\t(local)\n\x00"))
			noarch.Fprintf(fp, []byte("  Elmnt       Ux               Uy               Uz\n\x00"))
			for n = 1; n <= nU[lc]; n++ {
				noarch.Fprintf(fp, []byte(" %5d\x00"), int((U[lc][n][1])))
				for i = 2; i <= 4; i++ {
					noarch.Fprintf(fp, []byte(" %16.8f\x00"), float64(U[lc][n][i]))
				}
				noarch.Fprintf(fp, []byte("\n\x00"))
			}
		}
		if nW[lc] > 0 {
			noarch.Fprintf(fp, []byte(" T R A P E Z O I D A L   L O A D S\x00"))
			noarch.Fprintf(fp, []byte("\t\t\t\t\t(local)\n\x00"))
			noarch.Fprintf(fp, []byte("  Elmnt       x1               x2               W1               W2\n\x00"))
			for n = 1; n <= nW[lc]; n++ {
				noarch.Fprintf(fp, []byte(" %5d\x00"), int((W[lc][n][1])))
				for i = 2; i <= 5; i++ {
					noarch.Fprintf(fp, []byte(" %16.8f\x00"), float64(W[lc][n][i]))
				}
				noarch.Fprintf(fp, []byte("  (x)\n\x00"))
				noarch.Fprintf(fp, []byte(" %5d\x00"), int((W[lc][n][1])))
				for i = 6; i <= 9; i++ {
					noarch.Fprintf(fp, []byte(" %16.8f\x00"), float64(W[lc][n][i]))
				}
				noarch.Fprintf(fp, []byte("  (y)\n\x00"))
				noarch.Fprintf(fp, []byte(" %5d\x00"), int((W[lc][n][1])))
				for i = 10; i <= 13; i++ {
					noarch.Fprintf(fp, []byte(" %16.8f\x00"), float64(W[lc][n][i]))
				}
				noarch.Fprintf(fp, []byte("  (z)\n\x00"))
			}
		}
		if nP[lc] > 0 {
			noarch.Fprintf(fp, []byte(" C O N C E N T R A T E D   P O I N T   L O A D S\x00"))
			noarch.Fprintf(fp, []byte("\t\t\t\t(local)\n\x00"))
			noarch.Fprintf(fp, []byte("  Elmnt       Px          Py          Pz          x\n\x00"))
			for n = 1; n <= nP[lc]; n++ {
				noarch.Fprintf(fp, []byte(" %5d\x00"), int((P[lc][n][1])))
				for i = 2; i <= 5; i++ {
					noarch.Fprintf(fp, []byte(" %11.3f\x00"), float64(P[lc][n][i]))
				}
				noarch.Fprintf(fp, []byte("\n\x00"))
			}
		}
		if nT[lc] > 0 {
			noarch.Fprintf(fp, []byte(" T E M P E R A T U R E   C H A N G E S\x00"))
			noarch.Fprintf(fp, []byte("\t\t\t\t\t(local)\n\x00"))
			noarch.Fprintf(fp, []byte("  Elmnt     coef      hy        hz\x00"))
			noarch.Fprintf(fp, []byte("        Ty+       Ty-       Tz+       Tz-\n\x00"))
			for n = 1; n <= nT[lc]; n++ {
				noarch.Fprintf(fp, []byte(" %5d\x00"), int((T[lc][n][1])))
				noarch.Fprintf(fp, []byte(" %9.2e\x00"), float64(T[lc][n][2]))
				for i = 3; i <= 8; i++ {
					noarch.Fprintf(fp, []byte(" %9.3f\x00"), float64(T[lc][n][i]))
				}
				noarch.Fprintf(fp, []byte("\n\x00"))
			}
		}
		if nD[lc] > 0 {
			noarch.Fprintf(fp, []byte("\n P R E S C R I B E D   D I S P L A C E M E N T S\x00"))
			noarch.Fprintf(fp, []byte("                        (global)\n\x00"))
			noarch.Fprintf(fp, []byte("  Node        Dx          Dy          Dz\x00"))
			noarch.Fprintf(fp, []byte("          Dxx         Dyy         Dzz\n\x00"))
			for j = 1; j <= nN; j++ {
				i = 6 * (j - 1)
				if float64(Dp[lc][i+1]) != 0 || float64(Dp[lc][i+2]) != 0 || float64(Dp[lc][i+3]) != 0 || float64(Dp[lc][i+4]) != 0 || float64(Dp[lc][i+5]) != 0 || float64(Dp[lc][i+6]) != 0 {
					noarch.Fprintf(fp, []byte(" %5d\x00"), j)
					for i = 5; i >= 0; i-- {
						noarch.Fprintf(fp, []byte(" %11.3f\x00"), float64(Dp[lc][6*j-i]))
					}
					noarch.Fprintf(fp, []byte("\n\x00"))
				}
			}
		}
	}
	if anlyz != 0 {
		noarch.Fprintf(fp, []byte("\nE L A S T I C   S T I F F N E S S   A N A L Y S I S\x00"))
		noarch.Fprintf(fp, []byte("   via  L D L'  decomposition\n\n\x00"))
	} else {
		noarch.Fprintf(fp, []byte("D A T A   C H E C K   O N L Y\n\x00"))
	}
	noarch.Fflush(fp)
}

// write_static_results - transpiled function from  $GOPATH/src/github.com/Konstantin8105/History_frame3DD/src/frame3dd_io.c:1964
/*
 * WRITE_STATIC_RESULTS -  save node displacements and frame element end forces
 * 09 Sep 2008 , 2015-05-15
 */ //
/*	 return; */ //
//
func write_static_results(fp *noarch.File, nN int, nE int, nL int, lc int, DoF int, J1 []int, J2 []int, F []float64, D []float64, R []float64, r []int, Q [][]float64, err float64, ok int, axial_sign int) {
	var disp float64
	var i int
	var j int
	var n int
	if ok < 0 {
		noarch.Fprintf(fp, []byte("  * The Stiffness Matrix is not positive-definite *\n\x00"))
		noarch.Fprintf(fp, []byte("    Check that all six rigid-body translations are restrained\n\x00"))
		noarch.Fprintf(fp, []byte("    If geometric stiffness is included, reduce the loads.\n\x00"))
	}
	noarch.Fprintf(fp, []byte("\nL O A D   C A S E   %d   O F   %d  ... \n\n\x00"), lc, nL)
	noarch.Fprintf(fp, []byte("N O D E   D I S P L A C E M E N T S  \x00"))
	noarch.Fprintf(fp, []byte("\t\t\t\t\t(global)\n\x00"))
	noarch.Fprintf(fp, []byte("  Node    X-dsp       Y-dsp       Z-dsp\x00"))
	noarch.Fprintf(fp, []byte("       X-rot       Y-rot       Z-rot\n\x00"))
	for j = 1; j <= nN; j++ {
		disp = 0
		for i = 5; i >= 0; i-- {
			disp += math.Abs(D[6*j-i])
		}
		if disp > 0 {
			noarch.Fprintf(fp, []byte(" %5d\x00"), j)
			for i = 5; i >= 0; i-- {
				if math.Abs(D[6*j-i]) < 1e-08 {
					noarch.Fprintf(fp, []byte("    0.0     \x00"))
				} else {
					noarch.Fprintf(fp, []byte(" %11.6f\x00"), D[6*j-i])
				}
			}
			noarch.Fprintf(fp, []byte("\n\x00"))
		}
	}
	noarch.Fprintf(fp, []byte("F R A M E   E L E M E N T   E N D   F O R C E S\x00"))
	noarch.Fprintf(fp, []byte("\t\t\t\t(local)\n\x00"))
	noarch.Fprintf(fp, []byte("  Elmnt  Node       Nx          Vy         Vz\x00"))
	noarch.Fprintf(fp, []byte("        Txx        Myy        Mzz\n\x00"))
	for n = 1; n <= nE; n++ {
		noarch.Fprintf(fp, []byte(" %5d  %5d\x00"), n, J1[n])
		if math.Abs(Q[n][1]) < 0.0001 {
			noarch.Fprintf(fp, []byte("      0.0   \x00"))
		} else {
			noarch.Fprintf(fp, []byte(" %10.3f\x00"), Q[n][1])
		}
		if Q[n][1] >= 0.0001 && axial_sign != 0 {
			noarch.Fprintf(fp, []byte("c\x00"))
		}
		if Q[n][1] <= -0.0001 && axial_sign != 0 {
			noarch.Fprintf(fp, []byte("t\x00"))
		}
		if noarch.NotInt(axial_sign) != 0 {
			noarch.Fprintf(fp, []byte(" \x00"))
		}
		for i = 2; i <= 6; i++ {
			if math.Abs(Q[n][i]) < 0.0001 {
				noarch.Fprintf(fp, []byte("      0.0  \x00"))
			} else {
				noarch.Fprintf(fp, []byte(" %10.3f\x00"), Q[n][i])
			}
		}
		noarch.Fprintf(fp, []byte("\n\x00"))
		noarch.Fprintf(fp, []byte(" %5d  %5d\x00"), n, J2[n])
		if math.Abs(Q[n][7]) < 0.0001 {
			noarch.Fprintf(fp, []byte("      0.0   \x00"))
		} else {
			noarch.Fprintf(fp, []byte(" %10.3f\x00"), Q[n][7])
		}
		if Q[n][7] >= 0.0001 && axial_sign != 0 {
			noarch.Fprintf(fp, []byte("t\x00"))
		}
		if Q[n][7] <= -0.0001 && axial_sign != 0 {
			noarch.Fprintf(fp, []byte("c\x00"))
		}
		if noarch.NotInt(axial_sign) != 0 {
			noarch.Fprintf(fp, []byte(" \x00"))
		}
		for i = 8; i <= 12; i++ {
			if math.Abs(Q[n][i]) < 0.0001 {
				noarch.Fprintf(fp, []byte("      0.0  \x00"))
			} else {
				noarch.Fprintf(fp, []byte(" %10.3f\x00"), Q[n][i])
			}
		}
		noarch.Fprintf(fp, []byte("\n\x00"))
	}
	noarch.Fprintf(fp, []byte("R E A C T I O N S\t\t\t\t\t\t\t(global)\n\x00"))
	noarch.Fprintf(fp, []byte("  Node        Fx          Fy          Fz\x00"))
	noarch.Fprintf(fp, []byte("         Mxx         Myy         Mzz\n\x00"))
	for j = 1; j <= nN; j++ {
		i = 6 * (j - 1)
		if r[i+1] != 0 || r[i+2] != 0 || r[i+3] != 0 || r[i+4] != 0 || r[i+5] != 0 || r[i+6] != 0 {
			noarch.Fprintf(fp, []byte(" %5d\x00"), j)
			for i = 5; i >= 0; i-- {
				if r[6*j-i] != 0 {
					noarch.Fprintf(fp, []byte(" %11.3f\x00"), R[6*j-i])
				} else {
					noarch.Fprintf(fp, []byte("       0.0  \x00"))
				}
			}
			noarch.Fprintf(fp, []byte("\n\x00"))
		}
	}
	noarch.Fprintf(fp, []byte("R M S    R E L A T I V E    E Q U I L I B R I U M    E R R O R: %9.3e\n\x00"), err)
	noarch.Fflush(fp)
}

// CSV_filename - transpiled function from  $GOPATH/src/github.com/Konstantin8105/History_frame3DD/src/frame3dd_io.c:2058
/*
 * CSV_filename - return the file name for the .CSV file and
 * whether the file should be written or appended (wa)
 * 1 Nov 2015
 */ //
//	fprintf(stderr," 1. CSV_file = %s  wa = %s \n", CSV_file, wa );
//
func CSV_filename(CSV_file []byte, wa []byte, OUT_file []byte, lc int) {
	var i int
	var j int
	i = 0
	j = 0
	for i < 128 {
		CSV_file[j] = OUT_file[i]
		if int(CSV_file[j]) == int('+') || int(CSV_file[j]) == int('-') || int(CSV_file[j]) == int('*') || int(CSV_file[j]) == int('^') || int(CSV_file[j]) == int('.') || int(CSV_file[j]) == int('\x00') {
			CSV_file[j] = '_'
			break
		}
		i++
		j++
	}
	CSV_file[func() int {
		j++
		return j
	}()] = '\x00'
	noarch.Strcat(CSV_file, []byte("out.CSV\x00"))
	wa[0] = 'a'
	if lc == 1 {
		wa[0] = 'w'
	}
	wa[1] = '\x00'
}

// write_static_csv - transpiled function from  $GOPATH/src/github.com/Konstantin8105/History_frame3DD/src/frame3dd_io.c:2093
/*
 * WRITE_STATIC_CSV -  save node displacements and frame element end forces
 * 31 Dec 2008
 */ //
/* modern time variable type	*/ //
//frame3dd.sf.net/\"\n");
//www.fsf.org/copyleft/gpl.html\"\n");
/*	 return; */ //
//
func write_static_csv(OUT_file []byte, title []byte, nN int, nE int, nL int, lc int, DoF int, J1 []int, J2 []int, F []float64, D []float64, R []float64, r []int, Q [][]float64, err float64, ok int) {
	var fpcsv *noarch.File
	var i int
	var j int
	var n int
	var wa []byte = make([]byte, 4)
	var CSV_file []byte = make([]byte, 128)
	var now noarch.TimeT
	var errMsg []byte = make([]byte, 512)
	_ = noarch.Time((*[100000000]noarch.TimeT)(unsafe.Pointer(&now))[:])
	CSV_filename(CSV_file, wa, OUT_file, lc)
	if (func() *noarch.File {
		fpcsv = noarch.Fopen(CSV_file, wa)
		return fpcsv
	}()) == nil {
		noarch.Sprintf(errMsg, []byte("\n  error: cannot open CSV output data file: %s \n\x00"), CSV_file)
		errorMsg(errMsg)
		os.Exit(17)
	}
	if lc == 1 {
		noarch.Fprintf(fpcsv, []byte("\" Frame3DD version: %s \x00"), []byte("20140514+\x00"))
		noarch.Fprintf(fpcsv, []byte("              http://frame3dd.sf.net/\"\n\x00"))
		noarch.Fprintf(fpcsv, []byte("\"GPL Copyright (C) 1992-2015, Henri P. Gavin \"\n\x00"))
		noarch.Fprintf(fpcsv, []byte("\"Frame3DD is distributed in the hope that it will be useful\x00"))
		noarch.Fprintf(fpcsv, []byte(" but with no warranty.\"\n\x00"))
		noarch.Fprintf(fpcsv, []byte("\"For details see the GNU Public Licence:\x00"))
		noarch.Fprintf(fpcsv, []byte(" http://www.fsf.org/copyleft/gpl.html\"\n\x00"))
		noarch.Fprintf(fpcsv, []byte("\" %s \"\n\x00"), title)
		noarch.Fprintf(fpcsv, []byte("\" %s \"\n\x00"), noarch.Ctime((*[100000000]noarch.TimeT)(unsafe.Pointer(&now))[:]))
		noarch.Fprintf(fpcsv, []byte("\" .CSV formatted results of Frame3DD analysis \"\n\x00"))
		noarch.Fprintf(fpcsv, []byte("\n , Load Case , Displacements , End Forces , Reactions , Internal Forces \n\x00"))
		for i = 1; i <= nL; i++ {
			noarch.Fprintf(fpcsv, []byte(" First Row , %d , %d , %d , %d  , %d  \n\x00"), i, 15+(i-1)*(nN*2+nE*4+13)+2*nL, 17+(i-1)*(nN*2+nE*4+13)+2*nL+nN, 19+(i-1)*(nN*2+nE*4+13)+2*nL+nN+2*nE, 23+(i-1)*(nN*2+nE*4+13)+2*nL+2*nN+2*nE)
			noarch.Fprintf(fpcsv, []byte(" Last Row , %d , %d , %d , %d , %d \n\x00"), i, 15+(i-1)*(nN*2+nE*4+13)+2*nL+nN-1, 17+(i-1)*(nN*2+nE*4+13)+2*nL+nN+2*nE-1, 19+(i-1)*(nN*2+nE*4+13)+2*nL+2*nN+2*nE-1, 23+(i-1)*(nN*2+nE*4+13)+2*nL+2*nN+4*nE-1)
		}
	}
	if ok < 0 {
		noarch.Fprintf(fpcsv, []byte("\"  * The Stiffness Matrix is not positive-definite * \"\n\x00"))
		noarch.Fprintf(fpcsv, []byte("\" Check that all six rigid-body translations are restrained\"\n\x00"))
		noarch.Fprintf(fpcsv, []byte("\" If geometric stiffness is included, reduce the loads.\"\n\x00"))
	}
	noarch.Fprintf(fpcsv, []byte("\n\"L O A D   C A S E   %d   O F   %d  ... \"\n\n\x00"), lc, nL)
	noarch.Fprintf(fpcsv, []byte("\"N O D E   D I S P L A C E M E N T S\x00"))
	noarch.Fprintf(fpcsv, []byte("    (global)\"\n\x00"))
	noarch.Fprintf(fpcsv, []byte("Node  ,  X-dsp   ,   Y-dsp  ,    Z-dsp\x00"))
	noarch.Fprintf(fpcsv, []byte(" ,     X-rot  ,    Y-rot   ,   Z-rot\n\x00"))
	for j = 1; j <= nN; j++ {
		noarch.Fprintf(fpcsv, []byte(" %5d,\x00"), j)
		for i = 5; i >= 0; i-- {
			if math.Abs(D[6*j-i]) < 1e-08 {
				noarch.Fprintf(fpcsv, []byte("    0.0,    \x00"))
			} else {
				noarch.Fprintf(fpcsv, []byte(" %12.5e,\x00"), D[6*j-i])
			}
		}
		noarch.Fprintf(fpcsv, []byte("\n\x00"))
	}
	noarch.Fprintf(fpcsv, []byte("\"F R A M E   E L E M E N T   E N D   F O R C E S\x00"))
	noarch.Fprintf(fpcsv, []byte("  (local)\"\n\x00"))
	noarch.Fprintf(fpcsv, []byte("Elmnt , Node  ,    Nx     ,    Vy   ,     Vz\x00"))
	noarch.Fprintf(fpcsv, []byte("   ,     Txx   ,    Myy  ,     Mzz\n\x00"))
	for n = 1; n <= nE; n++ {
		noarch.Fprintf(fpcsv, []byte(" %5d, %5d,\x00"), n, J1[n])
		if math.Abs(Q[n][1]) < 0.0001 {
			noarch.Fprintf(fpcsv, []byte("      0.0,  \x00"))
		} else {
			noarch.Fprintf(fpcsv, []byte(" %12.5e,\x00"), Q[n][1])
		}
		for i = 2; i <= 6; i++ {
			if math.Abs(Q[n][i]) < 0.0001 {
				noarch.Fprintf(fpcsv, []byte("      0.0, \x00"))
			} else {
				noarch.Fprintf(fpcsv, []byte(" %12.5e,\x00"), Q[n][i])
			}
		}
		noarch.Fprintf(fpcsv, []byte("\n\x00"))
		noarch.Fprintf(fpcsv, []byte(" %5d, %5d,\x00"), n, J2[n])
		if math.Abs(Q[n][7]) < 0.0001 {
			noarch.Fprintf(fpcsv, []byte("      0.0,  \x00"))
		} else {
			noarch.Fprintf(fpcsv, []byte(" %12.5e,\x00"), Q[n][7])
		}
		for i = 8; i <= 12; i++ {
			if math.Abs(Q[n][i]) < 0.0001 {
				noarch.Fprintf(fpcsv, []byte("      0.0, \x00"))
			} else {
				noarch.Fprintf(fpcsv, []byte(" %12.5e,\x00"), Q[n][i])
			}
		}
		noarch.Fprintf(fpcsv, []byte("\n\x00"))
	}
	noarch.Fprintf(fpcsv, []byte("\"R E A C T I O N S  (global)\"\n\x00"))
	noarch.Fprintf(fpcsv, []byte(" Node   ,    Fx      ,   Fy   ,      Fz\x00"))
	noarch.Fprintf(fpcsv, []byte("   ,     Mxx    ,    Myy    ,    Mzz\n\x00"))
	for j = 1; j <= nN; j++ {
		i = 6 * (j - 1)
		noarch.Fprintf(fpcsv, []byte(" %5d,\x00"), j)
		for i = 5; i >= 0; i-- {
			if r[6*j-i] != 0 {
				noarch.Fprintf(fpcsv, []byte(" %12.5e,\x00"), R[6*j-i])
			} else {
				noarch.Fprintf(fpcsv, []byte("       0.0, \x00"))
			}
		}
		noarch.Fprintf(fpcsv, []byte("\n\x00"))
	}
	noarch.Fprintf(fpcsv, []byte("\"R M S    R E L A T I V E    E Q U I L I B R I U M    E R R O R:\", %9.3e\n\x00"), err)
	noarch.Fclose(fpcsv)
}

// write_static_mfile - transpiled function from  $GOPATH/src/github.com/Konstantin8105/History_frame3DD/src/frame3dd_io.c:2248
/*
 * WRITE_VALUE - write a value in %f or %e notation depending on numerical values
 * and the number of available significant figures
 * 12 Dec 2009
 */ //
/*
void write_value (
		FILE *fp,
		int sig_figs,
		float threshold,
		char *spaces,
		double x
){
	int nZspaces;

	nZspaces = (int) strlen(*spaces);

	if ( fabs(x) < threshold ) fprintf ( fp, "0.0 \n");

}
*/ //
/*
 * WRITE_STATIC_MFILE -
 * save node displacements and frame element end forces in an m-file
 * this function interacts with frame_3dd.m, an m-file interface to frame3dd
 * 09 Sep 2008
 */ //
/* modern time variable type	*/ //
//frame3dd.sf.net/\n");
//www.fsf.org/copyleft/gpl.html\n");
/*	 return; */ //
//
func write_static_mfile(OUT_file []byte, title []byte, nN int, nE int, nL int, lc int, DoF int, J1 []int, J2 []int, F []float64, D []float64, R []float64, r []int, Q [][]float64, err float64, ok int) {
	var fpm *noarch.File
	var i int
	var j int
	var n int
	var wa []byte
	var M_file []byte = make([]byte, 128)
	var now noarch.TimeT
	var errMsg []byte = make([]byte, 512)
	_ = noarch.Time((*[100000000]noarch.TimeT)(unsafe.Pointer(&now))[:])
	i = 0
	j = 0
	for i < 128 {
		M_file[j] = OUT_file[i]
		if int(M_file[j]) == int('+') || int(M_file[j]) == int('-') || int(M_file[j]) == int('*') || int(M_file[j]) == int('^') || int(M_file[j]) == int('.') || int(M_file[j]) == int('\x00') {
			M_file[j] = '_'
			break
		}
		i++
		j++
	}
	M_file[func() int {
		j++
		return j
	}()] = '\x00'
	noarch.Strcat(M_file, []byte("out.m\x00"))
	wa = []byte("a\x00")
	if lc == 1 {
		wa = []byte("w\x00")
	}
	if (func() *noarch.File {
		fpm = noarch.Fopen(M_file, wa)
		return fpm
	}()) == nil {
		noarch.Sprintf(errMsg, []byte("\n  error: cannot open Matlab output data file: %s \n\x00"), M_file)
		errorMsg(errMsg)
		os.Exit(18)
	}
	if lc == 1 {
		noarch.Fprintf(fpm, []byte("%% Frame3DD version: %s \x00"), []byte("20140514+\x00"))
		noarch.Fprintf(fpm, []byte("              http://frame3dd.sf.net/\n\x00"))
		noarch.Fprintf(fpm, []byte("%%GPL Copyright (C) 1992-2015, Henri P. Gavin \n\x00"))
		noarch.Fprintf(fpm, []byte("%%Frame3DD is distributed in the hope that it will be useful\x00"))
		noarch.Fprintf(fpm, []byte(" but with no warranty.\n\x00"))
		noarch.Fprintf(fpm, []byte("%%For details see the GNU Public Licence:\x00"))
		noarch.Fprintf(fpm, []byte(" http://www.fsf.org/copyleft/gpl.html\n\x00"))
		noarch.Fprintf(fpm, []byte("%% %s\n\x00"), title)
		noarch.Fprintf(fpm, []byte("%% %s\x00"), noarch.Ctime((*[100000000]noarch.TimeT)(unsafe.Pointer(&now))[:]))
		noarch.Fprintf(fpm, []byte("%% m-file formatted results of frame3dd analysis\n\x00"))
		noarch.Fprintf(fpm, []byte("%% to be read by frame_3dd.m\n\x00"))
	}
	if ok < 0 {
		noarch.Fprintf(fpm, []byte("%%  The Stiffness Matrix is not positive-definite *\n\x00"))
		noarch.Fprintf(fpm, []byte("%%  Check that all six rigid-body translations are restrained\n\x00"))
		noarch.Fprintf(fpm, []byte("%%  If geometric stiffness is included, reduce the loads.\n\x00"))
	}
	noarch.Fprintf(fpm, []byte("\n%% L O A D   C A S E   %d   O F   %d  ... \n\n\x00"), lc, nL)
	noarch.Fprintf(fpm, []byte("%% N O D E   D I S P L A C E M E N T S  \x00"))
	noarch.Fprintf(fpm, []byte("\t\t(global)\n\x00"))
	noarch.Fprintf(fpm, []byte("%%\tX-dsp\t\tY-dsp\t\tZ-dsp\t\tX-rot\t\tY-rot\t\tZ-rot\n\x00"))
	noarch.Fprintf(fpm, []byte("D%d=[\x00"), lc)
	for j = 1; j <= nN; j++ {
		for i = 5; i >= 0; i-- {
			if math.Abs(D[6*j-i]) < 1e-08 {
				noarch.Fprintf(fpm, []byte("\t0.0\t\x00"))
			} else {
				noarch.Fprintf(fpm, []byte("\t%13.6e\x00"), D[6*j-i])
			}
		}
		if j < nN {
			noarch.Fprintf(fpm, []byte(" ; \n\x00"))
		} else {
			noarch.Fprintf(fpm, []byte(" ]'; \n\n\x00"))
		}
	}
	noarch.Fprintf(fpm, []byte("%% F R A M E   E L E M E N T   E N D   F O R C E S\x00"))
	noarch.Fprintf(fpm, []byte("\t\t(local)\n\x00"))
	noarch.Fprintf(fpm, []byte("%%\tNx_1\t\tVy_1\t\tVz_1\t\tTxx_1\t\tMyy_1\t\tMzz_1\t\x00"))
	noarch.Fprintf(fpm, []byte("  \tNx_2\t\tVy_2\t\tVz_2\t\tTxx_2\t\tMyy_2\t\tMzz_2\n\x00"))
	noarch.Fprintf(fpm, []byte("F%d=[\x00"), lc)
	for n = 1; n <= nE; n++ {
		if math.Abs(Q[n][1]) < 0.0001 {
			noarch.Fprintf(fpm, []byte("\t0.0\t\x00"))
		} else {
			noarch.Fprintf(fpm, []byte("\t%13.6e\x00"), Q[n][1])
		}
		for i = 2; i <= 6; i++ {
			if math.Abs(Q[n][i]) < 0.0001 {
				noarch.Fprintf(fpm, []byte("\t0.0\t\x00"))
			} else {
				noarch.Fprintf(fpm, []byte("\t%13.6e\x00"), Q[n][i])
			}
		}
		if math.Abs(Q[n][7]) < 0.0001 {
			noarch.Fprintf(fpm, []byte("\t0.0\t\x00"))
		} else {
			noarch.Fprintf(fpm, []byte("\t%13.6e\x00"), Q[n][7])
		}
		for i = 8; i <= 12; i++ {
			if math.Abs(Q[n][i]) < 0.0001 {
				noarch.Fprintf(fpm, []byte("\t0.0\t\x00"))
			} else {
				noarch.Fprintf(fpm, []byte("\t%13.6e\x00"), Q[n][i])
			}
		}
		if n < nE {
			noarch.Fprintf(fpm, []byte(" ; \n\x00"))
		} else {
			noarch.Fprintf(fpm, []byte(" ]'; \n\n\x00"))
		}
	}
	noarch.Fprintf(fpm, []byte("%% R E A C T I O N S\t\t\t\t(global)\n\x00"))
	noarch.Fprintf(fpm, []byte("%%\tFx\t\tFy\t\tFz\t\tMxx\t\tMyy\t\tMzz\n\x00"))
	noarch.Fprintf(fpm, []byte("R%d=[\x00"), lc)
	for j = 1; j <= nN; j++ {
		i = 6 * (j - 1)
		for i = 5; i >= 0; i-- {
			if noarch.NotInt(r[6*j-i]) != 0 || math.Abs(R[6*j-i]) < 0.0001 {
				noarch.Fprintf(fpm, []byte("\t0.0\t\x00"))
			} else {
				noarch.Fprintf(fpm, []byte("\t%13.6e\x00"), R[6*j-i])
			}
		}
		if j < nN {
			noarch.Fprintf(fpm, []byte(" ; \n\x00"))
		} else {
			noarch.Fprintf(fpm, []byte(" ]'; \n\n\x00"))
		}
	}
	noarch.Fprintf(fpm, []byte("%% R M S    R E L A T I V E    E Q U I L I B R I U M    E R R O R: %9.3e\n\x00"), err)
	noarch.Fprintf(fpm, []byte("\n\n  load Ks \n\n\x00"))
	noarch.Fclose(fpm)
}

// peak_internal_forces - transpiled function from  $GOPATH/src/github.com/Konstantin8105/History_frame3DD/src/frame3dd_io.c:2387
/*
 * PEAK_INTERNAL_FORCES
 * calculate frame element internal forces, Nx, Vy, Vz, Tx, My, Mz
 * calculate frame element local displacements, Rx, Dx, Dy, Dz
 * return the peak values of the internal forces, moments, slopes, and displacements
 * 18jun13
 */ //
// load case number
// total number of load cases
// node locations
// x-axis increment along frame element
// vectors of peak forces, moments, displacements and slopes
// for each frame element, for load case "lc"
/* coord transformation */ //
/*, u7, u8, u9, u10, u11, u12*/ //
/* displ. */ //
/* trapz load data, local x dir */ //
/* trapz load data, local y dir */ //
/* trapz load data, local z dir */ //
// distributed loads in local coords at x[i]
// distributed loads in local coords at x[i-1]
// gravity loads in local x, y, z coord's
// distributed torque about local x coord
/* location of internal point loads	*/ //
/* distance along frame element		*/ //
// underscored "_" variables correspond to x=(i-1)*dx;
// non-underscored variables correspond to x=i*dx;
/* axial force within frame el.		*/ //
/* shear forces within frame el.	*/ //
/* torsional moment within frame el.	*/ //
/* bending moments within frame el.	*/ //
/* transverse slopes of frame el.	*/ //
/* frame el. displ. in local x,y,z, dir's */ //
/* twist rotation about the local x-axis */ //
// frame element number
// number of sections alont x axis
// counters for U, W, and P loads
// counter along x axis from node N1 to node N2
/*,i2*/ //
// starting and stopping node numbers
// skip calculation of internal forces and displ
// initialize peak values to zero
// loop over all frame elements
// node 1 and node 2 of elmnt m
// x-axis increment, same for each element
// no need to allocate memory for interior force or displacement data
// find interior axial force, shear forces, torsion and bending moments
// distributed gravity load in local x, y, z coordinates
// add uniformly-distributed loads to gravity load
// load n on element m
// interior forces for frame element "m" at (x=0)
// positive Nx is tensile
// positive Vy in local y direction
// positive Vz in local z direction
// positive Tx r.h.r. about local x axis
// positive My -> positive x-z curvature
// positive Mz -> positive x-y curvature
//	i2 = 6*(n2-1);
/* compute end deflections in local coordinates */ //
/* u7  = t1*D[i2+1] + t2*D[i2+2] + t3*D[i2+3]; */ //
/* u8  = t4*D[i2+1] + t5*D[i2+2] + t6*D[i2+3]; */ //
/* u9  = t7*D[i2+1] + t8*D[i2+2] + t9*D[i2+3]; */ //
/*  */ //
/* u10 = t1*D[i2+4] + t2*D[i2+5] + t3*D[i2+6]; */ //
/* u11 = t4*D[i2+4] + t5*D[i2+5] + t6*D[i2+6]; */ //
/* u12 = t7*D[i2+4] + t8*D[i2+5] + t9*D[i2+6]; */ //
// rotations and displacements for frame element "m" at (x=0)
// displacement in  local x dir  at node N1
// displacement in  local y dir  at node N1
// displacement in  local z dir  at node N1
// rotationin about local x axis at node N1
// slope in  local y  direction  at node N1
// slope in  local z  direction  at node N1
// accumulate interior span loads, forces, moments, slopes, and displacements
// all in a single loop
// location from node N1 along the x-axis
// start with gravitational plus uniform loads
// add trapezoidally-distributed loads
// load n on element m
// trapezoidal integration of distributed loads
// for axial forces, shear forces and torques
// update distributed loads at x = (i-1)*dx
// add interior point loads
// load n on element m
// trapezoidal integration of shear force for bending momemnt
// displacement along frame element "m"
// torsional rotation along frame element "m"
// transverse slope along frame element "m"
// displacement along frame element "m"
// update forces, moments, and slopes at x = (i-1)*dx
// update the peak forces, moments, slopes and displacements
// and their locations along the frame element
// end of loop along element "m"
// at the end of this loop,
// the variables Nx; Vy; Vz; Tx; My; Mz; Dx; Dy; Dz; Rx; Sy; Sz;
// contain the forces, moments, displacements, and slopes
// at node N2 of element "m"
// comparing the internal forces and displacements at node N2
// to the values conmputed via trapezoidal rule could give an estimate
// of the accuracy of the trapezoidal rule, (how small "dx" needs to be)
// linear correction for bias in trapezoidal integration
// is not implemented, the peak values are affected by accumulation
// of round-off error in trapezoidal rule integration.
// round-off errors are larger in the peak displacements than in the peak forces
// end of loop over all frame elements
// DEBUG --- write output to terminal
//
func peak_internal_forces(lc int, nL int, xyz []vec3, Q [][]float64, nN int, nE int, L []float64, N1 []int, N2 []int, Ax []float32, Asy []float32, Asz []float32, Jx []float32, Iy []float32, Iz []float32, E []float32, G []float32, p []float32, d []float32, gX float32, gY float32, gZ float32, nU int, U [][]float32, nW int, W [][]float32, nP int, P [][]float32, D []float64, shear int, dx float32, pkNx [][]float64, pkVy [][]float64, pkVz [][]float64, pkTx [][]float64, pkMy [][]float64, pkMz [][]float64, pkDx [][]float64, pkDy [][]float64, pkDz [][]float64, pkRx [][]float64, pkSy [][]float64, pkSz [][]float64) {
	var t1 float64
	var t2 float64
	var t3 float64
	var t4 float64
	var t5 float64
	var t6 float64
	var t7 float64
	var t8 float64
	var t9 float64
	var u1 float64
	var u2 float64
	var u3 float64
	var u4 float64
	var u5 float64
	var u6 float64
	var xx1 float64
	var xx2 float64
	var wx1 float64
	var wx2 float64
	var xy1 float64
	var xy2 float64
	var wy1 float64
	var wy2 float64
	var xz1 float64
	var xz2 float64
	var wz1 float64
	var wz2 float64
	var wx float64
	var wy float64
	var wz float64
	var wx_ float64
	var wy_ float64
	var wz_ float64
	var wxg float64
	var wyg float64
	var wzg float64
	var tx float64
	var tx_ float64
	var xp float64
	var x float64
	var Nx_ float64
	var Nx float64
	var Vy_ float64
	var Vy float64
	var Vz_ float64
	var Vz float64
	var Tx_ float64
	var Tx float64
	var My_ float64
	var My float64
	var Mz_ float64
	var Mz float64
	var Sy_ float64
	var Sy float64
	var Sz_ float64
	var Sz float64
	var Dx float64
	var Dy float64
	var Dz float64
	var Rx float64
	var n int
	var m int
	var nx int = 1000
	var cU int
	var cW int
	var cP int
	var i int
	var n1 int
	var n2 int
	var i1 int
	if float64(dx) == -1 {
		return
	}
	for m = 1; m <= nE; m++ {
		pkVz[lc][m] = 0
		pkVy[lc][m] = pkVz[lc][m]
		pkNx[lc][m] = pkVy[lc][m]
		pkMz[lc][m] = 0
		pkMy[lc][m] = pkMz[lc][m]
		pkTx[lc][m] = pkMy[lc][m]
		pkDz[lc][m] = 0
		pkDy[lc][m] = pkDz[lc][m]
		pkDx[lc][m] = pkDy[lc][m]
		pkSz[lc][m] = 0
		pkSy[lc][m] = pkSz[lc][m]
		pkRx[lc][m] = pkSy[lc][m]
	}
	for m = 1; m <= nE; m++ {
		n1 = N1[m]
		n2 = N2[m]
		dx = float32(L[m] / float64(float32(nx)))
		coord_trans(xyz, L[m], n1, n2, (*[100000000]float64)(unsafe.Pointer(&t1))[:], (*[100000000]float64)(unsafe.Pointer(&t2))[:], (*[100000000]float64)(unsafe.Pointer(&t3))[:], (*[100000000]float64)(unsafe.Pointer(&t4))[:], (*[100000000]float64)(unsafe.Pointer(&t5))[:], (*[100000000]float64)(unsafe.Pointer(&t6))[:], (*[100000000]float64)(unsafe.Pointer(&t7))[:], (*[100000000]float64)(unsafe.Pointer(&t8))[:], (*[100000000]float64)(unsafe.Pointer(&t9))[:], p[m])
		wxg = float64(d[m]*Ax[m]) * (t1*float64(gX) + t2*float64(gY) + t3*float64(gZ))
		wyg = float64(d[m]*Ax[m]) * (t4*float64(gX) + t5*float64(gY) + t6*float64(gZ))
		wzg = float64(d[m]*Ax[m]) * (t7*float64(gX) + t8*float64(gY) + t9*float64(gZ))
		for n = 1; n <= nE && cU < nU; n++ {
			if int(U[n][1]) == m {
				wxg += float64(U[n][2])
				wyg += float64(U[n][3])
				wzg += float64(U[n][4])
				cU++
			}
		}
		Nx = -Q[m][1]
		Nx_ = Nx
		Vy = -Q[m][2]
		Vy_ = Vy
		Vz = -Q[m][3]
		Vz_ = Vz
		Tx = -Q[m][4]
		Tx_ = Tx
		My = Q[m][5]
		My_ = My
		Mz = -Q[m][6]
		Mz_ = Mz
		i1 = 6 * (n1 - 1)
		u1 = t1*D[i1+1] + t2*D[i1+2] + t3*D[i1+3]
		u2 = t4*D[i1+1] + t5*D[i1+2] + t6*D[i1+3]
		u3 = t7*D[i1+1] + t8*D[i1+2] + t9*D[i1+3]
		u4 = t1*D[i1+4] + t2*D[i1+5] + t3*D[i1+6]
		u5 = t4*D[i1+4] + t5*D[i1+5] + t6*D[i1+6]
		u6 = t7*D[i1+4] + t8*D[i1+5] + t9*D[i1+6]
		Dx = u1
		Dy = u2
		Dz = u3
		Rx = u4
		Sy = u6
		Sy_ = Sy
		Sz = -u5
		Sz_ = Sz
		for i = 1; i <= nx; i++ {
			x = float64(float32(i) * dx)
			wx = wxg
			wy = wyg
			wz = wzg
			if i == 1 {
				wx_ = wxg
				wy_ = wyg
				wz_ = wzg
				tx_ = tx
			}
			for n = 1; n <= 10*nE && cW < nW; n++ {
				if int(W[n][1]) == m {
					if i == nx {
						cW++
					}
					xx1 = float64(W[n][2])
					xx2 = float64(W[n][3])
					wx1 = float64(W[n][4])
					wx2 = float64(W[n][5])
					xy1 = float64(W[n][6])
					xy2 = float64(W[n][7])
					wy1 = float64(W[n][8])
					wy2 = float64(W[n][9])
					xz1 = float64(W[n][10])
					xz2 = float64(W[n][11])
					wz1 = float64(W[n][12])
					wz2 = float64(W[n][13])
					if x > xx1 && x <= xx2 {
						wx += wx1 + (wx2-wx1)*(x-xx1)/(xx2-xx1)
					}
					if x > xy1 && x <= xy2 {
						wy += wy1 + (wy2-wy1)*(x-xy1)/(xy2-xy1)
					}
					if x > xz1 && x <= xz2 {
						wz += wz1 + (wz2-wz1)*(x-xz1)/(xz2-xz1)
					}
				}
			}
			Nx = Nx - 0.5*(wx+wx_)*float64(dx)
			Vy = Vy - 0.5*(wy+wy_)*float64(dx)
			Vz = Vz - 0.5*(wz+wz_)*float64(dx)
			Tx = Tx - 0.5*(tx+tx_)*float64(dx)
			wx_ = wx
			wy_ = wy
			wz_ = wz
			tx_ = tx
			for n = 1; n <= 10*nE && cP < nP; n++ {
				if int(P[n][1]) == m {
					if i == nx {
						cP++
					}
					xp = float64(P[n][5])
					if x <= xp && xp < x+float64(dx) {
						Nx -= float64(P[n][2]) * 0.5 * (1 - (xp-x)/float64(dx))
						Vy -= float64(P[n][3]) * 0.5 * (1 - (xp-x)/float64(dx))
						Vz -= float64(P[n][4]) * 0.5 * (1 - (xp-x)/float64(dx))
					}
					if x-float64(dx) <= xp && xp < x {
						Nx -= float64(P[n][2]) * 0.5 * (1 - (x-float64(dx)-xp)/float64(dx))
						Vy -= float64(P[n][3]) * 0.5 * (1 - (x-float64(dx)-xp)/float64(dx))
						Vz -= float64(P[n][4]) * 0.5 * (1 - (x-float64(dx)-xp)/float64(dx))
					}
				}
			}
			My = My - 0.5*(Vz_+Vz)*float64(dx)
			Mz = Mz - 0.5*(Vy_+Vy)*float64(dx)
			Dx = Dx + 0.5*(Nx_+Nx)/float64(E[m]*Ax[m])*float64(dx)
			Rx = Rx + 0.5*(Tx_+Tx)/float64(G[m]*Jx[m])*float64(dx)
			Sy = Sy + 0.5*(Mz_+Mz)/float64(E[m]*Iz[m])*float64(dx)
			Sz = Sz + 0.5*(My_+My)/float64(E[m]*Iy[m])*float64(dx)
			if shear != 0 {
				Sy += Vy / float64(G[m]*Asy[m])
				Sz += Vz / float64(G[m]*Asz[m])
			}
			Dy = Dy + 0.5*(Sy_+Sy)*float64(dx)
			Dz = Dz + 0.5*(Sz_+Sz)*float64(dx)
			Nx_ = Nx
			Vy_ = Vy
			Vz_ = Vz
			Tx_ = Tx
			My_ = My
			Mz_ = Mz
			Sy_ = Sy
			Sz_ = Sz
			pkNx[lc][m] = func() float64 {
				if math.Abs(Nx) > pkNx[lc][m] {
					return math.Abs(Nx)
				}
				return pkNx[lc][m]
			}()
			pkVy[lc][m] = func() float64 {
				if math.Abs(Vy) > pkVy[lc][m] {
					return math.Abs(Vy)
				}
				return pkVy[lc][m]
			}()
			pkVz[lc][m] = func() float64 {
				if math.Abs(Vz) > pkVz[lc][m] {
					return math.Abs(Vz)
				}
				return pkVz[lc][m]
			}()
			pkTx[lc][m] = func() float64 {
				if math.Abs(Tx) > pkTx[lc][m] {
					return math.Abs(Tx)
				}
				return pkTx[lc][m]
			}()
			pkMy[lc][m] = func() float64 {
				if math.Abs(My) > pkMy[lc][m] {
					return math.Abs(My)
				}
				return pkMy[lc][m]
			}()
			pkMz[lc][m] = func() float64 {
				if math.Abs(Mz) > pkMz[lc][m] {
					return math.Abs(Mz)
				}
				return pkMz[lc][m]
			}()
			pkDx[lc][m] = func() float64 {
				if math.Abs(Dx) > pkDx[lc][m] {
					return math.Abs(Dx)
				}
				return pkDx[lc][m]
			}()
			pkDy[lc][m] = func() float64 {
				if math.Abs(Dy) > pkDy[lc][m] {
					return math.Abs(Dy)
				}
				return pkDy[lc][m]
			}()
			pkDz[lc][m] = func() float64 {
				if math.Abs(Dz) > pkDz[lc][m] {
					return math.Abs(Dz)
				}
				return pkDz[lc][m]
			}()
			pkRx[lc][m] = func() float64 {
				if math.Abs(Rx) > pkRx[lc][m] {
					return math.Abs(Rx)
				}
				return pkRx[lc][m]
			}()
			pkSy[lc][m] = func() float64 {
				if math.Abs(Sy) > pkSy[lc][m] {
					return math.Abs(Sy)
				}
				return pkSy[lc][m]
			}()
			pkSz[lc][m] = func() float64 {
				if math.Abs(Sz) > pkSz[lc][m] {
					return math.Abs(Sz)
				}
				return pkSz[lc][m]
			}()
		}
	}
	noarch.Fprintf(stderr, []byte("P E A K   F R A M E   E L E M E N T   I N T E R N A L   F O R C E S\x00"))
	noarch.Fprintf(stderr, []byte("\t(local)\n\x00"))
	noarch.Fprintf(stderr, []byte("  Elmnt       Nx          Vy         Vz\x00"))
	noarch.Fprintf(stderr, []byte("        Txx        Myy        Mzz\n\x00"))
	for m = 1; m <= nE; m++ {
		noarch.Fprintf(stderr, []byte(" %5d  %10.3f  %10.3f %10.3f %10.3f %10.3f %10.3f\n\x00"), m, pkNx[lc][m], pkVy[lc][m], pkVz[lc][m], pkTx[lc][m], pkMy[lc][m], pkMz[lc][m])
	}
	noarch.Fprintf(stderr, []byte("\n P E A K   I N T E R N A L   D I S P L A C E M E N T S\x00"))
	noarch.Fprintf(stderr, []byte("\t\t\t(local)\n\x00"))
	noarch.Fprintf(stderr, []byte("  Elmnt  X-dsp       Y-dsp       Z-dsp       X-rot       Y-rot       Z-rot\n\x00"))
	for m = 1; m <= nE; m++ {
		noarch.Fprintf(stderr, []byte(" %5d %10.6f  %10.6f  %10.6f  %10.6f  %10.6f  %10.6f\n\x00"), m, pkDx[lc][m], pkDy[lc][m], pkDz[lc][m], pkRx[lc][m], pkSy[lc][m], pkSz[lc][m])
	}
}

// write_internal_forces - transpiled function from  $GOPATH/src/github.com/Konstantin8105/History_frame3DD/src/frame3dd_io.c:2674
/*
 * WRITE_INTERNAL_FORCES -
 * calculate frame element internal forces, Nx, Vy, Vz, Tx, My, Mz
 * calculate frame element local displacements, Rx, Dx, Dy, Dz
 * write internal forces and local displacements to an output data file
 * 4jan10, 7mar11, 21jan14
 */ //
/* coord transformation */ //
/* displ. */ //
/* trapz load data, local x dir */ //
/* trapz load data, local y dir */ //
/* trapz load data, local z dir */ //
// distributed loads in local coords at x[i]
// distributed loads in local coords at x[i-1]
// gravity loads in local x, y, z coord's
// distributed torque about local x coord
/* location of internal point loads	*/ //
/* distance along frame element		*/ //
/* axial force within frame el.		*/ //
/* shear forces within frame el.	*/ //
/* torsional moment within frame el.	*/ //
/* bending moments within frame el.	*/ //
/* transverse slopes of frame el.	*/ //
/* frame el. displ. in local x,y,z, dir's */ //
/* twist rotation about the local x-axis */ //
/*  maximum internal forces	*/ //
/*  maximum internal moments	*/ //
/*  maximum element displacements */ //
/*  maximum element rotations	*/ //
/*  minimum internal forces	*/ //
/*  minimum internal moments	*/ //
/*  minimum element displacements */ //
/*  minimum element rotations	*/ //
/* frame element number			*/ //
/* counters for U, W, and P loads	*/ //
/* number of sections alont x axis	*/ //
/* starting and stopping node no's	*/ //
/* file name    for internal force data	*/ //
/* indicate 'write' or 'append' to file */ //
/* file pointer for internal force data */ //
/* file pointer to .CSV output data file */ //
/* modern time variable type		*/ //
// skip calculation of internal forces and displ
/* file name for internal force data for load case "lc" */ //
/* open the interior force data file */ //
//frame3dd.sf.net/");
// write header information for each frame element to txt output data file
// write header information for each frame element to CSV output data file
/*	fprintf(fp,"\n P E A K   I N T E R N A L   D I S P L A C E M E N T S");
 *	fprintf(fp,"\t\t\t(local)\n");
 * 	fprintf(fp,"  Elmnt  X-dsp       Y-dsp       Z-dsp       X-rot       Y-rot       Z-rot\n");
 */ //
// loop over all frame elements
// node 1 and node 2 of elmnt m
// number of x-axis increments
// at least one x-axis increment
// allocate memory for interior force data for frame element "m"
// the local x-axis for frame element "m" starts at 0 and ends at L[m]
// length of the last x-axis increment
// write header information for each frame element
// find interior axial force, shear forces, torsion and bending moments
// distributed gravity load in local x, y, z coordinates
// add uniformly-distributed loads to gravity load
// load n on element m
// interior forces for frame element "m" at (x=0)
// positive Nx is tensile
// positive Vy in local y direction
// positive Vz in local z direction
// positive Tx r.h.r. about local x axis
// positive My -> positive x-z curvature
// positive Mz -> positive x-y curvature
/*  accumulate interior span loads */ //
// start with gravitational plus uniform loads
// add trapezoidally-distributed loads
// load n on element m
// trapezoidal integration of distributed loads
// for axial forces, shear forces and torques
// update distributed loads at x[i-1]
// add interior point loads
// load n on element m
// linear correction of forces for bias in trapezoidal integration
// trapezoidal integration of shear force for bending momemnt
// linear correction of moments for bias in trapezoidal integration
// find interior transverse displacements
/* compute end deflections in local coordinates */ //
// rotations and displacements for frame element "m" at (x=0)
// displacement in  local x dir  at node N1
// displacement in  local y dir  at node N1
// displacement in  local z dir  at node N1
// rotationin about local x axis at node N1
// slope in  local y  direction  at node N1
// slope in  local z  direction  at node N1
// axial displacement along frame element "m"
// linear correction of axial displacement for bias in trapezoidal integration
// torsional rotation along frame element "m"
// linear correction of torsional rot'n for bias in trapezoidal integration
// transverse slope along frame element "m"
// linear correction for bias in trapezoidal integration
// add-in slope due to shear deformation
// displacement along frame element "m"
// linear correction for bias in trapezoidal integration
// initialize the maximum and minimum element forces and displacements
//  maximum internal forces
//  maximum internal moments
//  maximum element displacements
//  maximum element rotations
// find maximum and minimum internal element forces
// find maximum and minimum internal element displacements
// write max and min element forces to the internal frame element force output data file
// write results to the internal frame element force output data file
// write max and min element forces to the Frame3DD text output data file
// write max and min element forces to the Frame3DD CSV output data file
/*
	fprintf(fp," %5d %10.6f  %10.6f  %10.6f  %10.6f  %10.6f  %10.6f\n",
			m, maxDx, maxDy, maxDz, maxRx, maxSy, maxSz );
	fprintf(fp," %5d %10.6f  %10.6f  %10.6f  %10.6f  %10.6f  %10.6f\n",
			m, minDx, minDy, minDz, minRx, minSy, minSz );
*/ //
// free memory
// end of loop over all frame elements
//
func write_internal_forces(OUT_file []byte, fp *noarch.File, infcpath []byte, lc int, nL int, title []byte, dx float32, xyz []vec3, Q [][]float64, nN int, nE int, L []float64, J1 []int, J2 []int, Ax []float32, Asy []float32, Asz []float32, Jx []float32, Iy []float32, Iz []float32, E []float32, G []float32, p []float32, d []float32, gX float32, gY float32, gZ float32, nU int, U [][]float32, nW int, W [][]float32, nP int, P [][]float32, D []float64, shear int, error float64) {
	var t1 float64
	var t2 float64
	var t3 float64
	var t4 float64
	var t5 float64
	var t6 float64
	var t7 float64
	var t8 float64
	var t9 float64
	var u1 float64
	var u2 float64
	var u3 float64
	var u4 float64
	var u5 float64
	var u6 float64
	var u7 float64
	var u8 float64
	var u9 float64
	var u10 float64
	var u11 float64
	var u12 float64
	var xx1 float64
	var xx2 float64
	var wx1 float64
	var wx2 float64
	var xy1 float64
	var xy2 float64
	var wy1 float64
	var wy2 float64
	var xz1 float64
	var xz2 float64
	var wz1 float64
	var wz2 float64
	var wx float64
	var wy float64
	var wz float64
	var wx_ float64
	var wy_ float64
	var wz_ float64
	var wxg float64
	var wyg float64
	var wzg float64
	var tx float64
	var tx_ float64
	var xp float64
	var x []float64
	var dx_ float64
	var dxnx float64
	var Nx []float64
	var Vy []float64
	var Vz []float64
	var Tx []float64
	var My []float64
	var Mz []float64
	var Sy []float64
	var Sz []float64
	var Dx []float64
	var Dy []float64
	var Dz []float64
	var Rx []float64
	var maxNx float64
	var maxVy float64
	var maxVz float64
	var maxTx float64
	var maxMy float64
	var maxMz float64
	var maxDx float64
	var maxDy float64
	var maxDz float64
	var maxRx float64
	var maxSy float64
	var maxSz float64
	var minNx float64
	var minVy float64
	var minVz float64
	var minTx float64
	var minMy float64
	var minMz float64
	var minDx float64
	var minDy float64
	var minDz float64
	var minRx float64
	var minSy float64
	var minSz float64
	var n int
	var m int
	var cU int
	var cW int
	var cP int
	var i int
	var nx int
	var n1 int
	var n2 int
	var i1 int
	var i2 int
	var fnif []byte = make([]byte, 128)
	var CSV_file []byte = make([]byte, 128)
	var errMsg []byte = make([]byte, 512)
	var wa []byte = make([]byte, 4)
	var fpif *noarch.File
	var fpcsv *noarch.File
	var now noarch.TimeT
	if float64(dx) == -1 {
		return
	}
	_ = noarch.Time((*[100000000]noarch.TimeT)(unsafe.Pointer(&now))[:])
	CSV_filename(CSV_file, wa, OUT_file, lc)
	if (func() *noarch.File {
		fpcsv = noarch.Fopen(CSV_file, []byte("a\x00"))
		return fpcsv
	}()) == nil {
		noarch.Sprintf(errMsg, []byte("\n  error: cannot open CSV output data file: %s \n\x00"), CSV_file)
		errorMsg(errMsg)
		os.Exit(17)
	}
	noarch.Sprintf(fnif, []byte("%s%02d\x00"), infcpath, lc)
	if (func() *noarch.File {
		fpif = noarch.Fopen(fnif, []byte("w\x00"))
		return fpif
	}()) == nil {
		noarch.Sprintf(errMsg, []byte("\n  error: cannot open interior force data file: %s \n\x00"), fnif)
		errorMsg(errMsg)
		os.Exit(19)
	}
	noarch.Fprintf(fpif, []byte("# FRAME3DD ANALYSIS RESULTS  http://frame3dd.sf.net/\x00"))
	noarch.Fprintf(fpif, []byte(" VERSION %s \n\x00"), []byte("20140514+\x00"))
	noarch.Fprintf(fpif, []byte("# %s\n\x00"), title)
	noarch.Fprintf(fpif, []byte("# %s\n\x00"), fnif)
	noarch.Fprintf(fpif, []byte("# %s\x00"), noarch.Ctime((*[100000000]noarch.TimeT)(unsafe.Pointer(&now))[:]))
	noarch.Fprintf(fpif, []byte("# L O A D  C A S E   %d  of   %d \n\x00"), lc, nL)
	noarch.Fprintf(fpif, []byte("# F R A M E   E L E M E N T   I N T E R N A L   F O R C E S (local)\n\x00"))
	noarch.Fprintf(fpif, []byte("# F R A M E   E L E M E N T   T R A N S V E R S E   D I S P L A C E M E N T S (local)\n\n\x00"))
	noarch.Fprintf(fp, []byte("\nP E A K   F R A M E   E L E M E N T   I N T E R N A L   F O R C E S\x00"))
	noarch.Fprintf(fp, []byte("(local)\", \n\x00"))
	noarch.Fprintf(fp, []byte("  Elmnt   .         Nx          Vy         Vz\x00"))
	noarch.Fprintf(fp, []byte("        Txx        Myy        Mzz\n\x00"))
	noarch.Fprintf(fpcsv, []byte("\n\"P E A K   F R A M E   E L E M E N T   I N T E R N A L   F O R C E S \x00"))
	noarch.Fprintf(fpcsv, []byte("   (local)\",\n\x00"))
	noarch.Fprintf(fpcsv, []byte(" \"Elmnt\",  \".\", \"Nx\", \"Vy\", \"Vz\", \x00"))
	noarch.Fprintf(fpcsv, []byte(" \"Txx\", \"Myy\", \"Mzz\", \n\x00"))
	for m = 1; m <= nE; m++ {
		n1 = J1[m]
		n2 = J2[m]
		nx = int(math.Floor(L[m] / float64(dx)))
		if nx < 1 {
			nx = 1
		}
		x = dvector(0, int32(nx))
		Nx = dvector(0, int32(nx))
		Vy = dvector(0, int32(nx))
		Vz = dvector(0, int32(nx))
		Tx = dvector(0, int32(nx))
		My = dvector(0, int32(nx))
		Mz = dvector(0, int32(nx))
		Sy = dvector(0, int32(nx))
		Sz = dvector(0, int32(nx))
		Rx = dvector(0, int32(nx))
		Dx = dvector(0, int32(nx))
		Dy = dvector(0, int32(nx))
		Dz = dvector(0, int32(nx))
		for i = 0; i < nx; i++ {
			x[i] = float64(float32(i) * dx)
		}
		x[nx] = L[m]
		dxnx = x[nx] - x[nx-1]
		noarch.Fprintf(fpif, []byte("#\tElmnt\tN1\tN2        \tX1        \tY1        \tZ1        \tX2        \tY2        \tZ2\tnx\n\x00"))
		noarch.Fprintf(fpif, []byte("# @\t%5d\t%5d\t%5d\t%14.6e\t%14.6e\t%14.6e\t%14.6e\t%14.6e\t%14.6e\t%5d\n\x00"), m, n1, n2, xyz[n1].x, xyz[n1].y, xyz[n1].z, xyz[n2].x, xyz[n2].y, xyz[n2].z, nx+1)
		coord_trans(xyz, L[m], n1, n2, (*[100000000]float64)(unsafe.Pointer(&t1))[:], (*[100000000]float64)(unsafe.Pointer(&t2))[:], (*[100000000]float64)(unsafe.Pointer(&t3))[:], (*[100000000]float64)(unsafe.Pointer(&t4))[:], (*[100000000]float64)(unsafe.Pointer(&t5))[:], (*[100000000]float64)(unsafe.Pointer(&t6))[:], (*[100000000]float64)(unsafe.Pointer(&t7))[:], (*[100000000]float64)(unsafe.Pointer(&t8))[:], (*[100000000]float64)(unsafe.Pointer(&t9))[:], p[m])
		wxg = float64(d[m]*Ax[m]) * (t1*float64(gX) + t2*float64(gY) + t3*float64(gZ))
		wyg = float64(d[m]*Ax[m]) * (t4*float64(gX) + t5*float64(gY) + t6*float64(gZ))
		wzg = float64(d[m]*Ax[m]) * (t7*float64(gX) + t8*float64(gY) + t9*float64(gZ))
		for n = 1; n <= nE && cU < nU; n++ {
			if int(U[n][1]) == m {
				wxg += float64(U[n][2])
				wyg += float64(U[n][3])
				wzg += float64(U[n][4])
				cU++
			}
		}
		Nx[0] = -Q[m][1]
		Vy[0] = -Q[m][2]
		Vz[0] = -Q[m][3]
		Tx[0] = -Q[m][4]
		My[0] = Q[m][5]
		Mz[0] = -Q[m][6]
		dx_ = float64(dx)
		for i = 1; i <= nx; i++ {
			wx = wxg
			wy = wyg
			wz = wzg
			if i == 1 {
				wx_ = wxg
				wy_ = wyg
				wz_ = wzg
				tx_ = tx
			}
			for n = 1; n <= 10*nE && cW < nW; n++ {
				if int(W[n][1]) == m {
					if i == nx {
						cW++
					}
					xx1 = float64(W[n][2])
					xx2 = float64(W[n][3])
					wx1 = float64(W[n][4])
					wx2 = float64(W[n][5])
					xy1 = float64(W[n][6])
					xy2 = float64(W[n][7])
					wy1 = float64(W[n][8])
					wy2 = float64(W[n][9])
					xz1 = float64(W[n][10])
					xz2 = float64(W[n][11])
					wz1 = float64(W[n][12])
					wz2 = float64(W[n][13])
					if x[i] > xx1 && x[i] <= xx2 {
						wx += wx1 + (wx2-wx1)*(x[i]-xx1)/(xx2-xx1)
					}
					if x[i] > xy1 && x[i] <= xy2 {
						wy += wy1 + (wy2-wy1)*(x[i]-xy1)/(xy2-xy1)
					}
					if x[i] > xz1 && x[i] <= xz2 {
						wz += wz1 + (wz2-wz1)*(x[i]-xz1)/(xz2-xz1)
					}
				}
			}
			if i == nx {
				dx_ = dxnx
			}
			Nx[i] = Nx[i-1] - 0.5*(wx+wx_)*dx_
			Vy[i] = Vy[i-1] - 0.5*(wy+wy_)*dx_
			Vz[i] = Vz[i-1] - 0.5*(wz+wz_)*dx_
			Tx[i] = Tx[i-1] - 0.5*(tx+tx_)*dx_
			wx_ = wx
			wy_ = wy
			wz_ = wz
			tx_ = tx
			for n = 1; n <= 10*nE && cP < nP; n++ {
				if int(P[n][1]) == m {
					if i == nx {
						cP++
					}
					xp = float64(P[n][5])
					if x[i] <= xp && xp < x[i]+float64(dx) {
						Nx[i] -= float64(P[n][2]) * 0.5 * (1 - (xp-x[i])/float64(dx))
						Vy[i] -= float64(P[n][3]) * 0.5 * (1 - (xp-x[i])/float64(dx))
						Vz[i] -= float64(P[n][4]) * 0.5 * (1 - (xp-x[i])/float64(dx))
					}
					if x[i]-float64(dx) <= xp && xp < x[i] {
						Nx[i] -= float64(P[n][2]) * 0.5 * (1 - (x[i]-float64(dx)-xp)/float64(dx))
						Vy[i] -= float64(P[n][3]) * 0.5 * (1 - (x[i]-float64(dx)-xp)/float64(dx))
						Vz[i] -= float64(P[n][4]) * 0.5 * (1 - (x[i]-float64(dx)-xp)/float64(dx))
					}
				}
			}
		}
		for i = 1; i <= nx; i++ {
			Nx[i] -= (Nx[nx] - Q[m][7]) * float64(i) / float64(nx)
			Vy[i] -= (Vy[nx] - Q[m][8]) * float64(i) / float64(nx)
			Vz[i] -= (Vz[nx] - Q[m][9]) * float64(i) / float64(nx)
			Tx[i] -= (Tx[nx] - Q[m][10]) * float64(i) / float64(nx)
		}
		dx_ = float64(dx)
		for i = 1; i <= nx; i++ {
			if i == nx {
				dx_ = dxnx
			}
			My[i] = My[i-1] - 0.5*(Vz[i]+Vz[i-1])*dx_
			Mz[i] = Mz[i-1] - 0.5*(Vy[i]+Vy[i-1])*dx_
		}
		for i = 1; i <= nx; i++ {
			My[i] -= (My[nx] + Q[m][11]) * float64(i) / float64(nx)
			Mz[i] -= (Mz[nx] - Q[m][12]) * float64(i) / float64(nx)
		}
		i1 = 6 * (n1 - 1)
		i2 = 6 * (n2 - 1)
		u1 = t1*D[i1+1] + t2*D[i1+2] + t3*D[i1+3]
		u2 = t4*D[i1+1] + t5*D[i1+2] + t6*D[i1+3]
		u3 = t7*D[i1+1] + t8*D[i1+2] + t9*D[i1+3]
		u4 = t1*D[i1+4] + t2*D[i1+5] + t3*D[i1+6]
		u5 = t4*D[i1+4] + t5*D[i1+5] + t6*D[i1+6]
		u6 = t7*D[i1+4] + t8*D[i1+5] + t9*D[i1+6]
		u7 = t1*D[i2+1] + t2*D[i2+2] + t3*D[i2+3]
		u8 = t4*D[i2+1] + t5*D[i2+2] + t6*D[i2+3]
		u9 = t7*D[i2+1] + t8*D[i2+2] + t9*D[i2+3]
		u10 = t1*D[i2+4] + t2*D[i2+5] + t3*D[i2+6]
		u11 = t4*D[i2+4] + t5*D[i2+5] + t6*D[i2+6]
		u12 = t7*D[i2+4] + t8*D[i2+5] + t9*D[i2+6]
		Dx[0] = u1
		Dy[0] = u2
		Dz[0] = u3
		Rx[0] = u4
		Sy[0] = u6
		Sz[0] = -u5
		dx_ = float64(dx)
		for i = 1; i <= nx; i++ {
			if i == nx {
				dx_ = dxnx
			}
			Dx[i] = Dx[i-1] + 0.5*(Nx[i-1]+Nx[i])/float64(E[m]*Ax[m])*dx_
		}
		for i = 1; i <= nx; i++ {
			Dx[i] -= (Dx[nx] - u7) * float64(i) / float64(nx)
		}
		dx_ = float64(dx)
		for i = 1; i <= nx; i++ {
			if i == nx {
				dx_ = dxnx
			}
			Rx[i] = Rx[i-1] + 0.5*(Tx[i-1]+Tx[i])/float64(G[m]*Jx[m])*dx_
		}
		for i = 1; i <= nx; i++ {
			Rx[i] -= (Rx[nx] - u10) * float64(i) / float64(nx)
		}
		dx_ = float64(dx)
		for i = 1; i <= nx; i++ {
			if i == nx {
				dx_ = dxnx
			}
			Sy[i] = Sy[i-1] + 0.5*(Mz[i-1]+Mz[i])/float64(E[m]*Iz[m])*dx_
			Sz[i] = Sz[i-1] + 0.5*(My[i-1]+My[i])/float64(E[m]*Iy[m])*dx_
		}
		for i = 1; i <= nx; i++ {
			Sy[i] -= (Sy[nx] - u12) * float64(i) / float64(nx)
			Sz[i] -= (Sz[nx] + u11) * float64(i) / float64(nx)
		}
		if shear != 0 {
			for i = 0; i <= nx; i++ {
				Sy[i] += Vy[i] / float64(G[m]*Asy[m])
				Sz[i] += Vz[i] / float64(G[m]*Asz[m])
			}
		}
		dx_ = float64(dx)
		for i = 1; i <= nx; i++ {
			if i == nx {
				dx_ = dxnx
			}
			Dy[i] = Dy[i-1] + 0.5*(Sy[i-1]+Sy[i])*dx_
			Dz[i] = Dz[i-1] + 0.5*(Sz[i-1]+Sz[i])*dx_
		}
		for i = 1; i <= nx; i++ {
			Dy[i] -= (Dy[nx] - u8) * float64(i) / float64(nx)
			Dz[i] -= (Dz[nx] - u9) * float64(i) / float64(nx)
		}
		minNx = Nx[0]
		maxNx = minNx
		minVy = Vy[0]
		maxVy = minVy
		minVz = Vz[0]
		maxVz = minVz
		minTx = Tx[0]
		maxTx = minTx
		minMy = My[0]
		maxMy = minMy
		minMz = Mz[0]
		maxMz = minMz
		minDx = Dx[0]
		maxDx = minDx
		minDy = Dy[0]
		maxDy = minDy
		minDz = Dz[0]
		maxDz = minDz
		minRx = Rx[0]
		maxRx = minRx
		minSy = Sy[0]
		maxSy = minSy
		minSz = Sz[0]
		maxSz = minSz
		for i = 1; i <= nx; i++ {
			maxNx = func() float64 {
				if Nx[i] > maxNx {
					return Nx[i]
				}
				return maxNx
			}()
			minNx = func() float64 {
				if Nx[i] < minNx {
					return Nx[i]
				}
				return minNx
			}()
			maxVy = func() float64 {
				if Vy[i] > maxVy {
					return Vy[i]
				}
				return maxVy
			}()
			minVy = func() float64 {
				if Vy[i] < minVy {
					return Vy[i]
				}
				return minVy
			}()
			maxVz = func() float64 {
				if Vz[i] > maxVz {
					return Vz[i]
				}
				return maxVz
			}()
			minVz = func() float64 {
				if Vz[i] < minVz {
					return Vz[i]
				}
				return minVz
			}()
			maxTx = func() float64 {
				if Tx[i] > maxTx {
					return Tx[i]
				}
				return maxTx
			}()
			minTx = func() float64 {
				if Tx[i] < minTx {
					return Tx[i]
				}
				return minTx
			}()
			maxMy = func() float64 {
				if My[i] > maxMy {
					return My[i]
				}
				return maxMy
			}()
			minMy = func() float64 {
				if My[i] < minMy {
					return My[i]
				}
				return minMy
			}()
			maxMz = func() float64 {
				if Mz[i] > maxMz {
					return Mz[i]
				}
				return maxMz
			}()
			minMz = func() float64 {
				if Mz[i] < minMz {
					return Mz[i]
				}
				return minMz
			}()
		}
		for i = 1; i <= nx; i++ {
			maxDx = func() float64 {
				if Dx[i] > maxDx {
					return Dx[i]
				}
				return maxDx
			}()
			minDx = func() float64 {
				if Dx[i] < minDx {
					return Dx[i]
				}
				return minDx
			}()
			maxDy = func() float64 {
				if Dy[i] > maxDy {
					return Dy[i]
				}
				return maxDy
			}()
			minDy = func() float64 {
				if Dy[i] < minDy {
					return Dy[i]
				}
				return minDy
			}()
			maxDz = func() float64 {
				if Dz[i] > maxDz {
					return Dz[i]
				}
				return maxDz
			}()
			minDz = func() float64 {
				if Dz[i] < minDz {
					return Dz[i]
				}
				return minDz
			}()
			maxRx = func() float64 {
				if Rx[i] > maxRx {
					return Rx[i]
				}
				return maxRx
			}()
			minRx = func() float64 {
				if Rx[i] < minRx {
					return Rx[i]
				}
				return minRx
			}()
			maxSy = func() float64 {
				if Sy[i] > maxSy {
					return Sy[i]
				}
				return maxSy
			}()
			minSy = func() float64 {
				if Sy[i] < minSy {
					return Sy[i]
				}
				return minSy
			}()
			maxSz = func() float64 {
				if Sz[i] > maxSz {
					return Sz[i]
				}
				return maxSz
			}()
			minSz = func() float64 {
				if Sz[i] < minSz {
					return Sz[i]
				}
				return minSz
			}()
		}
		noarch.Fprintf(fpif, []byte("#                \tNx        \tVy        \tVz        \tTx        \tMy        \tMz        \tDx        \tDy        \tDz         \tRx\t*\n\x00"))
		noarch.Fprintf(fpif, []byte("# MAXIMUM\t%14.6e\t%14.6e\t%14.6e\t%14.6e\t%14.6e\t%14.6e\t%14.6e\t%14.6e\t%14.6e\t%14.6e\n\x00"), maxNx, maxVy, maxVz, maxTx, maxMy, maxMz, maxDx, maxDy, maxDz, maxRx)
		noarch.Fprintf(fpif, []byte("# MINIMUM\t%14.6e\t%14.6e\t%14.6e\t%14.6e\t%14.6e\t%14.6e\t%14.6e\t%14.6e\t%14.6e\t%14.6e\n\x00"), minNx, minVy, minVz, minTx, minMy, minMz, minDx, minDy, minDz, minRx)
		noarch.Fprintf(fpif, []byte("#.x                \tNx        \tVy        \tVz        \tTx       \tMy        \tMz        \tDx        \tDy        \tDz        \tRx\t~\n\x00"))
		for i = 0; i <= nx; i++ {
			noarch.Fprintf(fpif, []byte("%14.6e\t\x00"), x[i])
			noarch.Fprintf(fpif, []byte("%14.6e\t%14.6e\t%14.6e\t\x00"), Nx[i], Vy[i], Vz[i])
			noarch.Fprintf(fpif, []byte("%14.6e\t%14.6e\t%14.6e\t\x00"), Tx[i], My[i], Mz[i])
			noarch.Fprintf(fpif, []byte("%14.6e\t%14.6e\t%14.6e\t%14.6e\n\x00"), Dx[i], Dy[i], Dz[i], Rx[i])
		}
		noarch.Fprintf(fpif, []byte("#---------------------------------------\n\n\n\x00"))
		noarch.Fprintf(fp, []byte(" %5d   max  %10.3f  %10.3f %10.3f %10.3f %10.3f %10.3f\n\x00"), m, maxNx, maxVy, maxVz, maxTx, maxMy, maxMz)
		noarch.Fprintf(fp, []byte(" %5d   min  %10.3f  %10.3f %10.3f %10.3f %10.3f %10.3f\n\x00"), m, minNx, minVy, minVz, minTx, minMy, minMz)
		noarch.Fprintf(fpcsv, []byte(" %5d, \"max\", %10.3f,  %10.3f, %10.3f, %10.3f, %10.3f, %10.3f\n\x00"), m, maxNx, maxVy, maxVz, maxTx, maxMy, maxMz)
		noarch.Fprintf(fpcsv, []byte(" %5d, \"min\", %10.3f,  %10.3f, %10.3f, %10.3f, %10.3f, %10.3f\n\x00"), m, minNx, minVy, minVz, minTx, minMy, minMz)
		free_dvector(x, 0, int32(nx))
		free_dvector(Nx, 0, int32(nx))
		free_dvector(Vy, 0, int32(nx))
		free_dvector(Vz, 0, int32(nx))
		free_dvector(Tx, 0, int32(nx))
		free_dvector(My, 0, int32(nx))
		free_dvector(Mz, 0, int32(nx))
		free_dvector(Rx, 0, int32(nx))
		free_dvector(Sy, 0, int32(nx))
		free_dvector(Sz, 0, int32(nx))
		free_dvector(Dx, 0, int32(nx))
		free_dvector(Dy, 0, int32(nx))
		free_dvector(Dz, 0, int32(nx))
	}
	noarch.Fclose(fpif)
	noarch.Fclose(fpcsv)
}

// write_modal_results - transpiled function from  $GOPATH/src/github.com/Konstantin8105/History_frame3DD/src/frame3dd_io.c:3120
/*
 * WRITE_MODAL_RESULTS -  save modal frequencies and mode shapes
 * 16 Aug 2001
 */ //
/* mode participation factors	*/ //
//
func write_modal_results(fp *noarch.File, nN int, nE int, nI int, DoF int, M [][]float64, f []float64, V [][]float64, total_mass float64, struct_mass float64, iter int, sumR int, nM int, shift float64, lump int, tol float64, ok int) {
	var i int
	var j int
	var k int
	var m int
	var num_modes int
	var mpfX float64
	var mpfY float64
	var mpfZ float64
	var msX []float64
	var msY []float64
	var msZ []float64
	var fs float64
	msX = dvector(1, int32(DoF))
	msY = dvector(1, int32(DoF))
	msZ = dvector(1, int32(DoF))
	for i = 1; i <= DoF; i++ {
		msZ[i] = 0
		msY[i] = msZ[i]
		msX[i] = msY[i]
		for j = 1; j <= DoF; j += 6 {
			msX[i] += M[i][j]
		}
		for j = 2; j <= DoF; j += 6 {
			msY[i] += M[i][j]
		}
		for j = 3; j <= DoF; j += 6 {
			msZ[i] += M[i][j]
		}
	}
	if DoF-sumR > nM {
		num_modes = nM
	} else {
		num_modes = DoF - sumR
	}
	noarch.Fprintf(fp, []byte("\nM O D A L   A N A L Y S I S   R E S U L T S\n\x00"))
	noarch.Fprintf(fp, []byte("  Total Mass:  %e   \x00"), total_mass)
	noarch.Fprintf(fp, []byte("  Structural Mass:  %e \n\x00"), struct_mass)
	noarch.Fprintf(fp, []byte("N O D A L   M A S S E S\x00"))
	noarch.Fprintf(fp, []byte("\t(diagonal of the mass matrix)\t\t\t(global)\n\x00"))
	noarch.Fprintf(fp, []byte("  Node  X-mass      Y-mass      Z-mass\x00"))
	noarch.Fprintf(fp, []byte("      X-inrta     Y-inrta     Z-inrta\n\x00"))
	for j = 1; j <= nN; j++ {
		k = 6 * (j - 1)
		noarch.Fprintf(fp, []byte(" %5d\x00"), j)
		for i = 1; i <= 6; i++ {
			noarch.Fprintf(fp, []byte(" %11.5e\x00"), M[k+i][k+i])
		}
		noarch.Fprintf(fp, []byte("\n\x00"))
	}
	if lump != 0 {
		noarch.Fprintf(fp, []byte("  Lump masses at nodes.\n\x00"))
	} else {
		noarch.Fprintf(fp, []byte("  Use consistent mass matrix.\n\x00"))
	}
	noarch.Fprintf(fp, []byte("N A T U R A L   F R E Q U E N C I E S   & \n\x00"))
	noarch.Fprintf(fp, []byte("M A S S   N O R M A L I Z E D   M O D E   S H A P E S \n\x00"))
	noarch.Fprintf(fp, []byte(" convergence tolerance: %.3e \n\x00"), tol)
	for m = 1; m <= num_modes; m++ {
		mpfX = 0
		for i = 1; i <= DoF; i++ {
			mpfX += V[i][m] * msX[i]
		}
		mpfY = 0
		for i = 1; i <= DoF; i++ {
			mpfY += V[i][m] * msY[i]
		}
		mpfZ = 0
		for i = 1; i <= DoF; i++ {
			mpfZ += V[i][m] * msZ[i]
		}
		noarch.Fprintf(fp, []byte("  MODE %5d:   f= %lf Hz,  T= %lf sec\n\x00"), m, f[m], 1/f[m])
		noarch.Fprintf(fp, []byte("\t\tX- modal participation factor = %12.4e \n\x00"), mpfX)
		noarch.Fprintf(fp, []byte("\t\tY- modal participation factor = %12.4e \n\x00"), mpfY)
		noarch.Fprintf(fp, []byte("\t\tZ- modal participation factor = %12.4e \n\x00"), mpfZ)
		noarch.Fprintf(fp, []byte("  Node    X-dsp       Y-dsp       Z-dsp\x00"))
		noarch.Fprintf(fp, []byte("       X-rot       Y-rot       Z-rot\n\x00"))
		for j = 1; j <= nN; j++ {
			noarch.Fprintf(fp, []byte(" %5d\x00"), j)
			for i = 5; i >= 0; i-- {
				noarch.Fprintf(fp, []byte(" %11.3e\x00"), V[6*j-i][m])
			}
			noarch.Fprintf(fp, []byte("\n\x00"))
		}
	}
	noarch.Fprintf(fp, []byte("M A T R I X    I T E R A T I O N S: %d\n\x00"), iter)
	fs = math.Sqrt(4*3.141592653589793*3.141592653589793*f[nM]*f[nM]+tol) / (2 * 3.141592653589793)
	noarch.Fprintf(fp, []byte("There are %d modes below %f Hz.\x00"), -ok, fs)
	if -ok > nM {
		noarch.Fprintf(fp, []byte(" ... %d modes were not found.\n\x00"), -ok-nM)
		noarch.Fprintf(fp, []byte(" Try increasing the number of modes in \n\x00"))
		noarch.Fprintf(fp, []byte(" order to get the missing modes below %f Hz.\n\x00"), fs)
	} else {
		noarch.Fprintf(fp, []byte(" ... All %d modes were found.\n\x00"), nM)
	}
	free_dvector(msX, 1, int32(DoF))
	free_dvector(msY, 1, int32(DoF))
	free_dvector(msZ, 1, int32(DoF))
	noarch.Fflush(fp)
}

// static_mesh - transpiled function from  $GOPATH/src/github.com/Konstantin8105/History_frame3DD/src/frame3dd_io.c:3213
/*
 * STATIC_MESH  - create mesh data of deformed and undeformed mesh  22 Feb 1999
 * use gnuplot
 * useful gnuplot options: unset xtics ytics ztics border view key
 * This function illustrates how to read the internal force output data file.
 * The internal force output data file contains all the information required
 * to plot deformed meshes, internal axial force, internal shear force, internal
 * torsion, and internal bending moment diagrams.
 */ //
/* coordinates of the frame element number labels */ //
/* indicates plotting in 2D or 3D	*/ //
/* *scanf return value			*/ //
/* frame element number, number of increments */ //
/* node numbers			*/ //
/* coordinates of node n1		*/ //
/* coordinates of node n2		*/ //
/*  line width of deformed mesh		*/ //
/* modern time variable type		*/ //
// write gnuplot plotting script commands
// check for three-dimensional frame
// open plotting script file for writing
// open plotting script file for appending
// file name for deformed mesh data for load case "lc"
// write header, plot-setup cmds, node label, and element label data
// header & node number & element number labels
//frame3dd.sf.net/");
/* fprintf(fpm,"#  X=%d , Y=%d , Z=%d, D3=%d  \n", X,Y,Z,D3_flag); */ //
// 3D plot setup commands
//	 fprintf(fpm,"%c unset label\n", D3 );
// different plot title for each load case
// requires Gnuplot >= 4.6
// 2D plot command
// 3D plot command
// write undeformed mesh data
// open the undeformed mesh data file for writing
//frame3dd.sf.net/");
// i = 6*(n-1);
// i = 6*(n-1);
// no deformed mesh
// write deformed mesh data
// open the deformed mesh data file for writing
//frame3dd.sf.net/");
// open the interior force data file for reading
// file name for internal force data for load case "lc"
// write deformed shape data for each element
/* debugging ... check mesh data
printf("  frel = %3d; m = %3d; n1 =%4d; n2 = %4d; nx = %3d L = %f \n", frel,m,n1,n2,nx,L[m] );
*/ //
//
func static_mesh(IN_file []byte, infcpath []byte, meshpath []byte, plotpath []byte, title []byte, nN int, nE int, nL int, lc int, DoF int, xyz []vec3, L []float64, N1 []int, N2 []int, p []float32, D []float64, exagg_static float64, D3_flag int, anlyz int, dx float32, scale float32) {
	var fpif *noarch.File
	var fpm *noarch.File
	var mx float64
	var my float64
	var mz float64
	var fnif []byte = make([]byte, 128)
	var meshfl []byte = make([]byte, 128)
	var D2 byte = '#'
	var D3 byte = '#'
	var errMsg []byte = make([]byte, 512)
	var ch byte = 'a'
	var sfrv int
	var frel int
	var nx int
	var n1 int
	var n2 int
	var x1 float32
	var y1 float32
	var z1 float32
	var x2 float32
	var y2 float32
	var z2 float32
	var j int
	var m int
	var n int
	var X int
	var Y int
	var Z int
	var lw int = 1
	var now noarch.TimeT
	_ = noarch.Time((*[100000000]noarch.TimeT)(unsafe.Pointer(&now))[:])
	for j = 1; j <= nN; j++ {
		if xyz[j].x != 0 {
			X = 1
		}
		if xyz[j].y != 0 {
			Y = 1
		}
		if xyz[j].z != 0 {
			Z = 1
		}
	}
	if X != 0 && Y != 0 && Z != 0 || D3_flag != 0 {
		D3 = ' '
		D2 = '#'
	} else {
		D3 = '#'
		D2 = ' '
	}
	if lc <= 1 {
		if (func() *noarch.File {
			fpm = noarch.Fopen(plotpath, []byte("w\x00"))
			return fpm
		}()) == nil {
			noarch.Sprintf(errMsg, []byte("\n  error: cannot open gnuplot script file: %s \n\x00"), plotpath)
			errorMsg(errMsg)
			os.Exit(23)
		}
	} else {
		if (func() *noarch.File {
			fpm = noarch.Fopen(plotpath, []byte("a\x00"))
			return fpm
		}()) == nil {
			noarch.Sprintf(errMsg, []byte("\n  error: cannot open gnuplot script file: %s \n\x00"), plotpath)
			errorMsg(errMsg)
			os.Exit(24)
		}
	}
	if lc >= 1 && anlyz != 0 {
		noarch.Sprintf(meshfl, []byte("%sf.%03d\x00"), meshpath, lc)
	}
	if lc <= 1 {
		noarch.Fprintf(fpm, []byte("# FRAME3DD ANALYSIS RESULTS  http://frame3dd.sf.net/\x00"))
		noarch.Fprintf(fpm, []byte(" VERSION %s \n\x00"), []byte("20140514+\x00"))
		noarch.Fprintf(fpm, []byte("# %s\n\x00"), title)
		noarch.Fprintf(fpm, []byte("# %s\x00"), noarch.Ctime((*[100000000]noarch.TimeT)(unsafe.Pointer(&now))[:]))
		noarch.Fprintf(fpm, []byte("# G N U P L O T   S C R I P T   F I L E \n\x00"))
		noarch.Fprintf(fpm, []byte("set autoscale\n\x00"))
		noarch.Fprintf(fpm, []byte("unset border\n\x00"))
		noarch.Fprintf(fpm, []byte("set pointsize 1.0\n\x00"))
		noarch.Fprintf(fpm, []byte("set xtics; set ytics; set ztics; \n\x00"))
		noarch.Fprintf(fpm, []byte("unset zeroaxis\n\x00"))
		noarch.Fprintf(fpm, []byte("unset key\n\x00"))
		noarch.Fprintf(fpm, []byte("unset label\n\x00"))
		noarch.Fprintf(fpm, []byte("set size ratio -1    # 1:1 2D axis scaling \n\x00"))
		noarch.Fprintf(fpm, []byte("# set view equal xyz # 1:1 3D axis scaling \n\x00"))
		noarch.Fprintf(fpm, []byte("# NODE NUMBER LABELS\n\x00"))
		for j = 1; j <= nN; j++ {
			noarch.Fprintf(fpm, []byte("set label ' %d' at %12.4e, %12.4e, %12.4e\n\x00"), j, xyz[j].x, xyz[j].y, xyz[j].z)
		}
		noarch.Fprintf(fpm, []byte("# ELEMENT NUMBER LABELS\n\x00"))
		for m = 1; m <= nE; m++ {
			n1 = N1[m]
			n2 = N2[m]
			mx = 0.5 * (xyz[n1].x + xyz[n2].x)
			my = 0.5 * (xyz[n1].y + xyz[n2].y)
			mz = 0.5 * (xyz[n1].z + xyz[n2].z)
			noarch.Fprintf(fpm, []byte("set label ' %d' at %12.4e, %12.4e, %12.4e\n\x00"), m, mx, my, mz)
		}
		noarch.Fprintf(fpm, []byte("%c set parametric\n\x00"), int(D3))
		noarch.Fprintf(fpm, []byte("%c set view 60, 70, %5.2f \n\x00"), int(D3), float64(scale))
		noarch.Fprintf(fpm, []byte("%c set view equal xyz # 1:1 3D axis scaling \n\x00"), int(D3))
		noarch.Fprintf(fpm, []byte("%c unset key\n\x00"), int(D3))
		noarch.Fprintf(fpm, []byte("%c set xlabel 'x'\n\x00"), int(D3))
		noarch.Fprintf(fpm, []byte("%c set ylabel 'y'\n\x00"), int(D3))
		noarch.Fprintf(fpm, []byte("%c set zlabel 'z'\n\x00"), int(D3))
	}
	noarch.Fprintf(fpm, []byte("set title \"%s\\n\x00"), title)
	noarch.Fprintf(fpm, []byte("analysis file: %s \x00"), IN_file)
	if anlyz != 0 {
		noarch.Fprintf(fpm, []byte("  deflection exaggeration: %.1f \x00"), exagg_static)
		noarch.Fprintf(fpm, []byte("  load case %d of %d \"\n\x00"), lc, nL)
	} else {
		noarch.Fprintf(fpm, []byte("  data check only \"\n\x00"))
	}
	noarch.Fprintf(fpm, []byte("unset clip; \nset clip one; set clip two\n\x00"))
	noarch.Fprintf(fpm, []byte("set xyplane 0 \n\x00"))
	noarch.Fprintf(fpm, []byte("%c plot '%s' u 2:3 t 'undeformed mesh' w lp \x00"), int(D2), meshpath)
	if noarch.NotInt(anlyz) != 0 {
		noarch.Fprintf(fpm, []byte("lw %d lt 1 pt 6 \n\x00"), lw)
	} else {
		noarch.Fprintf(fpm, []byte("lw 1 lt 5 pt 6, '%s' u 1:2 t 'load case %d of %d' w l lw %d lt 3\n\x00"), meshfl, lc, nL, lw)
	}
	noarch.Fprintf(fpm, []byte("%c splot '%s' u 2:3:4 t 'load case %d of %d' w lp \x00"), int(D3), meshpath, lc, nL)
	if noarch.NotInt(anlyz) != 0 {
		noarch.Fprintf(fpm, []byte(" lw %d lt 1 pt 6 \n\x00"), lw)
	} else {
		noarch.Fprintf(fpm, []byte(" lw 1 lt 5 pt 6, '%s' u 1:2:3 t 'load case %d of %d' w l lw %d lt 3\n\x00"), meshfl, lc, nL, lw)
	}
	if lc < nL && anlyz != 0 {
		noarch.Fprintf(fpm, []byte("pause -1\n\x00"))
	}
	noarch.Fclose(fpm)
	if lc <= 1 {
		if (func() *noarch.File {
			fpm = noarch.Fopen(meshpath, []byte("w\x00"))
			return fpm
		}()) == nil {
			noarch.Sprintf(errMsg, []byte("\n  error: cannot open gnuplot undeformed mesh data file: %s\n\x00"), meshpath)
			errorMsg(errMsg)
			os.Exit(21)
		}
		noarch.Fprintf(fpm, []byte("# FRAME3DD ANALYSIS RESULTS  http://frame3dd.sf.net/\x00"))
		noarch.Fprintf(fpm, []byte(" VERSION %s \n\x00"), []byte("20140514+\x00"))
		noarch.Fprintf(fpm, []byte("# %s\n\x00"), title)
		noarch.Fprintf(fpm, []byte("# %s\x00"), noarch.Ctime((*[100000000]noarch.TimeT)(unsafe.Pointer(&now))[:]))
		noarch.Fprintf(fpm, []byte("# U N D E F O R M E D   M E S H   D A T A   (global coordinates)\n\x00"))
		noarch.Fprintf(fpm, []byte("# Node        X            Y            Z \n\x00"))
		for m = 1; m <= nE; m++ {
			n = N1[m]
			noarch.Fprintf(fpm, []byte("%5d %12.4e %12.4e %12.4e \n\x00"), n, xyz[n].x, xyz[n].y, xyz[n].z)
			n = N2[m]
			noarch.Fprintf(fpm, []byte("%5d %12.4e %12.4e %12.4e\x00"), n, xyz[n].x, xyz[n].y, xyz[n].z)
			noarch.Fprintf(fpm, []byte("\n\n\n\x00"))
		}
		noarch.Fclose(fpm)
	}
	if noarch.NotInt(anlyz) != 0 {
		return
	}
	if (func() *noarch.File {
		fpm = noarch.Fopen(meshfl, []byte("w\x00"))
		return fpm
	}()) == nil {
		noarch.Sprintf(errMsg, []byte("\n  error: cannot open gnuplot deformed mesh data file %s \n\x00"), meshfl)
		errorMsg(errMsg)
		os.Exit(22)
	}
	noarch.Fprintf(fpm, []byte("# FRAME3DD ANALYSIS RESULTS  http://frame3dd.sf.net/\x00"))
	noarch.Fprintf(fpm, []byte(" VERSION %s \n\x00"), []byte("20140514+\x00"))
	noarch.Fprintf(fpm, []byte("# %s\n\x00"), title)
	noarch.Fprintf(fpm, []byte("# L O A D  C A S E   %d  of   %d \n\x00"), lc, nL)
	noarch.Fprintf(fpm, []byte("# %s\x00"), noarch.Ctime((*[100000000]noarch.TimeT)(unsafe.Pointer(&now))[:]))
	noarch.Fprintf(fpm, []byte("# D E F O R M E D   M E S H   D A T A \x00"))
	noarch.Fprintf(fpm, []byte("  deflection exaggeration: %.1f\n\x00"), exagg_static)
	noarch.Fprintf(fpm, []byte("#       X-dsp        Y-dsp        Z-dsp\n\x00"))
	if float64(dx) > 0 && anlyz != 0 {
		noarch.Sprintf(fnif, []byte("%s%02d\x00"), infcpath, lc)
		if (func() *noarch.File {
			fpif = noarch.Fopen(fnif, []byte("r\x00"))
			return fpif
		}()) == nil {
			noarch.Sprintf(errMsg, []byte("\n  error: cannot open interior force data file: %s \n\x00"), fnif)
			errorMsg(errMsg)
			os.Exit(20)
		}
	}
	for m = 1; m <= nE; m++ {
		ch = 'a'
		noarch.Fprintf(fpm, []byte("\n# element %5d \n\x00"), m)
		if float64(dx) < 0 && anlyz != 0 {
			cubic_bent_beam(fpm, N1[m], N2[m], xyz, L[m], p[m], D, exagg_static)
		}
		if float64(dx) > 0 && anlyz != 0 {
			for int(ch) != int('@') {
				ch = byte(noarch.Fgetc(fpif))
			}
			sfrv = noarch.Fscanf(fpif, []byte("%d %d %d %f %f %f %f %f %f %d\x00"), (*[100000000]int)(unsafe.Pointer(&frel))[:], (*[100000000]int)(unsafe.Pointer(&n1))[:], (*[100000000]int)(unsafe.Pointer(&n2))[:], (*[100000000]float32)(unsafe.Pointer(&x1))[:], (*[100000000]float32)(unsafe.Pointer(&y1))[:], (*[100000000]float32)(unsafe.Pointer(&z1))[:], (*[100000000]float32)(unsafe.Pointer(&x2))[:], (*[100000000]float32)(unsafe.Pointer(&y2))[:], (*[100000000]float32)(unsafe.Pointer(&z2))[:], (*[100000000]int)(unsafe.Pointer(&nx))[:])
			if sfrv != 10 {
				sferr(fnif)
			}
			if frel != m || N1[m] != n1 || N2[m] != n2 {
				noarch.Fprintf(stderr, []byte(" error in static_mesh parsing\n\x00"))
				noarch.Fprintf(stderr, []byte("  frel = %d; m = %d; nx = %d \n\x00"), frel, m, nx)
			}
			for int(ch) != int('~') {
				ch = byte(noarch.Fgetc(fpif))
			}
			force_bent_beam(fpm, fpif, fnif, nx, N1[m], N2[m], xyz, L[m], p[m], D, exagg_static)
		}
	}
	if float64(dx) > 0 && anlyz != 0 {
		noarch.Fclose(fpif)
	}
	noarch.Fclose(fpm)
}

// modal_mesh - transpiled function from  $GOPATH/src/github.com/Konstantin8105/History_frame3DD/src/frame3dd_io.c:3449
/*
 * MODAL_MESH  -  create mesh data of the mode-shape meshes, use gnuplot	19oct98
 * useful gnuplot options: unset xtics ytics ztics border view key
 */ //
/* mode participation factors	*/ //
/* a mode-shape vector */ //
/*  line thickness of deformed mesh	*/ //
/* indicate 2D or 3D frame	*/ //
/* modal participation factors */ //
//frame3dd.sf.net/");
// check for three-dimensional frame
// 2D plot command
// 3D plot command
//
func modal_mesh(IN_file []byte, meshpath []byte, modepath []byte, plotpath []byte, title []byte, nN int, nE int, DoF int, nM int, xyz []vec3, L []float64, J1 []int, J2 []int, p []float32, M [][]float64, f []float64, V [][]float64, exagg_modal float64, D3_flag int, anlyz int) {
	var fpm *noarch.File
	var mpfX float64
	var mpfY float64
	var mpfZ float64
	var msX []float64
	var msY []float64
	var msZ []float64
	var v []float64
	var i int
	var j int
	var m int
	var n int
	var X int
	var Y int
	var Z int
	var lw int = 1
	var D2 byte = '#'
	var D3 byte = '#'
	var modefl []byte = make([]byte, 128)
	var errMsg []byte = make([]byte, 512)
	msX = dvector(1, int32(DoF))
	msY = dvector(1, int32(DoF))
	msZ = dvector(1, int32(DoF))
	v = dvector(1, int32(DoF))
	for i = 1; i <= DoF; i++ {
		msZ[i] = 0
		msY[i] = msZ[i]
		msX[i] = msY[i]
		for j = 1; j <= DoF; j += 6 {
			msX[i] += M[i][j]
		}
		for j = 2; j <= DoF; j += 6 {
			msY[i] += M[i][j]
		}
		for j = 3; j <= DoF; j += 6 {
			msZ[i] += M[i][j]
		}
	}
	if noarch.NotInt(anlyz) != 0 {
		exagg_modal = 0
	}
	for m = 1; m <= nM; m++ {
		noarch.Sprintf(modefl, []byte("%s-%02d-\x00"), modepath, m)
		if (func() *noarch.File {
			fpm = noarch.Fopen(modefl, []byte("w\x00"))
			return fpm
		}()) == nil {
			noarch.Sprintf(errMsg, []byte("\n  error: cannot open gnuplot modal mesh file: %s \n\x00"), modefl)
			errorMsg(errMsg)
			os.Exit(27)
		}
		noarch.Fprintf(fpm, []byte("# FRAME3DD ANALYSIS RESULTS  http://frame3dd.sf.net/\x00"))
		noarch.Fprintf(fpm, []byte(" VERSION %s \n\x00"), []byte("20140514+\x00"))
		noarch.Fprintf(fpm, []byte("# %s\n\x00"), title)
		noarch.Fprintf(fpm, []byte("# M O D E   S H A P E   D A T A   F O R   M O D E\x00"))
		noarch.Fprintf(fpm, []byte("   %d\t(global coordinates)\n\x00"), m)
		noarch.Fprintf(fpm, []byte("# deflection exaggeration: %.1f\n\n\x00"), exagg_modal)
		mpfX = 0
		for i = 1; i <= DoF; i++ {
			mpfX += V[i][m] * msX[i]
		}
		mpfY = 0
		for i = 1; i <= DoF; i++ {
			mpfY += V[i][m] * msY[i]
		}
		mpfZ = 0
		for i = 1; i <= DoF; i++ {
			mpfZ += V[i][m] * msZ[i]
		}
		noarch.Fprintf(fpm, []byte("# MODE %5d:   f= %lf Hz, T= %lf sec\n\x00"), m, f[m], 1/f[m])
		noarch.Fprintf(fpm, []byte("#\t\tX- modal participation factor = %12.4e \n\x00"), mpfX)
		noarch.Fprintf(fpm, []byte("#\t\tY- modal participation factor = %12.4e \n\x00"), mpfY)
		noarch.Fprintf(fpm, []byte("#\t\tZ- modal participation factor = %12.4e \n\x00"), mpfZ)
		for i = 1; i <= DoF; i++ {
			v[i] = V[i][m]
		}
		noarch.Fprintf(fpm, []byte("#      X-dsp       Y-dsp       Z-dsp\n\n\x00"))
		for n = 1; n <= nE; n++ {
			noarch.Fprintf(fpm, []byte("\n# element %5d \n\x00"), n)
			cubic_bent_beam(fpm, J1[n], J2[n], xyz, L[n], p[n], v, exagg_modal)
		}
		noarch.Fclose(fpm)
		for j = 1; j <= nN; j++ {
			if xyz[j].x != 0 {
				X = 1
			}
			if xyz[j].y != 0 {
				Y = 1
			}
			if xyz[j].z != 0 {
				Z = 1
			}
		}
		if X != 0 && Y != 0 && Z != 0 || D3_flag != 0 {
			D3 = ' '
			D2 = '#'
		} else {
			D3 = '#'
			D2 = ' '
		}
		if (func() *noarch.File {
			fpm = noarch.Fopen(plotpath, []byte("a\x00"))
			return fpm
		}()) == nil {
			noarch.Sprintf(errMsg, []byte("\n  error: cannot append gnuplot script file: %s \n\x00"), plotpath)
			errorMsg(errMsg)
			os.Exit(25)
		}
		noarch.Fprintf(fpm, []byte("pause -1\n\x00"))
		if m == 1 {
			noarch.Fprintf(fpm, []byte("unset label\n\x00"))
			noarch.Fprintf(fpm, []byte("%c unset key\n\x00"), int(D3))
		}
		noarch.Fprintf(fpm, []byte("set title '%s     mode %d     %lf Hz'\n\x00"), IN_file, m, f[m])
		noarch.Fprintf(fpm, []byte("%c plot '%s' u 2:3 t 'undeformed mesh' w l \x00"), int(D2), meshpath)
		if noarch.NotInt(anlyz) != 0 {
			noarch.Fprintf(fpm, []byte(" lw %d lt 1 \n\x00"), lw)
		} else {
			noarch.Fprintf(fpm, []byte(" lw 1 lt 5 , '%s' u 1:2 t 'mode-shape %d' w l lw %d lt 3\n\x00"), modefl, m, lw)
		}
		noarch.Fprintf(fpm, []byte("%c splot '%s' u 2:3:4 t 'undeformed mesh' w l \x00"), int(D3), meshpath)
		if noarch.NotInt(anlyz) != 0 {
			noarch.Fprintf(fpm, []byte(" lw %d lt 1 \n\x00"), lw)
		} else {
			noarch.Fprintf(fpm, []byte(" lw 1 lt 5 , '%s' u 1:2:3 t 'mode-shape %d' w l lw %d lt 3\n\x00"), modefl, m, lw)
		}
		noarch.Fclose(fpm)
	}
	free_dvector(msX, 1, int32(DoF))
	free_dvector(msY, 1, int32(DoF))
	free_dvector(msZ, 1, int32(DoF))
	free_dvector(v, 1, int32(DoF))
}

// animate - transpiled function from  $GOPATH/src/github.com/Konstantin8105/History_frame3DD/src/frame3dd_io.c:3577
/*
 * ANIMATE -  create mesh data of animated mode-shape meshes, use gnuplot	16dec98
 * useful gnuplot options: unset xtics ytics ztics border view key
 * mpeg movie example:   % convert mesh_file-03-f-*.ps mode-03.mpeg
 * ... requires ImageMagick and mpeg2vidcodec packages
 */ //
/* pan rate for animation	     */ //
/* inital zoom scale in 3D animation */ //
/* "diameter" of the structure	*/ //
/* inital x-rotation in 3D animation */ //
/* final  x-rotation in 3D animation */ //
/* inital z-rotation in 3D animation */ //
/* final  z-rotation in 3D animation */ //
/* init.  zoom scale in 3D animation */ //
/* final  zoom scale in 3D animation */ //
/* number of frames in animation */ //
/* an exageration factor, for animation */ //
/*  line thickness of deformed mesh	*/ //
/* total number of frames in animation	*/ //
/* indicate 2D or 3D frame	*/ //
/* use '#' for no-movie  -OR-  ' ' for movie */ //
// check for three-dimensional frame
/*
 *	   if ( x_min != x_max )
 *		fprintf(fpm,"set xrange [ %lf : %lf ] \n",
 *	 		x_min-0.2*(x_max-x_min), x_max+0.2*(x_max-x_min) );
 *	   else fprintf(fpm,"set xrange [ %lf : %lf ] \n",
 *			x_min-exagg_modal, x_max+exagg_modal );
 *	   if (y_min != y_max)
 *		fprintf(fpm,"set yrange [ %lf : %lf ] \n",
 *	 		y_min-0.2*(y_max-y_min), y_max+0.2*(y_max-y_min) );
 *	   else fprintf(fpm,"set yrange [ %lf : %lf ] \n",
 *			y_min-exagg_modal, y_max+exagg_modal );
 *	   if (z_min != z_max)
 *	   	fprintf(fpm,"set zrange [ %lf : %lf ] \n",
 *			z_min-0.2*(z_max-z_min), z_max+0.2*(z_max-z_min) );
 *	   else fprintf(fpm,"set zrange [ %lf : %lf ] \n",
 *			z_min-exagg_modal, z_max+exagg_modal );
 */ //
//frame3dd.sf.net/");
/* mode "m" */ //
//
func animate(IN_file []byte, meshpath []byte, modepath []byte, plotpath []byte, title []byte, anim []int, nN int, nE int, DoF int, nM int, xyz []vec3, L []float64, p []float32, J1 []int, J2 []int, f []float64, V [][]float64, exagg_modal float64, D3_flag int, pan float32, scale float32) {
	var fpm *noarch.File
	var x_min float32
	var x_max float32
	var y_min float32
	var y_max float32
	var z_min float32
	var z_max float32
	var Dxyz float32
	var rot_x_init float32
	var rot_x_final float32
	var rot_z_init float32
	var rot_z_final float32
	var zoom_init float32 = float32(0 * float64(scale))
	var zoom_final float32 = float32(0 * float64(scale))
	var frames float32 = 25
	var ex float64 = 10
	var v []float64
	var fr int
	var i int
	var j int
	var m int
	var n int
	var X int
	var Y int
	var Z int
	var c int
	var CYCLES int = 3
	var frame_number int
	var lw int = 1
	var total_frames int
	var D2 byte = '#'
	var D3 byte = '#'
	var Movie byte = '#'
	var modefl []byte = make([]byte, 128)
	var framefl []byte = make([]byte, 128)
	var errMsg []byte = make([]byte, 512)
	for j = 1; j <= nN; j++ {
		if xyz[j].x != 0 {
			X = 1
		}
		if xyz[j].y != 0 {
			Y = 1
		}
		if xyz[j].z != 0 {
			Z = 1
		}
		if j == 1 {
			x_max = float32(xyz[j].x)
			x_min = x_max
			y_max = float32(xyz[j].y)
			y_min = y_max
			z_max = float32(xyz[j].z)
			z_min = z_max
		}
		if xyz[j].x < float64(x_min) {
			x_min = float32(xyz[j].x)
		}
		if xyz[j].y < float64(y_min) {
			y_min = float32(xyz[j].y)
		}
		if xyz[j].z < float64(z_min) {
			z_min = float32(xyz[j].z)
		}
		if float64(x_max) < xyz[j].x {
			x_max = float32(xyz[j].x)
		}
		if float64(y_max) < xyz[j].y {
			y_max = float32(xyz[j].y)
		}
		if float64(z_max) < xyz[j].z {
			z_max = float32(xyz[j].z)
		}
	}
	if X != 0 && Y != 0 && Z != 0 || D3_flag != 0 {
		D3 = ' '
		D2 = '#'
	} else {
		D3 = '#'
		D2 = ' '
	}
	Dxyz = float32(math.Sqrt(float64((x_max-x_min)*(x_max-x_min) + (y_max-y_min)*(y_max-y_min) + (z_max-z_min)*(z_max-z_min))))
	if (func() *noarch.File {
		fpm = noarch.Fopen(plotpath, []byte("a\x00"))
		return fpm
	}()) == nil {
		noarch.Sprintf(errMsg, []byte("\n  error: cannot append gnuplot script file: %s \n\x00"), plotpath)
		errorMsg(errMsg)
		os.Exit(26)
	}
	i = 1
	for (func() int {
		m = anim[i]
		return m
	}()) != 0 && i < 100 {
		if i == 1 {
			noarch.Fprintf(fpm, []byte("\n# --- M O D E   S H A P E   A N I M A T I O N ---\n\x00"))
			noarch.Fprintf(fpm, []byte("# rot_x_init  = %7.2f\n\x00"), float64(rot_x_init))
			noarch.Fprintf(fpm, []byte("# rot_x_final = %7.2f\n\x00"), float64(rot_x_final))
			noarch.Fprintf(fpm, []byte("# rot_z_init  = %7.2f\n\x00"), float64(rot_z_init))
			noarch.Fprintf(fpm, []byte("# rot_z_final = %7.2f\n\x00"), float64(rot_z_final))
			noarch.Fprintf(fpm, []byte("# zoom_init   = %7.2f\n\x00"), float64(zoom_init))
			noarch.Fprintf(fpm, []byte("# zoom_final  = %7.2f\n\x00"), float64(zoom_init))
			noarch.Fprintf(fpm, []byte("# pan rate    = %7.2f \n\x00"), float64(pan))
			noarch.Fprintf(fpm, []byte("set autoscale\n\x00"))
			noarch.Fprintf(fpm, []byte("unset border\n\x00"))
			noarch.Fprintf(fpm, []byte("%c unset xlabel \n\x00"), int(D3))
			noarch.Fprintf(fpm, []byte("%c unset ylabel \n\x00"), int(D3))
			noarch.Fprintf(fpm, []byte("%c unset zlabel \n\x00"), int(D3))
			noarch.Fprintf(fpm, []byte("%c unset label \n\x00"), int(D3))
			noarch.Fprintf(fpm, []byte("unset key\n\x00"))
			noarch.Fprintf(fpm, []byte("%c set parametric\n\x00"), int(D3))
			noarch.Fprintf(fpm, []byte("# x_min = %12.5e     x_max = %12.5e \n\x00"), float64(x_min), float64(x_max))
			noarch.Fprintf(fpm, []byte("# y_min = %12.5e     y_max = %12.5e \n\x00"), float64(y_min), float64(y_max))
			noarch.Fprintf(fpm, []byte("# z_min = %12.5e     z_max = %12.5e \n\x00"), float64(z_min), float64(z_max))
			noarch.Fprintf(fpm, []byte("# Dxyz = %12.5e \n\x00"), float64(Dxyz))
			noarch.Fprintf(fpm, []byte("set xrange [ %lf : %lf ] \n\x00"), float64(x_min)-0.2*float64(Dxyz), float64(x_max)+0.1*float64(Dxyz))
			noarch.Fprintf(fpm, []byte("set yrange [ %lf : %lf ] \n\x00"), float64(y_min)-0.2*float64(Dxyz), float64(y_max)+0.1*float64(Dxyz))
			noarch.Fprintf(fpm, []byte("set zrange [ %lf : %lf ] \n\x00"), float64(z_min)-0.2*float64(Dxyz), float64(z_max)+0.1*float64(Dxyz))
			noarch.Fprintf(fpm, []byte("unset xzeroaxis; unset yzeroaxis; unset zzeroaxis\n\x00"))
			noarch.Fprintf(fpm, []byte("unset xtics; unset ytics; unset ztics; \n\x00"))
			noarch.Fprintf(fpm, []byte("%c set view 60, 70, %5.2f \n\x00"), int(D3), float64(scale))
			noarch.Fprintf(fpm, []byte("set size ratio -1    # 1:1 2D axis scaling \n\x00"))
			noarch.Fprintf(fpm, []byte("%c set view equal xyz # 1:1 3D axis scaling \n\x00"), int(D3))
		}
		noarch.Fprintf(fpm, []byte("pause -1 \n\x00"))
		noarch.Fprintf(fpm, []byte("set title '%s     mode %d      %lf Hz'\n\x00"), IN_file, m, f[m])
		frame_number = 0
		total_frames = int(float32(2*CYCLES) * frames)
		for c = 1; c <= CYCLES; c++ {
			for fr = 0; float32(fr) <= frames; fr++ {
				frame_number++
				noarch.Sprintf(modefl, []byte("%s-%02d.%03d\x00"), modepath, m, fr)
				noarch.Sprintf(framefl, []byte("%s-%02d-f-%03d.ps\x00"), modepath, m, fr)
				noarch.Fprintf(fpm, []byte("%c plot '%s' u 2:3 w l lw 1 lt 5, \x00"), int(D2), meshpath)
				noarch.Fprintf(fpm, []byte(" '%s' u 1:2 w l lw %d lt 3 ; \n\x00"), modefl, lw)
				if float64(pan) != 0 {
					noarch.Fprintf(fpm, []byte("%c set view %7.2f, %7.2f, %5.3f # pan = %f\n\x00"), int(D3), float64(rot_x_init+pan*(rot_x_final-rot_x_init)*float32(frame_number)/float32(total_frames)), float64(rot_z_init+pan*(rot_z_final-rot_z_init)*float32(frame_number)/float32(total_frames)), float64(zoom_init+pan*(zoom_final-zoom_init)*float32(frame_number)/float32(total_frames)), float64(pan))
				}
				noarch.Fprintf(fpm, []byte("%c splot '%s' u 2:3:4 w l lw 1 lt 5, \x00"), int(D3), meshpath)
				noarch.Fprintf(fpm, []byte(" '%s' u 1:2:3 w l lw %d lt 3;\x00"), modefl, lw)
				if fr == 0 && c == 1 {
					noarch.Fprintf(fpm, []byte("  pause 1.5 \n\x00"))
				} else {
					noarch.Fprintf(fpm, []byte("  pause 0.05 \n\x00"))
				}
				noarch.Fprintf(fpm, []byte("%c  load 'saveplot';\n\x00"), int(Movie))
				noarch.Fprintf(fpm, []byte("%c  !mv my-plot.ps %s\n\x00"), int(Movie), framefl)
			}
			for fr = int(frames - 1); fr > 0; fr-- {
				frame_number++
				noarch.Sprintf(modefl, []byte("%s-%02d.%03d\x00"), modepath, m, fr)
				noarch.Sprintf(framefl, []byte("%s-%02d-f-%03d.ps\x00"), modepath, m, fr)
				noarch.Fprintf(fpm, []byte("%c plot '%s' u 2:3 w l lw 1 lt 5, \x00"), int(D2), meshpath)
				noarch.Fprintf(fpm, []byte(" '%s' u 1:2 w l lw %d lt 3; \n\x00"), modefl, lw)
				if float64(pan) != 0 {
					noarch.Fprintf(fpm, []byte("%c set view %7.2f, %7.2f, %5.3f # pan = %f\n\x00"), int(D3), float64(rot_x_init+pan*(rot_x_final-rot_x_init)*float32(frame_number)/float32(total_frames)), float64(rot_z_init+pan*(rot_z_final-rot_z_init)*float32(frame_number)/float32(total_frames)), float64(zoom_init+pan*(zoom_final-zoom_init)*float32(frame_number)/float32(total_frames)), float64(pan))
				}
				noarch.Fprintf(fpm, []byte("%c splot '%s' u 2:3:4 w l lw 1 lt 5, \x00"), int(D3), meshpath)
				noarch.Fprintf(fpm, []byte(" '%s' u 1:2:3 w l lw %d lt 3;\x00"), modefl, lw)
				noarch.Fprintf(fpm, []byte("  pause 0.05 \n\x00"))
				noarch.Fprintf(fpm, []byte("%c  load 'saveplot';\n\x00"), int(Movie))
				noarch.Fprintf(fpm, []byte("%c  !mv my-plot.ps %s\n\x00"), int(Movie), framefl)
			}
		}
		fr = 0
		noarch.Sprintf(modefl, []byte("%s-%02d.%03d\x00"), modepath, m, fr)
		noarch.Fprintf(fpm, []byte("%c plot '%s' u 2:3 w l lw %d lt 5, \x00"), int(D2), meshpath, lw)
		noarch.Fprintf(fpm, []byte(" '%s' u 1:2 w l lw 3 lt 3 \n\x00"), modefl)
		noarch.Fprintf(fpm, []byte("%c splot '%s' u 2:3:4 w l lw %d lt 5, \x00"), int(D3), meshpath, lw)
		noarch.Fprintf(fpm, []byte(" '%s' u 1:2:3 w l lw 3 lt 3 \n\x00"), modefl)
		i++
	}
	noarch.Fclose(fpm)
	v = dvector(1, int32(DoF))
	i = 1
	for (func() int {
		m = anim[i]
		return m
	}()) != 0 {
		for fr = 0; float32(fr) <= frames; fr++ {
			noarch.Sprintf(modefl, []byte("%s-%02d.%03d\x00"), modepath, m, fr)
			if (func() *noarch.File {
				fpm = noarch.Fopen(modefl, []byte("w\x00"))
				return fpm
			}()) == nil {
				noarch.Sprintf(errMsg, []byte("\n  error: cannot open gnuplot modal mesh data file: %s \n\x00"), modefl)
				errorMsg(errMsg)
				os.Exit(28)
			}
			ex = exagg_modal * math.Cos(3.141592653589793*float64(fr)/float64(frames))
			noarch.Fprintf(fpm, []byte("# FRAME3DD ANALYSIS RESULTS  http://frame3dd.sf.net/\x00"))
			noarch.Fprintf(fpm, []byte(" VERSION %s \n\x00"), []byte("20140514+\x00"))
			noarch.Fprintf(fpm, []byte("# %s\n\x00"), title)
			noarch.Fprintf(fpm, []byte("# A N I M A T E D   M O D E   S H A P E   D A T A \n\x00"))
			noarch.Fprintf(fpm, []byte("# deflection exaggeration: %.1f\n\x00"), ex)
			noarch.Fprintf(fpm, []byte("# MODE %5d: f= %lf Hz  T= %lf sec\n\n\x00"), m, f[m], 1/f[m])
			for j = 1; j <= DoF; j++ {
				v[j] = V[j][m]
			}
			noarch.Fprintf(fpm, []byte("#      X-dsp       Y-dsp       Z-dsp\n\n\x00"))
			for n = 1; n <= nE; n++ {
				noarch.Fprintf(fpm, []byte("\n# element %5d \n\x00"), n)
				cubic_bent_beam(fpm, J1[n], J2[n], xyz, L[n], p[n], v, ex)
			}
			noarch.Fclose(fpm)
		}
		i++
	}
	free_dvector(v, 1, int32(DoF))
}

// cubic_bent_beam - transpiled function from  $GOPATH/src/github.com/Konstantin8105/History_frame3DD/src/frame3dd_io.c:3816
/*
 * CUBIC_BENT_BEAM  -  computes cubic deflection functions from end deflections
 * and end rotations.  Saves deflected shapes to a file.  These bent shapes
 * are exact for mode-shapes, and for frames loaded at their nodes.
 * 15 May 2009
 */ //
/* coord xfmn	*/ //
/* u4,*/ //
/* u10,*/ //
/* compute end deflections in local coordinates */ //
/* u4  = exagg*(t1*D[i1+4] + t2*D[i1+5] + t3*D[i1+6]); */ //
/* u10 = exagg*(t1*D[i2+4] + t2*D[i2+5] + t3*D[i2+6]); */ //
/* curve-fitting problem for a cubic polynomial */ //
/* solve for cubic coef's */ //
/* solve for cubic coef's */ //
// debug ... if deformed mesh exageration is too big, some elements
// may not be plotted.
//fprintf( fpm, "# u1=%e  L+u7=%e, dx = %e \n",
//				u1, fabs(L+u7), fabs(L+u7-u1)/10.0);
/* deformed shape in local coordinates */ //
/* deformed shape in global coordinates */ //
//
func cubic_bent_beam(fpm *noarch.File, n1 int, n2 int, xyz []vec3, L float64, p float32, D []float64, exagg float64) {
	var t1 float64
	var t2 float64
	var t3 float64
	var t4 float64
	var t5 float64
	var t6 float64
	var t7 float64
	var t8 float64
	var t9 float64
	var u1 float64
	var u2 float64
	var u3 float64
	var u5 float64
	var u6 float64
	var u7 float64
	var u8 float64
	var u9 float64
	var u11 float64
	var u12 float64
	var a []float64
	var b []float64
	var A [][]float64
	var s float64
	var v float64
	var w float64
	var dX float64
	var dY float64
	var dZ float64
	var i1 int
	var i2 int
	var pd int
	var errMsg []byte = make([]byte, 512)
	A = dmatrix(1, 4, 1, 4)
	a = dvector(1, 4)
	b = dvector(1, 4)
	coord_trans(xyz, L, n1, n2, (*[100000000]float64)(unsafe.Pointer(&t1))[:], (*[100000000]float64)(unsafe.Pointer(&t2))[:], (*[100000000]float64)(unsafe.Pointer(&t3))[:], (*[100000000]float64)(unsafe.Pointer(&t4))[:], (*[100000000]float64)(unsafe.Pointer(&t5))[:], (*[100000000]float64)(unsafe.Pointer(&t6))[:], (*[100000000]float64)(unsafe.Pointer(&t7))[:], (*[100000000]float64)(unsafe.Pointer(&t8))[:], (*[100000000]float64)(unsafe.Pointer(&t9))[:], p)
	i1 = 6 * (n1 - 1)
	i2 = 6 * (n2 - 1)
	u1 = exagg * (t1*D[i1+1] + t2*D[i1+2] + t3*D[i1+3])
	u2 = exagg * (t4*D[i1+1] + t5*D[i1+2] + t6*D[i1+3])
	u3 = exagg * (t7*D[i1+1] + t8*D[i1+2] + t9*D[i1+3])
	u5 = exagg * (t4*D[i1+4] + t5*D[i1+5] + t6*D[i1+6])
	u6 = exagg * (t7*D[i1+4] + t8*D[i1+5] + t9*D[i1+6])
	u7 = exagg * (t1*D[i2+1] + t2*D[i2+2] + t3*D[i2+3])
	u8 = exagg * (t4*D[i2+1] + t5*D[i2+2] + t6*D[i2+3])
	u9 = exagg * (t7*D[i2+1] + t8*D[i2+2] + t9*D[i2+3])
	u11 = exagg * (t4*D[i2+4] + t5*D[i2+5] + t6*D[i2+6])
	u12 = exagg * (t7*D[i2+4] + t8*D[i2+5] + t9*D[i2+6])
	a[1] = u2
	b[1] = u3
	a[2] = u8
	b[2] = u9
	a[3] = u6
	b[3] = -u5
	a[4] = u12
	b[4] = -u11
	u7 += L
	A[1][1] = 1
	A[1][2] = u1
	A[1][3] = u1 * u1
	A[1][4] = u1 * u1 * u1
	A[2][1] = 1
	A[2][2] = u7
	A[2][3] = u7 * u7
	A[2][4] = u7 * u7 * u7
	A[3][1] = 0
	A[3][2] = 1
	A[3][3] = 2 * u1
	A[3][4] = 3 * u1 * u1
	A[4][1] = 0
	A[4][2] = 1
	A[4][3] = 2 * u7
	A[4][4] = 3 * u7 * u7
	u7 -= L
	lu_dcmp(A, 4, a, 1, 1, (*[100000000]int)(unsafe.Pointer(&pd))[:])
	if noarch.NotInt(pd) != 0 {
		noarch.Sprintf(errMsg, []byte(" n1 = %d  n2 = %d  L = %e  u7 = %e \n\x00"), n1, n2, L, u7)
		errorMsg(errMsg)
		os.Exit(30)
	}
	lu_dcmp(A, 4, b, 0, 1, (*[100000000]int)(unsafe.Pointer(&pd))[:])
	for s = u1; math.Abs(s) <= 1.01*math.Abs(L+u7); s += math.Abs(L+u7-u1) / 10 {
		v = a[1] + a[2]*s + a[3]*s*s + a[4]*s*s*s
		w = b[1] + b[2]*s + b[3]*s*s + b[4]*s*s*s
		dX = t1*s + t4*v + t7*w
		dY = t2*s + t5*v + t8*w
		dZ = t3*s + t6*v + t9*w
		noarch.Fprintf(fpm, []byte(" %12.4e %12.4e %12.4e\n\x00"), xyz[n1].x+dX, xyz[n1].y+dY, xyz[n1].z+dZ)
	}
	noarch.Fprintf(fpm, []byte("\n\n\x00"))
	free_dmatrix(A, 1, 4, 1, 4)
	free_dvector(a, 1, 4)
	free_dvector(b, 1, 4)
}

// force_bent_beam - transpiled function from  $GOPATH/src/github.com/Konstantin8105/History_frame3DD/src/frame3dd_io.c:3914
/*
 * FORCE_BENT_BEAM  -  reads internal frame element forces and deflections
 * from the internal force and deflection data file.
 * Saves deflected shapes to a file.  These bent shapes are exact.
 * Note: It would not be difficult to adapt this function to plot
 * internal axial force, shear force, torques, or bending moments.
 * 9 Jan 2010
 */ //
/* coord xfmn	*/ //
/* *scanf return value	*/ //
/* read the deformed shape in local coordinates */ //
//		    printf("x = %12.4f\n", x );		/* debug */
/* exaggerated deformed shape in global coordinates */ //
//		printf("...  x = %7.3f  n = %3d  Dx = %10.3e   Dy = %10.3e   Dz = %10.3e \n", x,n,Dx,Dy,Dz ); /* debug */
//		printf("                           dX = %10.3e   dY = %10.3e   dZ = %10.3e \n", dX,dY,dZ ); /* debug */
//
func force_bent_beam(fpm *noarch.File, fpif *noarch.File, fnif []byte, nx int, n1 int, n2 int, xyz []vec3, L float64, p float32, D []float64, exagg float64) {
	var t1 float64
	var t2 float64
	var t3 float64
	var t4 float64
	var t5 float64
	var t6 float64
	var t7 float64
	var t8 float64
	var t9 float64
	var xi float64
	var dX float64
	var dY float64
	var dZ float64
	var x float32
	var Nx float32
	var Vy float32
	var Vz float32
	var Tx float32
	var My float32
	var Mz float32
	var Dx float32
	var Dy float32
	var Dz float32
	var Rx float32
	var Lx float64
	var Ly float64
	var Lz float64
	var n int
	var sfrv int
	Lx = xyz[n2].x - xyz[n1].x
	Ly = xyz[n2].y - xyz[n1].y
	Lz = xyz[n2].z - xyz[n1].z
	coord_trans(xyz, L, n1, n2, (*[100000000]float64)(unsafe.Pointer(&t1))[:], (*[100000000]float64)(unsafe.Pointer(&t2))[:], (*[100000000]float64)(unsafe.Pointer(&t3))[:], (*[100000000]float64)(unsafe.Pointer(&t4))[:], (*[100000000]float64)(unsafe.Pointer(&t5))[:], (*[100000000]float64)(unsafe.Pointer(&t6))[:], (*[100000000]float64)(unsafe.Pointer(&t7))[:], (*[100000000]float64)(unsafe.Pointer(&t8))[:], (*[100000000]float64)(unsafe.Pointer(&t9))[:], p)
	x = float32(-1)
	n = 0
	for xi = 0; xi <= 1.01*L && n < nx; xi += 0.1 * L {
		for float64(x) < xi && n < nx {
			sfrv = noarch.Fscanf(fpif, []byte("%f %f %f %f %f %f %f %f %f %f %f\x00"), (*[100000000]float32)(unsafe.Pointer(&x))[:], (*[100000000]float32)(unsafe.Pointer(&Nx))[:], (*[100000000]float32)(unsafe.Pointer(&Vy))[:], (*[100000000]float32)(unsafe.Pointer(&Vz))[:], (*[100000000]float32)(unsafe.Pointer(&Tx))[:], (*[100000000]float32)(unsafe.Pointer(&My))[:], (*[100000000]float32)(unsafe.Pointer(&Mz))[:], (*[100000000]float32)(unsafe.Pointer(&Dx))[:], (*[100000000]float32)(unsafe.Pointer(&Dy))[:], (*[100000000]float32)(unsafe.Pointer(&Dz))[:], (*[100000000]float32)(unsafe.Pointer(&Rx))[:])
			if sfrv != 11 {
				sferr(fnif)
			}
			n++
		}
		dX = exagg * (t1*float64(Dx) + t4*float64(Dy) + t7*float64(Dz))
		dY = exagg * (t2*float64(Dx) + t5*float64(Dy) + t8*float64(Dz))
		dZ = exagg * (t3*float64(Dx) + t6*float64(Dy) + t9*float64(Dz))
		noarch.Fprintf(fpm, []byte(" %12.4e %12.4e %12.4e\n\x00"), xyz[n1].x+float64(x)/L*Lx+dX, xyz[n1].y+float64(x)/L*Ly+dY, xyz[n1].z+float64(x)/L*Lz+dZ)
	}
	noarch.Fprintf(fpm, []byte("\n\n\x00"))
}

// my_itoa - transpiled function from  $GOPATH/src/github.com/Konstantin8105/History_frame3DD/src/frame3dd_io.c:3982
/*
 * SFERR  -  Display error message upon an erronous *scanf operation
 *
void sferr ( char s[] ) {
	char	errMsg[MAXL];
	sprintf(errMsg,">> Input Data file error while reading %s\n",s);
	errorMsg(errMsg);
	return;
}
*/ //
/*
 * MY_ITOA  -  Convert an integer n to charcters in s, from K&R, 1978,   p. 59-60
 * ... specialized for portability between GNU GCC and DJGPP GCC
 */ //
/* record sign */ //
/* make n positive */ //
/* generate digits in reverse order */ //
/* get next digit */ //
/* delete it */ //
/* add leading '0' */ //
/* reverse order of string s */ //
/* j is length of s - 1 */ //
//
func my_itoa(n int, s []byte, k int) {
	var c int
	var i int
	var j int
	var sign int
	if (func() int {
		sign = n
		return sign
	}()) < 0 {
		n = -n
	}
	i = 0
	for {
		s[func() int {
			defer func() {
				i++
			}()
			return i
		}()] = byte(n%10 + int('0'))
		if !(func() int {
			n /= 10
			defer func() {
			}()
			return n
		}() > 0) {
			break
		}
	}
	for i < k {
		s[func() int {
			defer func() {
				i++
			}()
			return i
		}()] = '0'
	}
	if sign < 0 {
		s[func() int {
			defer func() {
				i++
			}()
			return i
		}()] = '-'
	}
	s[i] = '\x00'
	j = 0
	for int(s[j]) != int('\x00') {
		j++
	}
	j--
	for i = 0; i < j; {
		c = int(s[i])
		s[i] = s[j]
		s[j] = byte(c)
		i++
		j--
	}
}

// get_file_ext - transpiled function from  $GOPATH/src/github.com/Konstantin8105/History_frame3DD/src/frame3dd_io.c:4015
/*
 * GET_FILE_EXT  -  get the file extension,
 *		return 1 if the extension is ".csv"
 *		return 2 if the extension is ".fmm"
 *		return 0 otherwise
 */ //
/* the length of file filename */ //
/* the last '.' in filename */ //
/* debugging ... check file names
printf(" filename '%s' has length %d and extension = '%s' \n",
						filename, len, ext);
printf(" Is .CSV? ... = %d \n", !strcmp(ext,".csv") );
*/ //
//
func get_file_ext(filename []byte, ext []byte) int {
	noarch.Printf([]byte("0001\x00"))
	var i int
	var full_len int
	var len int
	noarch.Printf([]byte("0002\x00"))
	for int(filename[func() int {
		defer func() {
			len++
		}()
		return len
	}()]) != int('\x00') {
	}
	full_len = len
	noarch.Printf([]byte("0003\x00"))
	for int(filename[func() int {
		defer func() {
			len--
		}()
		return len
	}()]) != int('.') && len > 0 {
	}
	if len == 0 {
		len = full_len
	}
	noarch.Printf([]byte("0004\x00"))
	len++
	noarch.Printf([]byte("0005\x00"))
	for i = 0; len < full_len; {
		ext[i] = byte(linux.ToLower(int(filename[len])))
		i++
		len++
	}
	noarch.Printf([]byte("0006\x00"))
	if noarch.NotInt(noarch.Strcmp(ext, []byte(".csv\x00"))) != 0 {
		return (1)
	}
	noarch.Printf([]byte("0007\x00"))
	if noarch.NotInt(noarch.Strcmp(ext, []byte(".fmm\x00"))) != 0 {
		return (2)
	}
	noarch.Printf([]byte("0008\n\x00"))
	return (0)
}

// dots - transpiled function from  $GOPATH/src/github.com/Konstantin8105/History_frame3DD/src/frame3dd_io.c:4050
/*
 * DOTS  -  print a set of dots (periods)
 */ //
//
func dots(fp *noarch.File, n int) {
	var i int
	for i = 1; i <= n; i++ {
		noarch.Fprintf(fp, []byte(".\x00"))
	}
}

// evaluate - transpiled function from  $GOPATH/src/github.com/Konstantin8105/History_frame3DD/src/frame3dd_io.c:4059
/*
 * EVALUATE -  displays a randomly-generated goodbye message.
 */ //
//
func evaluate(error float32, rms_resid float32, tol float32, geom int) {
	var r int
	r = rand.Int() % 10
	color(0)
	noarch.Fprintf(stdout, []byte("  RMS relative equilibrium error  = %9.3e \x00"), float64(error))
	if error < tol {
		noarch.Fprintf(stdout, []byte(" < tol = %7.1e \x00"), float64(tol))
		_ = noarch.Fflush(stdout)
		textColor('y', 'b', 'b', 'x')
		noarch.Fprintf(stdout, []byte(" ** converged ** \x00"))
	}
	if error > tol {
		noarch.Fprintf(stdout, []byte(" > tol = %7.1e \x00"), float64(tol))
		_ = noarch.Fflush(stdout)
		textColor('y', 'r', 'b', 'x')
		noarch.Fprintf(stdout, []byte(" !! not converged !! \x00"))
	}
	_ = noarch.Fflush(stdout)
	color(0)
	noarch.Fprintf(stdout, []byte("\n\x00"))
	noarch.Fprintf(stdout, []byte("  RMS residual incremental displ. = %9.3e \x00"), float64(rms_resid))
	dots(stdout, 17)
	_ = noarch.Fflush(stdout)
	if float64(rms_resid) < 1e-24 {
		textColor('y', 'b', 'b', 'x')
		switch r {
		case 0:
			noarch.Fprintf(stdout, []byte(" * brilliant!  * \x00"))
		case 1:
			noarch.Fprintf(stdout, []byte(" *  chuffed!   * \x00"))
		case 2:
			noarch.Fprintf(stdout, []byte(" *  woo-hoo!   * \x00"))
		case 3:
			noarch.Fprintf(stdout, []byte(" *  wicked!    * \x00"))
		case 4:
			noarch.Fprintf(stdout, []byte(" *   beaut!    * \x00"))
		case 5:
			noarch.Fprintf(stdout, []byte(" *   flash!    * \x00"))
		case 6:
			noarch.Fprintf(stdout, []byte(" *  well done! * \x00"))
		case 7:
			noarch.Fprintf(stdout, []byte(" *  priceless! * \x00"))
		case 8:
			noarch.Fprintf(stdout, []byte(" *  sweet as!  * \x00"))
		case 9:
			{
				noarch.Fprintf(stdout, []byte(" *good as gold!* \x00"))
				break
			}
		}
		_ = noarch.Fflush(stdout)
		color(0)
		noarch.Fprintf(stdout, []byte("\n\x00"))
		return
	}
	if float64(rms_resid) < 1e-16 {
		textColor('y', 'g', 'b', 'x')
		switch r {
		case 0:
			noarch.Fprintf(stdout, []byte("   acceptable!   \x00"))
		case 1:
			noarch.Fprintf(stdout, []byte("      bling!     \x00"))
		case 2:
			noarch.Fprintf(stdout, []byte("  that will do!  \x00"))
		case 3:
			noarch.Fprintf(stdout, []byte("   not shabby!   \x00"))
		case 4:
			noarch.Fprintf(stdout, []byte("   reasonable!   \x00"))
		case 5:
			noarch.Fprintf(stdout, []byte("   very good!    \x00"))
		case 6:
			noarch.Fprintf(stdout, []byte("   up to snuff!  \x00"))
		case 7:
			noarch.Fprintf(stdout, []byte("     bully!      \x00"))
		case 8:
			noarch.Fprintf(stdout, []byte("      nice!      \x00"))
		case 9:
			{
				noarch.Fprintf(stdout, []byte("     choice!     \x00"))
				break
			}
		}
		_ = noarch.Fflush(stdout)
		color(0)
		noarch.Fprintf(stdout, []byte("\n\x00"))
		return
	}
	if float64(rms_resid) < 1e-12 {
		textColor('y', 'c', 'b', 'x')
		switch r {
		case 0:
			noarch.Fprintf(stdout, []byte(" adequate. \x00"))
		case 1:
			noarch.Fprintf(stdout, []byte(" passable. \x00"))
		case 2:
			noarch.Fprintf(stdout, []byte(" all right. \x00"))
		case 3:
			noarch.Fprintf(stdout, []byte(" ok. \x00"))
		case 4:
			noarch.Fprintf(stdout, []byte(" not bad. \x00"))
		case 5:
			noarch.Fprintf(stdout, []byte(" fine. \x00"))
		case 6:
			noarch.Fprintf(stdout, []byte(" fair. \x00"))
		case 7:
			noarch.Fprintf(stdout, []byte(" respectable. \x00"))
		case 8:
			noarch.Fprintf(stdout, []byte(" tolerable. \x00"))
		case 9:
			{
				noarch.Fprintf(stdout, []byte(" just ok. \x00"))
				break
			}
		}
		_ = noarch.Fflush(stdout)
		color(0)
		noarch.Fprintf(stdout, []byte("\n\x00"))
		return
	}
	if float64(rms_resid) > 1e-12 {
		textColor('y', 'r', 'b', 'x')
		switch r {
		case 0:
			noarch.Fprintf(stdout, []byte(" abominable! \x00"))
		case 1:
			noarch.Fprintf(stdout, []byte(" puckeroo! \x00"))
		case 2:
			noarch.Fprintf(stdout, []byte(" atrocious! \x00"))
		case 3:
			noarch.Fprintf(stdout, []byte(" not ok! \x00"))
		case 4:
			noarch.Fprintf(stdout, []byte(" wonky! \x00"))
		case 5:
			noarch.Fprintf(stdout, []byte(" crappy! \x00"))
		case 6:
			noarch.Fprintf(stdout, []byte(" oh noooo! \x00"))
		case 7:
			noarch.Fprintf(stdout, []byte(" abominable! \x00"))
		case 8:
			noarch.Fprintf(stdout, []byte(" munted! \x00"))
		case 9:
			{
				noarch.Fprintf(stdout, []byte(" awful! \x00"))
				break
			}
		}
		_ = noarch.Fflush(stdout)
		color(0)
		noarch.Fprintf(stdout, []byte("\n\x00"))
		return
	}
}

// coord_trans - transpiled function from  $GOPATH/src/github.com/Konstantin8105/History_frame3DD/src/coordtrans.c:51
/*
 This file is part of FRAME3DD:
 Static and dynamic structural analysis of 2D and 3D frames and trusses with
 elastic and geometric stiffness.
 ---------------------------------------------------------------------------
 http://frame3dd.sourceforge.net/
 ---------------------------------------------------------------------------
 Copyright (C) 1992-2009  Henri P. Gavin

    FRAME3DD is free software: you can redistribute it and/or modify
    it under the terms of the GNU General Public License as published by
    the Free Software Foundation, either version 3 of the License, or
    (at your option) any later version.

    FRAME3DD is distributed in the hope that it will be useful,
    but WITHOUT ANY WARRANTY; without even the implied warranty of
    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
    GNU General Public License for more details.

    You should have received a copy of the GNU General Public License
    along with FRAME3DD.  If not, see <http://www.gnu.org/licenses/>.
*/ //
/* -------------------------------------------------------------------------
COORD_TRANS - calculate the 9 elements of the block-diagonal 12-by-12
coordinate transformation matrix, t1, t2, ..., t9.

These coordinate transformation factors are used to:
* transform frame element end forces from the element (local) coordinate system
to the structral (global) coordinate system.
* transfrom end displacements from the structural (global) coordinate system
to the element (local) coordinate system,
* transform the frame element stiffness and mass matrices
from element (local) coordinates to structral (global) coordinates.

Element matrix coordinate transformations are carried out by function ATMA
in frame3dd.c

Currently coordinate transformations do not consider the effect of
finite node sizes ... this needs work, and could require a substantial
re-write of much of the code.

Currently the effect of finite node sizes is used only in the calculation
of the element stiffness matrices.
------------------------------------------------------------------------- */ //
/**< the roll angle (radians) */ //
/* direction cosines	*/ //
/* cosine and sine of roll angle */ //
//
func coord_trans(xyz []vec3, L float64, n1 int, n2 int, t1 []float64, t2 []float64, t3 []float64, t4 []float64, t5 []float64, t6 []float64, t7 []float64, t8 []float64, t9 []float64, p float32) {
	var Cx float64
	var Cy float64
	var Cz float64
	var den float64
	var Cp float64
	var Sp float64
	Cx = (xyz[n2].x - xyz[n1].x) / L
	Cy = (xyz[n2].y - xyz[n1].y) / L
	Cz = (xyz[n2].z - xyz[n1].z) / L
	t9[0] = 0
	t8[0] = t9[0]
	t7[0] = t8[0]
	t6[0] = t7[0]
	t5[0] = t6[0]
	t4[0] = t5[0]
	t3[0] = t4[0]
	t2[0] = t3[0]
	t1[0] = t2[0]
	Cp = math.Cos(float64(p))
	Sp = math.Sin(float64(p))
	if math.Abs(Cz) == 1 {
		t3[0] = Cz
		t4[0] = -Cz * Sp
		t5[0] = Cp
		t7[0] = -Cz * Cp
		t8[0] = -Sp
	} else {
		den = math.Sqrt(1 - Cz*Cz)
		t1[0] = Cx
		t2[0] = Cy
		t3[0] = Cz
		t4[0] = (-Cx*Cz*Sp - Cy*Cp) / den
		t5[0] = (-Cy*Cz*Sp + Cx*Cp) / den
		t6[0] = Sp * den
		t7[0] = (-Cx*Cz*Cp + Cy*Sp) / den
		t8[0] = (-Cy*Cz*Cp - Cx*Sp) / den
		t9[0] = Cp * den
	}
}

// atma - transpiled function from  $GOPATH/src/github.com/Konstantin8105/History_frame3DD/src/coordtrans.c:131
/* ------------------------------------------------------------------------------
 * ATMA  -  perform the coordinate transformation from local to global     6jan96
 *	  include effects of a finite node radii, r1 and r2.	    9dec04
 *	  ------------------------------------------------------------------------------*/ //
/*  effect of finite node radius on coordinate transformation  ... */ //
/*  this needs work ... */ //
/*
 	a[5][1] =  r1*t7;
  	a[5][2] =  r1*t8;
  	a[5][3] =  r1*t9;
  	a[6][1] = -r1*t4;
  	a[6][2] = -r1*t5;
  	a[6][3] = -r1*t6;

  	a[11][7] = -r2*t7;
  	a[11][8] = -r2*t8;
  	a[11][9] = -r2*t9;
  	a[12][7] =  r2*t4;
  	a[12][8] =  r2*t5;
  	a[12][9] =  r2*t6;
*/ //
/*  MT = M T     */ //
/*  T'MT = T' MT */ //
//
func atma(t1 float64, t2 float64, t3 float64, t4 float64, t5 float64, t6 float64, t7 float64, t8 float64, t9 float64, m [][]float64, r1 float32, r2 float32) {
	var a [][]float64
	var ma [][]float64
	var i int
	var j int
	var k int
	a = dmatrix(1, 12, 1, 12)
	ma = dmatrix(1, 12, 1, 12)
	for i = 1; i <= 12; i++ {
		for j = i; j <= 12; j++ {
			a[i][j] = 0
			a[j][i] = a[i][j]
			ma[i][j] = a[j][i]
			ma[j][i] = ma[i][j]
		}
	}
	for i = 0; i <= 3; i++ {
		a[3*i+1][3*i+1] = t1
		a[3*i+1][3*i+2] = t2
		a[3*i+1][3*i+3] = t3
		a[3*i+2][3*i+1] = t4
		a[3*i+2][3*i+2] = t5
		a[3*i+2][3*i+3] = t6
		a[3*i+3][3*i+1] = t7
		a[3*i+3][3*i+2] = t8
		a[3*i+3][3*i+3] = t9
	}
	for j = 1; j <= 12; j++ {
		for i = 1; i <= 12; i++ {
			for k = 1; k <= 12; k++ {
				ma[i][j] += m[i][k] * a[k][j]
			}
		}
	}
	for i = 1; i <= 12; i++ {
		for j = i; j <= 12; j++ {
			m[i][j] = 0
			m[j][i] = m[i][j]
		}
	}
	for j = 1; j <= 12; j++ {
		for i = 1; i <= 12; i++ {
			for k = 1; k <= 12; k++ {
				m[i][j] += a[k][i] * ma[k][j]
			}
		}
	}
	free_dmatrix(a, 1, 12, 1, 12)
	free_dmatrix(ma, 1, 12, 1, 12)
}

// subspace - transpiled function from  $GOPATH/src/github.com/Konstantin8105/History_frame3DD/src/eig.c:66
/*
 This file is part of FRAME3DD:
 Static and dynamic structural analysis of 2D and 3D frames and trusses with
 elastic and geometric stiffness.
 ---------------------------------------------------------------------------
 http://frame3dd.sourceforge.net/
 ---------------------------------------------------------------------------
 Copyright (C) 1992-2009  Henri P. Gavin

    FRAME3DD is free software: you can redistribute it and/or modify
    it under the terms of the GNU General Public License as published by
    the Free Software Foundation, either version 3 of the License, or
    (at your option) any later version.

    FRAME3DD is distributed in the hope that it will be useful,
    but WITHOUT ANY WARRANTY; without even the implied warranty of
    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
    GNU General Public License for more details.

    You should have received a copy of the GNU General Public License
    along with FRAME3DD.  If not, see <http://www.gnu.org/licenses/>.
*/ //
/**
	@file
	Routines to solve the generalized eigenvalue problem

 Henri P. Gavin                                             hpgavin@duke.edu
 Department of Civil and Environmental Engineering
 Duke University, Box 90287
 Durham, NC  27708--0287
*/ //
/* #define EIG_DEBUG */ //
/* forward declarations */ //
/*-----------------------------------------------------------------------------
SUBSPACE - Find the lowest m eigen-values, w, and eigen-vectors, V, of the
general eigen-problem  ...       K V = w M V using sub-space / Jacobi iteration
where
  K is an n by n  symmetric real (stiffness) matrix
  M is an n by n  symmetric positive definate real (mass) matrix
  w is a diagonal matrix of eigen-values
  V is a  rectangular matrix of eigen-vectors

 H.P. Gavin, Civil Engineering, Duke University, hpgavin@duke.edu  1 March 2007
 Bathe, Finite Element Procecures in Engineering Analysis, Prentice Hall, 1982
-----------------------------------------------------------------------------*/ //
/**< DoF and number of required modes	*/ //
/**< sub-space iterations		*/ //
/**< Sturm check result			*/ //
/* display convergence info.	*/ //
/* shift eigen-values by this much */ //
/* use L D L' decomp  */ //
//	for (k=1; k<=m; k++) printf(" idx[%d] = %d \n", k, idx[k] ); /*debug*/
//		printf(" idx[%3d] = %3d   ok = %d \n", k , idx[k], *ok); /*debug*/
/*	for (i=1; i<=n; i++)	V[i][1] = M[i][i];	// diag(M)	*/ //
/* Begin sub-space iterations */ //
/* K Xb = M V	(12.10) */ //
/* LDL bk-sub */ //
/* improve the solution iteratively */ //
/* Kb = Xb' K Xb (12.11) */ //
/* Mb = Xb' M Xb (12.12) */ //
/* (12.13) */ //
/* V = Xb Qb (12.14) */ //
/* End   sub-space iterations */ //
/* shift eigen-values */ //
//
func subspace(K [][]float64, M [][]float64, n int, m int, w []float64, V [][]float64, tol float64, shift float64, iter []int, ok []int, verbose int) {
	var Kb [][]float64
	var Mb [][]float64
	var Xb [][]float64
	var Qb [][]float64
	var d []float64
	var u []float64
	var v []float64
	var km float64
	var km_old float64
	var error float64 = 1
	var w_old float64
	var i int
	var j int
	var k int
	var modes int
	var disp int
	var idx []int
	var errMsg []byte = make([]byte, 512)
	if m > n {
		noarch.Sprintf(errMsg, []byte("subspace: Number of eigen-values must be less than the problem dimension.\n Desired number of eigen-values=%d \n Dimension of the problem= %d \n\x00"), m, n)
		errorMsg(errMsg)
		os.Exit(32)
	}
	d = dvector(1, int32(n))
	u = dvector(1, int32(n))
	v = dvector(1, int32(n))
	Kb = dmatrix(1, int32(m), 1, int32(m))
	Mb = dmatrix(1, int32(m), 1, int32(m))
	Xb = dmatrix(1, int32(n), 1, int32(m))
	Qb = dmatrix(1, int32(m), 1, int32(m))
	idx = ivector(1, int32(m))
	for i = 1; i <= m; i++ {
		idx[i] = 0
		for j = i; j <= m; j++ {
			Qb[j][i] = 0
			Qb[i][j] = Qb[j][i]
			Mb[j][i] = Qb[i][j]
			Mb[i][j] = Mb[j][i]
			Kb[j][i] = Mb[i][j]
			Kb[i][j] = Kb[j][i]
		}
	}
	for i = 1; i <= n; i++ {
		for j = 1; j <= m; j++ {
			V[i][j] = 0
			Xb[i][j] = V[i][j]
		}
	}
	modes = (func() int {
		if (0.5 * float64(m)) > (float64(m) - 8) {
			return int((float64(m) / 2))
		}
		return m - 8
	}())
	for i = 1; i <= n; i++ {
		for j = i; j <= n; j++ {
			K[i][j] += shift * M[i][j]
		}
	}
	ldl_dcmp(K, n, u, v, v, 1, 0, ok)
	for i = 1; i <= n; i++ {
		if M[i][i] <= 0 {
			noarch.Sprintf(errMsg, []byte(" subspace: M[%d][%d] = %e \n\x00"), i, i, M[i][i])
			errorMsg(errMsg)
			os.Exit(32)
		}
		d[i] = K[i][i] / M[i][i]
	}
	km_old = 0
	for k = 1; k <= m; k++ {
		km = d[1]
		for i = 1; i <= n; i++ {
			if km_old <= d[i] && d[i] <= km {
				ok[0] = 1
				for j = 1; j <= k-1; j++ {
					if i == idx[j] {
						ok[0] = 0
					}
				}
				if ok[0] != 0 {
					km = d[i]
					idx[k] = i
				}
			}
		}
		if idx[k] == 0 {
			i = idx[1]
			for j = 1; j < k; j++ {
				if i < idx[j] {
					i = idx[j]
				}
			}
			idx[k] = i + 1
			km = d[i+1]
		}
		km_old = km
	}
	for k = 1; k <= m; k++ {
		V[idx[k]][k] = 1
		ok[0] = idx[k] % 6
		switch ok[0] {
		case 1:
			i = 1
			j = 2
		case 2:
			i = -1
			j = 1
		case 3:
			i = -1
			j = -2
		case 4:
			i = 1
			j = 2
		case 5:
			i = -1
			j = 1
		case 0:
			{
				i = -1
				j = -2
				break
			}
		}
		V[idx[k]+i][k] = 0.2
		V[idx[k]+j][k] = 0.2
	}
	iter[0] = 0
	for {
		for k = 1; k <= m; k++ {
			prodABj(M, V, v, n, k)
			ldl_dcmp(K, n, u, v, d, 0, 1, ok)
			if disp != 0 {
				noarch.Fprintf(stdout, []byte("  RMS matrix error:\x00"))
			}
			error = float64((func() int {
				ok[0] = 1
				return ok[0]
			}()))
			for {
				ldl_mprove(K, n, u, v, d, (*[100000000]float64)(unsafe.Pointer(&error))[:], ok)
				if disp != 0 {
					noarch.Fprintf(stdout, []byte("%9.2e\x00"), error)
				}
				if noarch.NotInt((ok[0])) != 0 {
					break
				}
			}
			if disp != 0 {
				noarch.Fprintf(stdout, []byte("\n\x00"))
			}
			for i = 1; i <= n; i++ {
				Xb[i][k] = d[i]
			}
		}
		xtAx(K, Xb, Kb, n, m)
		xtAx(M, Xb, Mb, n, m)
		jacobi(Kb, Mb, w, Qb, m)
		prodAB(Xb, Qb, V, n, m, m)
		eigsort(w, V, n, m)
		if w[modes] == 0 {
			noarch.Sprintf(errMsg, []byte(" subspace: Zero frequency found! \n w[%d] = %e \n\x00"), modes, w[modes])
			errorMsg(errMsg)
			os.Exit(32)
		}
		error = math.Abs(w[modes]-w_old) / w[modes]
		iter[0]++
		if disp != 0 {
			noarch.Fprintf(stdout, []byte(" iter = %d  w[%d] = %f error = %e\n\x00"), iter[0], modes, w[modes], error)
		}
		w_old = w[modes]
		if iter[0] > 1000 {
			noarch.Sprintf(errMsg, []byte("  subspace: Iteration limit exceeded\n rel. error = %e > %e\n\x00"), error, tol)
			errorMsg(errMsg)
			os.Exit(32)
		}
		if !(error > tol) {
			break
		}
	}
	for k = 1; k <= m; k++ {
		if w[k] > shift {
			w[k] = w[k] - shift
		} else {
			w[k] = shift - w[k]
		}
	}
	if verbose != 0 {
		noarch.Fprintf(stdout, []byte(" %4d sub-space iterations,   error: %.4e \n\x00"), iter[0], error)
		for k = 1; k <= m; k++ {
			noarch.Fprintf(stdout, []byte("  mode: %2d\tDoF: %5d\t %9.4lf Hz\n\x00"), k, idx[k], math.Sqrt(w[k])/(2*3.141592653589793))
		}
	}
	ok[0] = sturm(K, M, n, m, shift, w[modes]+tol, verbose)
	for i = 1; i <= n; i++ {
		for j = i; j <= n; j++ {
			K[i][j] -= shift * M[i][j]
		}
	}
	free_dmatrix(Kb, 1, int32(m), 1, int32(m))
	free_dmatrix(Mb, 1, int32(m), 1, int32(m))
	free_dmatrix(Xb, 1, int32(n), 1, int32(m))
	free_dmatrix(Qb, 1, int32(m), 1, int32(m))
}

// jacobi - transpiled function from  $GOPATH/src/github.com/Konstantin8105/History_frame3DD/src/eig.c:250
/*-----------------------------------------------------------------------------
 JACOBI - Find all eigen-values, E, and eigen-vectors, V,
 of the general eigen-problem  K V = E M V
 using Jacobi iteration, with efficient matrix rotations.
 K is a symmetric real (stiffness) matrix
 M is a symmetric positive definate real (mass) matrix
 E is a diagonal matrix of eigen-values
 V is a  square  matrix of eigen-vectors

 H.P. Gavin, Civil Engineering, Duke University, hpgavin@duke.edu  1 March 2007
 Bathe, Finite Element Procecures in Engineering Analysis, Prentice Hall, 1982
-----------------------------------------------------------------------------*/ //
/* Begin Sweep Iteration */ //
/* sweep along upper diagonals */ //
/* row */ //
/* column */ //
/* do a rotation */ //
/* make Kij zero */ //
/* make Mij zero */ //
/*  update eigen-vectors  V = V * P */ //
/* rotations complete */ //
/* row */ //
/* diagonal */ //
/* End Sweep Iteration */ //
/* scale eigen-vectors */ //
/* eigen-values */ //
//
func jacobi(K [][]float64, M [][]float64, E []float64, V [][]float64, n int) {
	var iter int
	var d int
	var i int
	var j int
	var k int
	var Kii float64
	var Kjj float64
	var Kij float64
	var Mii float64
	var Mjj float64
	var Mij float64
	var Vki float64
	var Vkj float64
	var alpha float64
	var beta float64
	var gamma float64
	var s float64
	var tol float64
	Vkj = 0
	Vki = Vkj
	Mij = Vki
	Mjj = Mij
	Mii = Mjj
	Kij = Mii
	Kjj = Kij
	Kii = Kjj
	for i = 1; i <= n; i++ {
		for j = i + 1; j <= n; j++ {
			V[j][i] = 0
			V[i][j] = V[j][i]
		}
	}
	for d = 1; d <= n; d++ {
		V[d][d] = 1
	}
	for iter = 1; iter <= 2*n; iter++ {
		tol = math.Pow(0.01, float64((2 * iter)))
		tol = 0
		for d = 1; d <= n-1; d++ {
			for i = 1; i <= n-d; i++ {
				j = i + d
				Kij = K[i][j]
				Mij = M[i][j]
				if Kij*Kij/(K[i][i]*K[j][j]) > tol || Mij*Mij/(M[i][i]*M[j][j]) > tol {
					Kii = K[i][i]*Mij - Kij*M[i][i]
					Kjj = K[j][j]*Mij - Kij*M[j][j]
					s = K[i][i]*M[j][j] - K[j][j]*M[i][i]
					if s >= 0 {
						gamma = 0.5*s + math.Sqrt(0.25*s*s+Kii*Kjj)
					} else {
						gamma = 0.5*s - math.Sqrt(0.25*s*s+Kii*Kjj)
					}
					alpha = Kjj / gamma
					beta = -Kii / gamma
					rotate(K, n, alpha, beta, i, j)
					rotate(M, n, alpha, beta, i, j)
					for k = 1; k <= n; k++ {
						Vki = V[k][i]
						Vkj = V[k][j]
						V[k][i] = Vki + beta*Vkj
						V[k][j] = Vkj + alpha*Vki
					}
				}
			}
		}
	}
	for j = 1; j <= n; j++ {
		Mjj = math.Sqrt(M[j][j])
		for i = 1; i <= n; i++ {
			V[i][j] /= Mjj
		}
	}
	for j = 1; j <= n; j++ {
		E[j] = K[j][j] / M[j][j]
	}
}

// rotate - transpiled function from  $GOPATH/src/github.com/Konstantin8105/History_frame3DD/src/eig.c:319
/*-----------------------------------------------------------------------------
ROTATE - rotate an n by n symmetric matrix A such that A[i][j] = A[j][i] = 0
     A = P' * A * P  where diag(P) = 1 and P[i][j] = alpha and P[j][i] = beta.
     Since P is sparse, this matrix multiplcation can be done efficiently.
-----------------------------------------------------------------------------*/ //
/* elements of A	*/ //
/* i-th and j-th rows of A */ //
//
func rotate(A [][]float64, n int, alpha float64, beta float64, i int, j int) {
	var Aii float64
	var Ajj float64
	var Aij float64
	var Ai []float64
	var Aj []float64
	var k int
	Ai = dvector(1, int32(n))
	Aj = dvector(1, int32(n))
	for k = 1; k <= n; k++ {
		Ai[k] = A[i][k]
		Aj[k] = A[j][k]
	}
	Aii = A[i][i]
	Ajj = A[j][j]
	Aij = A[i][j]
	A[i][i] = Aii + 2*beta*Aij + beta*beta*Ajj
	A[j][j] = Ajj + 2*alpha*Aij + alpha*alpha*Aii
	for k = 1; k <= n; k++ {
		if k != i && k != j {
			A[i][k] = Ai[k] + beta*Aj[k]
			A[k][i] = A[i][k]
			A[j][k] = Aj[k] + alpha*Ai[k]
			A[k][j] = A[j][k]
		}
	}
	A[i][j] = 0
	A[j][i] = A[i][j]
	free_dvector(Ai, 1, int32(n))
	free_dvector(Aj, 1, int32(n))
}

// stodola - transpiled function from  $GOPATH/src/github.com/Konstantin8105/History_frame3DD/src/eig.c:363
/*------------------------------------------------------------------------------
STODOLA  -  calculate the lowest m eigen-values and eigen-vectors of the
generalized eigen-problem, K v = w M v, using a matrix iteration approach
with shifting. 								15oct98

 H.P. Gavin, Civil Engineering, Duke University, hpgavin@duke.edu  12 Jul 2001
------------------------------------------------------------------------------*/ //
/* stiffness and mass matrices */ //
/* DoF and number of required modes	*/ //
/* the dynamics matrix, D = K^(-1) M	*/ //
/* minimum value of D[i][i]		*/ //
/* maximum value of D[i][i]		*/ //
/* previous extreme value of D[i][i]	*/ //
/* columns of the D, M, and V matrices	*/ //
/* trial eigen-vector vectors		*/ //
/* coefficients for lower mode purge	*/ //
/* factor for mass normalization	*/ //
/* Raliegh quotient			*/ //
/* location of minimum value of D[i][i]	*/ //
/* number of desired modes		*/ //
/* 1: display convergence error; 0: dont*/ //
/* shift eigen-values by this much */ //
/* use L D L' decomp	*/ //
/* calculate  D = K^(-1) M */ //
/* L D L' bk-sub */ //
/* improve the solution iteratively */ //
/* loop over lowest m modes */ //
/* initial guess */ //
/* mass-normalize */ //
/* purge lower modes */ //
/* mass-normalize */ //
/* Raleigh quotient */ //
/* iterate	*/ //
/* v = D u	*/ //
/* mass-normalize */ //
/* purge lower modes */ //
/* mass-normalize */ //
/* Raleigh quotient */ //
//
func stodola(K [][]float64, M [][]float64, n int, m int, w []float64, V [][]float64, tol float64, shift float64, iter []int, ok []int, verbose int) {
	var D [][]float64
	var d_min float64
	var d_max float64
	var d_old float64
	var d []float64
	var u []float64
	var v []float64
	var c []float64
	var vMv float64
	var RQ float64
	var RQold float64
	var error float64 = 1
	var i_ex int = 9999
	var modes int
	var disp int
	var i int
	var j int
	var k int
	var errMsg []byte = make([]byte, 512)
	D = dmatrix(1, int32(n), 1, int32(n))
	d = dvector(1, int32(n))
	u = dvector(1, int32(n))
	v = dvector(1, int32(n))
	c = dvector(1, int32(m))
	modes = (func() int {
		if (0.5 * float64(m)) > float64((m - 8)) {
			return int((float64(m) / 2))
		}
		return m - 8
	}())
	for i = 1; i <= n; i++ {
		for j = i; j <= n; j++ {
			K[i][j] += shift * M[i][j]
		}
	}
	ldl_dcmp(K, n, u, v, v, 1, 0, ok)
	if ok[0] < 0 {
		noarch.Sprintf(errMsg, []byte(" Make sure that all six rigid body translation are restrained.\n\x00"))
		errorMsg(errMsg)
		os.Exit(32)
	}
	for j = 1; j <= n; j++ {
		for i = 1; i <= n; i++ {
			v[i] = M[i][j]
		}
		ldl_dcmp(K, n, u, v, d, 0, 1, ok)
		if disp != 0 {
			noarch.Fprintf(stdout, []byte("  RMS matrix error:\x00"))
		}
		error = float64((func() int {
			ok[0] = 1
			return ok[0]
		}()))
		for {
			ldl_mprove(K, n, u, v, d, (*[100000000]float64)(unsafe.Pointer(&error))[:], ok)
			if disp != 0 {
				noarch.Fprintf(stdout, []byte("%9.2e\x00"), error)
			}
			if noarch.NotInt((ok[0])) != 0 {
				break
			}
		}
		if disp != 0 {
			noarch.Fprintf(stdout, []byte("\n\x00"))
		}
		for i = 1; i <= n; i++ {
			D[i][j] = d[i]
		}
	}
	iter[0] = 0
	for i = 1; i <= n; i++ {
		if D[i][i] > d_max {
			d_max = D[i][i]
		}
	}
	d_min = d_max
	d_old = d_min
	for i = 1; i <= n; i++ {
		if D[i][i] < d_min {
			d_min = D[i][i]
		}
	}
	for k = 1; k <= m; k++ {
		d_max = d_min
		for i = 1; i <= n; i++ {
			u[i] = 0
			if D[i][i] < d_old && D[i][i] > d_max {
				d_max = D[i][i]
				i_ex = i
			}
		}
		u[i_ex] = 1
		u[i_ex+1] = 0.0001
		d_old = d_max
		vMv = xtAy(u, M, u, n, d)
		for i = 1; i <= n; i++ {
			u[i] /= math.Sqrt(vMv)
		}
		for j = 1; j < k; j++ {
			for i = 1; i <= n; i++ {
				v[i] = V[i][j]
			}
			c[j] = xtAy(v, M, u, n, d)
		}
		for j = 1; j < k; j++ {
			for i = 1; i <= n; i++ {
				u[i] -= c[j] * V[i][j]
			}
		}
		vMv = xtAy(u, M, u, n, d)
		for i = 1; i <= n; i++ {
			u[i] /= math.Sqrt(vMv)
		}
		RQ = xtAy(u, K, u, n, d)
		for {
			for i = 1; i <= n; i++ {
				v[i] = 0
				for j = 1; j <= n; j++ {
					v[i] += D[i][j] * u[j]
				}
			}
			vMv = xtAy(v, M, v, n, d)
			for i = 1; i <= n; i++ {
				v[i] /= math.Sqrt(vMv)
			}
			for j = 1; j < k; j++ {
				for i = 1; i <= n; i++ {
					u[i] = V[i][j]
				}
				c[j] = xtAy(u, M, v, n, d)
			}
			for j = 1; j < k; j++ {
				for i = 1; i <= n; i++ {
					v[i] -= c[j] * V[i][j]
				}
			}
			vMv = xtAy(v, M, v, n, d)
			for i = 1; i <= n; i++ {
				u[i] = v[i] / math.Sqrt(vMv)
			}
			RQold = RQ
			RQ = xtAy(u, K, u, n, d)
			iter[0]++
			if iter[0] > 1000 {
				noarch.Sprintf(errMsg, []byte("  stodola: Iteration limit exceeded\n  rel. error = %e > %e\n\x00"), (math.Abs(RQ-RQold) / RQ), tol)
				errorMsg(errMsg)
				os.Exit(32)
			}
			if !(math.Abs(RQ-RQold)/RQ > tol) {
				break
			}
		}
		for i = 1; i <= n; i++ {
			V[i][k] = v[i]
		}
		w[k] = xtAy(u, K, u, n, d)
		if w[k] > shift {
			w[k] = w[k] - shift
		} else {
			w[k] = shift - w[k]
		}
		noarch.Fprintf(stdout, []byte("  mode: %2d\tDoF: %5d\t\x00"), k, i_ex)
		noarch.Fprintf(stdout, []byte(" %9.4f Hz\t iter: %4d   error: %.4e \n\x00"), math.Sqrt(w[k])/(2*3.141592653589793), iter[0], (math.Abs(RQ-RQold) / RQ))
	}
	eigsort(w, V, n, m)
	ok[0] = sturm(K, M, n, m, shift, w[modes]+tol, verbose)
	free_dmatrix(D, 1, int32(n), 1, int32(n))
	free_dvector(d, 1, int32(n))
	free_dvector(u, 1, int32(n))
	free_dvector(v, 1, int32(n))
	free_dvector(c, 1, int32(m))
}

// eigsort - transpiled function from  $GOPATH/src/github.com/Konstantin8105/History_frame3DD/src/eig.c:524
/*------------------------------------------------------------------------------
EIGSORT  -  Given the eigenvallues e[1..m] and eigenvectors v[1..n][1..m],
this routine sorts the eigenvalues into ascending order, and rearranges
the columns of v correspondingly.  The method is straight insertion.
Adapted from Numerical Recipes in C, Ch 11
------------------------------------------------------------------------------*/ //
/* find smallest eigen-value */ //
/* swap eigen-values	*/ //
/* swap eigen-vectors	*/ //
//
func eigsort(e []float64, v [][]float64, n int, m int) {
	var k int
	var j int
	var i int
	var p float64
	for i = 1; i < m; i++ {
		k = i
		p = e[k]
		for j = i + 1; j <= m; j++ {
			if e[j] <= p {
				p = e[(func() int {
					k = j
					return k
				}())]
			}
		}
		if k != i {
			e[k] = e[i]
			e[i] = p
			for j = 1; j <= n; j++ {
				p = v[j][i]
				v[j][i] = v[j][k]
				v[j][k] = p
			}
		}
	}
}

// sturm - transpiled function from  $GOPATH/src/github.com/Konstantin8105/History_frame3DD/src/eig.c:563
/*-----------------------------------------------------------------------------
STURM  -  Determine the number of eigenvalues, w, of the general eigen-problem
  K V = w M V which are below the value ws,
  K is an n by n  symmetric real (stiffness) matrix
  M is an n by n  symmetric positive definate real (mass) matrix
  w is a diagonal matrix of eigen-values
  ws is the limit
  n is the number of DoF
  m is the number of required modes


 H.P. Gavin, Civil Engineering, Duke University, hpgavin@duke.edu  30 Aug 2001
 Bathe, Finite Element Procecures in Engineering Analysis, Prentice Hall, 1982
-----------------------------------------------------------------------------*/ //
/* shift [K]	*/ //
//
func sturm(K [][]float64, M [][]float64, n int, m int, shift float64, ws float64, verbose int) int {
	var ws_shift float64
	var d []float64
	var ok int
	var i int
	var j int
	var modes int
	d = dvector(1, int32(n))
	modes = (func() int {
		if float32((0.5 * float64(m))) > float32((float64(m) - 8)) {
			return int((float64(m) / 2))
		}
		return m - 8
	}())
	ws_shift = ws + shift
	for i = 1; i <= n; i++ {
		for j = i; j <= n; j++ {
			K[i][j] -= ws_shift * M[i][j]
		}
	}
	ldl_dcmp(K, n, d, d, d, 1, 0, (*[100000000]int)(unsafe.Pointer(&ok))[:])
	if verbose != 0 {
		noarch.Fprintf(stdout, []byte("  There are %d modes below %f Hz.\x00"), -ok, math.Sqrt(ws)/(2*3.141592653589793))
	}
	if -ok > modes {
		noarch.Fprintf(stderr, []byte(" ... %d modes were not found.\n\x00"), -ok-modes)
		noarch.Fprintf(stderr, []byte(" Try increasing the number of modes in \n\x00"))
		noarch.Fprintf(stderr, []byte(" order to get the missing modes below %f Hz.\n\x00"), math.Sqrt(ws)/(2*3.141592653589793))
	} else if verbose != 0 {
		noarch.Fprintf(stdout, []byte("  All %d modes were found.\n\x00"), modes)
	}
	for i = 1; i <= n; i++ {
		for j = i; j <= n; j++ {
			K[i][j] += ws_shift * M[i][j]
		}
	}
	free_dvector(d, 1, int32(n))
	return ok
}

// check_non_negative - transpiled function from  $GOPATH/src/github.com/Konstantin8105/History_frame3DD/src/eig.c:601
/*----------------------------------------------------------------------------
CHECK_NON_NEGATIVE -  checks that a value is non-negative
-----------------------------------------------------------------------------*/ //
//
func check_non_negative(x float64, i int) {
	if x <= 1e-100 {
		noarch.Fprintf(stderr, []byte(" value %e is less than or equal to zero \x00"), x)
		noarch.Fprintf(stderr, []byte(" i = %d \n\x00"), i)
	} else {
		return
	}
}

// gaussj - transpiled function from  $GOPATH/src/github.com/Konstantin8105/History_frame3DD/src/HPGmatrix.c:55
/*
 * ==========================================================================
 *
 *       Filename:  HPGmatrix.c
 *
 *    Description:  Matrix math functions
 *
 *	Version:  1.0
 *	Created:  12/30/11 18:07:41
 *       Revision:  none
 *       Compiler:  gcc
 *
 *	 Author:  Henri P. Gavin (hpgavin), h p gavin ~at~ duke ~dot~ e d v
 *	Company:  Duke Univ.
 *
 * ==========================================================================
 */ //
/*
 Copyright (C) 2012 Henri P. Gavin

    HPGmatrix is free software: you can redistribute it and/or modify
    it under the terms of the GNU General Public License as published by
    the Free Software Foundation, either version 3 of the License, or
    (at your option) any later version.

    HPGmatrix is distributed in the hope that it will be useful,
    but WITHOUT ANY WARRANTY; without even the implied warranty of
    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
    GNU General Public License for more details.

    You should have received a copy of the GNU General Public License
    along with HPGmatrix.  If not, see <http://www.gnu.org/licenses/>.
*/ //
/*
 * GAUSSJ
 * Linear equation solution by Gauss-Jordan elimination, [A][X]=[B] above. A[1..n][1..n]
 * is the input matrix. B[1..n][1..m] is input containing the m right-hand side vectors. On
 * output, a is replaced by its matrix inverse, and B is replaced by the corresponding set of solution
 * vectors.
 */ //
/* The integer arrays ipiv, indxr, and indxc are used for bookkeeping on the pivoting. */ //
/*  This is the main loop over the columns to be reduced. */ //
/*  This is the outer loop for the search for a pivot element. */ //
/* We now have the pivot element, so we interchange rows, if needed, to put the pivot
 * element on the diagonal. The columns are not physically interchanged, only relabeled:
 * indxc[i], the column of the ith pivot element, is the ith column that is reduced, while
 * indxr[i] is the row in which that pivot element was originally located. If indxr[i] =
 * indxc[i] there is an implied column interchange. With this form of bookkeeping, the
 * solution b's will end up in the correct order, and the inverse matrix will be scrambled
 * by columns.
 */ //
/* We are now ready to divide the pivot row by the by the pivot element, located at irow,icol */ //
/*  Next, we reduce the rows ... except for the pivot one, of course. */ //
/* This is the end of the main loop over columns of the reduction. It only remains to unscram-
 * ble the solution in view of the column interchanges. We do this by interchanging pairs of
 * columns in the reverse order that the permutation was built up.
 */ //
/*  And we are done. */ //
//
func gaussj(A [][]float32, n int, B [][]float32, m int) {
	var indxc []int
	var indxr []int
	var ipiv []int
	var i int
	var icol int = 1
	var irow int = 1
	var j int
	var k int
	var l int
	var ll int
	var big float32
	var dum float32
	var pivinv float32
	var temp float32
	indxc = ivector(1, int32(n))
	indxr = ivector(1, int32(n))
	ipiv = ivector(1, int32(n))
	for j = 1; j <= n; j++ {
		ipiv[j] = 0
	}
	for i = 1; i <= n; i++ {
		big = float32(0)
		for j = 1; j <= n; j++ {
			if ipiv[j] != 1 {
				for k = 1; k <= n; k++ {
					if ipiv[k] == 0 {
						if math.Abs(float64(A[j][k])) >= float64(big) {
							big = float32(math.Abs(float64(A[j][k])))
							irow = j
							icol = k
						}
					} else if ipiv[k] > 1 {
						NRerror([]byte("gaussj: Singular Matrix-1\x00"))
					}
				}
			}
		}
		ipiv[icol]++
		if irow != icol {
			for l = 1; l <= n; l++ {
				temp = A[irow][l]
				A[irow][l] = A[icol][l]
				A[icol][l] = temp
			}
			for l = 1; l <= m; l++ {
				temp = B[irow][l]
				B[irow][l] = B[icol][l]
				B[icol][l] = temp
			}
		}
		indxr[i] = irow
		indxc[i] = icol
		if float64(A[icol][icol]) == 0 {
			NRerror([]byte("gaussj: Singular Matrix-2\x00"))
		}
		pivinv = float32(1 / float64(A[icol][icol]))
		A[icol][icol] = float32(1)
		for l = 1; l <= n; l++ {
			A[icol][l] *= pivinv
		}
		for l = 1; l <= m; l++ {
			B[icol][l] *= pivinv
		}
		for ll = 1; ll <= n; ll++ {
			if ll != icol {
				dum = A[ll][icol]
				A[ll][icol] = float32(0)
				for l = 1; l <= n; l++ {
					A[ll][l] -= A[icol][l] * dum
				}
				for l = 1; l <= m; l++ {
					B[ll][l] -= B[icol][l] * dum
				}
			}
		}
	}
	for l = n; l >= 1; l-- {
		if indxr[l] != indxc[l] {
			for k = 1; k <= n; k++ {
				temp = A[k][indxr[l]]
				A[k][indxr[l]] = A[k][indxc[l]]
				A[k][indxc[l]] = temp
			}
		}
	}
	free_ivector(ipiv, 1, int32(n))
	free_ivector(indxr, 1, int32(n))
	free_ivector(indxc, 1, int32(n))
}

// lu_dcmp - transpiled function from  $GOPATH/src/github.com/Konstantin8105/History_frame3DD/src/HPGmatrix.c:154
/*
 * LU_DCMP
 * Solves [A]{x} = {b} simply and efficiently by performing an
 * LU - decomposition of [A].  No pivoting is performed.
 * [A] is a diagonally dominant matrix of dimension [1..n][1..n].
 * {b} is a r.h.s. vector of dimension [1..n].
 * {b} is updated using [LU] and then back-substitution is done to obtain {x}.
 * {b} is replaced by {x} and [A] is replaced by the LU - reduction of itself.
 *
 *  usage:  double **A, *b;
 *	  int   n, reduce, solve, pd;
 *	  lu_dcmp ( A, n, b, reduce, solve, &pd );		     5may98
 */ //
// the system matrix, and it's LU- reduction
// the dimension of the matrix
// the right hand side vector, and the solution vector
// 1: do a forward reduction; 0: don't do the reduction
// 1: do a back substitution for {x};  0: do no bk-sub'n
// 1: positive diagonal  and  successful LU decomp'n
/* a diagonal element of [A]		*/ //
/* forward reduction of [A]	*/ //
/* the forward reduction of [A] is now complete	*/ //
/* back substitution to solve for {x}	*/ //
/* {b} is run through the same forward reduction as was [A]	*/ //
/* now back substitution is conducted on {b};  [A] is preserved */ //
/* finally we solve for the {x} vector			*/ //
/* TAA DAAAAAAAA! {b} is now {x} and is ready to be returned	*/ //
//
func lu_dcmp(A [][]float64, n int, b []float64, reduce int, solve int, pd []int) {
	var pivot float64
	var i int
	var j int
	var k int
	pd[0] = 1
	if reduce != 0 {
		for k = 1; k <= n; k++ {
			if 0 == (func() float64 {
				pivot = A[k][k]
				return pivot
			}()) {
				noarch.Fprintf(stderr, []byte(" lu_dcmp: zero found on the diagonal\n\x00"))
				noarch.Fprintf(stderr, []byte(" A[%d][%d] = %11.4e\n\x00"), k, k, A[k][k])
				pd[0] = 0
				return
			}
			for i = k + 1; i <= n; i++ {
				A[i][k] /= pivot
				for j = k + 1; j <= n; j++ {
					A[i][j] -= A[i][k] * A[k][j]
				}
			}
		}
	}
	if solve != 0 {
		for k = 1; k <= n; k++ {
			for i = k + 1; i <= n; i++ {
				b[i] -= A[i][k] * b[k]
			}
		}
		for j = n; j >= 2; j-- {
			for i = 1; i <= j-1; i++ {
				b[i] -= b[j] * A[i][j] / A[j][j]
			}
		}
		for i = 1; i <= n; i++ {
			b[i] /= A[i][i]
		}
	}
}

// ldl_dcmp - transpiled function from  $GOPATH/src/github.com/Konstantin8105/History_frame3DD/src/HPGmatrix.c:222
/*
 * LDL_DCMP  -  Solves [A]{x} = {b} simply and efficiently by performing an
 * L D L' - decomposition of [A].  No pivoting is performed.
 * [A] is a symmetric diagonally-dominant matrix of dimension [1..n][1..n].
 * {b} is a r.h.s. vector of dimension [1..n].
 * {b} is updated using L D L' and then back-substitution is done to obtain {x}
 * {b} is returned unchanged.  ldl_dcmp(A,n,d,x,x,1,1,&pd) is valid.
 *     The lower triangle of [A] is replaced by the lower triangle L of the
 *     L D L' reduction.  The diagonal of D is returned in the vector {d}
 *
 * usage: double **A, *d, *b, *x;
 *	int   n, reduce, solve, pd;
 *	ldl_dcmp ( A, n, d, b, x, reduce, solve, &pd );
 *
 * H.P. Gavin, Civil Engineering, Duke University, hpgavin@duke.edu  9 Oct 2001
 * Bathe, Finite Element Procecures in Engineering Analysis, Prentice Hall, 1982
 */ //
// the system matrix, and L of the L D L' decomp.
// the dimension of the matrix
// diagonal of D in the  L D L' - decomp'n
// the right hand side vector
// the solution vector
// 1: do a forward reduction of A; 0: don't
// 1: do a back substitution for {x}; 0: don't
// 1: definite matrix and successful L D L' decomp'n
/* number of negative elements on the diagonal of D */ //
/* forward column-wise reduction of [A]	*/ //
/* scan the sky-line	*/ //
/* the forward reduction of [A] is now complete	*/ //
/* back substitution to solve for {x}   */ //
/* {x} is run through the same forward reduction as was [A] */ //
/* now back substitution is conducted on {x};  [A] is preserved */ //
//
func ldl_dcmp(A [][]float64, n int, d []float64, b []float64, x []float64, reduce int, solve int, pd []int) {
	var i int
	var j int
	var k int
	var m int
	pd[0] = 0
	if reduce != 0 {
		for j = 1; j <= n; j++ {
			{
				m = 1
				i = 1
				for i = 1; i < j; i++ {
					if A[i][j] == 0 {
						m++
					} else {
						break
					}
				}
			}
			for i = m; i < j; i++ {
				A[j][i] = A[i][j]
				for k = m; k < i; k++ {
					A[j][i] -= A[j][k] * A[i][k]
				}
			}
			d[j] = A[j][j]
			for i = m; i < j; i++ {
				d[j] -= A[j][i] * A[j][i] / d[i]
			}
			for i = m; i < j; i++ {
				A[j][i] /= d[i]
			}
			if d[j] == 0 {
				noarch.Fprintf(stderr, []byte(" ldl_dcmp(): zero found on diagonal ...\n\x00"))
				noarch.Fprintf(stderr, []byte(" d[%d] = %11.4e\n\x00"), j, d[j])
				return
			}
			if d[j] < 0 {
				pd[0]--
			}
		}
	}
	if solve != 0 {
		for i = 1; i <= n; i++ {
			x[i] = b[i]
			for j = 1; j < i; j++ {
				x[i] -= A[i][j] * x[j]
			}
		}
		for i = 1; i <= n; i++ {
			x[i] /= d[i]
		}
		for i = n; i > 1; i-- {
			for j = 1; j < i; j++ {
				x[j] -= A[i][j] * x[i]
			}
		}
	}
}

// ldl_mprove - transpiled function from  $GOPATH/src/github.com/Konstantin8105/History_frame3DD/src/HPGmatrix.c:295
/*
 * LDL_MPROVE  Improves a solution vector x[1..n] of the linear set of equations
 * [A]{x} = {b}.  The matrix A[1..n][1..n], and the vectors b[1..n] and x[1..n]
 * are input, as is the dimension n.   The matrix [A] is the L D L'
 * decomposition of the original system matrix, as returned by ldl_dcmp().
 * Also input is the diagonal vector, {d} of [D] of the L D L' decompositon.
 * On output, only {x} is modified to an improved set of values.
 *
 * usage: double **A, *d, *b, *x, rms_resid;
 * 	int   n, ok;
 * 	ldl_mprove ( A, n, d, b, x, &rms_resid, &ok );
 *
 * H.P. Gavin, Civil Engineering, Duke University, hpgavin@duke.edu  4 May 2001
 */ //
/* accumulate the r.h.s. in double precision */ //
/* the residual error		  	*/ //
/* the RMS error of the mprvd solution	*/ //
/* allocate memory for a vector	of doubles */ //
/* calculate the r.h.s. of      */ //
/* [A]{r} = {b} - [A]{x+r}      */ //
/* A in upper triangle only     */ //
/* solve for the error term */ //
/* good improvement */ //
/* subtract the error from the old solution */ //
/* the solution has improved		*/ //
//
func ldl_mprove(A [][]float64, n int, d []float64, b []float64, x []float64, rms_resid []float64, ok []int) {
	var sdp float64
	var resid []float64
	var rms_resid_new float64
	var j int
	var i int
	var pd int
	resid = dvector(1, int32(n))
	for i = 1; i <= n; i++ {
		sdp = b[i]
		for j = 1; j <= n; j++ {
			if i <= j {
				sdp -= A[i][j] * x[j]
			} else {
				sdp -= A[j][i] * x[j]
			}
		}
		resid[i] = sdp
	}
	ldl_dcmp(A, n, d, resid, resid, 0, 1, (*[100000000]int)(unsafe.Pointer(&pd))[:])
	for i = 1; i <= n; i++ {
		rms_resid_new += resid[i] * resid[i]
	}
	rms_resid_new = math.Sqrt(rms_resid_new / float64(n))
	ok[0] = 0
	if rms_resid_new/rms_resid[0] < 0.9 {
		for i = 1; i <= n; i++ {
			x[i] += resid[i]
		}
		rms_resid[0] = rms_resid_new
		ok[0] = 1
	}
	free_dvector(resid, 1, int32(n))
}

// ldl_dcmp_pm - transpiled function from  $GOPATH/src/github.com/Konstantin8105/History_frame3DD/src/HPGmatrix.c:366
/*
 * LDL_DCMP_PM  -  Solves partitioned matrix equations
 *
 *           [A_qq]{x_q} + [A_qr]{x_r} = {b_q}
 *           [A_rq]{x_q} + [A_rr]{x_r} = {b_r}+{c_r}
 *           where {b_q}, {b_r}, and {x_r} are known and
 *           where {x_q} and {c_r} are unknown
 *
 * via L D L' - decomposition of [A_qq].  No pivoting is performed.
 * [A] is a symmetric diagonally-dominant matrix of dimension [1..n][1..n].
 * {b} is a r.h.s. vector of dimension [1..n].
 * {b} is updated using L D L' and then back-substitution is done to obtain {x}
 * {b_q} and {b_r}  are returned unchanged.
 * {c_r} is returned as a vector of [1..n] with {c_q}=0.
 * {q} is a vector of the indexes of known values {b_q}
 * {r} is a vector of the indexes of known values {x_r}
 *     The lower triangle of [A_qq] is replaced by the lower triangle L of its
 *     L D L' reduction.  The diagonal of D is returned in the vector {d}
 *
 * usage: double **A, *d, *b, *x;
 *	int   n, reduce, solve, pd;
 *	ldl_dcmp_pm ( A, n, d, b, x, c, q, r, reduce, solve, &pd );
 *
 * H.P. Gavin, Civil Engineering, Duke University, hpgavin@duke.edu
 * Bathe, Finite Element Procecures in Engineering Analysis, Prentice Hall, 1982
 * 2014-05-14
 */ //
/**< the system matrix, and L of the L D L' decomp.*/ //
/**< the dimension of the matrix		*/ //
/**< diagonal of D in the  L D L' - decomp'n    */ //
/**< the right hand side vector			*/ //
/**< part of the solution vector		*/ //
/**< the part of the solution vector in the rhs */ //
/**< q[j]=1 if  b[j] is known; q[j]=0 otherwise	*/ //
/**< r[j]=1 if  x[j] is known; r[j]=0 otherwise	*/ //
/**< 1: do a forward reduction of A; 0: don't   */ //
/**< 1: do a back substitution for {x}; 0: don't */ //
/**< 1: definite matrix and successful L D L' decomp'n*/ //
/* number of negative elements on the diagonal of D */ //
/* forward column-wise reduction of [A]	*/ //
/* reduce column j, except where q[i]==0	*/ //
/* scan the sky-line	*/ //
/* the forward reduction of [A] is now complete	*/ //
/* back substitution to solve for {x}   */ //
/* {x} is run through the same forward reduction as was [A] */ //
/* now back substitution is conducted on {x};  [A] is preserved */ //
/* finally, evaluate c_r	*/ //
// changed from 0.0 to -b[i]; 2014-05-14
//
func ldl_dcmp_pm(A [][]float64, n int, d []float64, b []float64, x []float64, c []float64, q []int, r []int, reduce int, solve int, pd []int) {
	var i int
	var j int
	var k int
	var m int
	pd[0] = 0
	if reduce != 0 {
		for j = 1; j <= n; j++ {
			d[j] = 0
			if q[j] != 0 {
				{
					m = 1
					i = 1
					for i = 1; i < j; i++ {
						if A[i][j] == 0 {
							m++
						} else {
							break
						}
					}
				}
				for i = m; i < j; i++ {
					if q[i] != 0 {
						A[j][i] = A[i][j]
						for k = m; k < i; k++ {
							if q[k] != 0 {
								A[j][i] -= A[j][k] * A[i][k]
							}
						}
					}
				}
				d[j] = A[j][j]
				for i = m; i < j; i++ {
					if q[i] != 0 {
						d[j] -= A[j][i] * A[j][i] / d[i]
					}
				}
				for i = m; i < j; i++ {
					if q[i] != 0 {
						A[j][i] /= d[i]
					}
				}
				if d[j] == 0 {
					noarch.Fprintf(stderr, []byte(" ldl_dcmp_pm(): zero found on diagonal ...\n\x00"))
					noarch.Fprintf(stderr, []byte(" d[%d] = %11.4e\n\x00"), j, d[j])
					return
				}
				if d[j] < 0 {
					pd[0]--
				}
			}
		}
	}
	if solve != 0 {
		for i = 1; i <= n; i++ {
			if q[i] != 0 {
				x[i] = b[i]
				for j = 1; j <= n; j++ {
					if r[j] != 0 {
						x[i] -= A[i][j] * x[j]
					}
				}
			}
		}
		for i = 1; i <= n; i++ {
			if q[i] != 0 {
				for j = 1; j < i; j++ {
					if q[j] != 0 {
						x[i] -= A[i][j] * x[j]
					}
				}
			}
		}
		for i = 1; i <= n; i++ {
			if q[i] != 0 {
				x[i] /= d[i]
			}
		}
		for i = n; i > 1; i-- {
			if q[i] != 0 {
				for j = 1; j < i; j++ {
					if q[j] != 0 {
						x[j] -= A[i][j] * x[i]
					}
				}
			}
		}
		for i = 1; i <= n; i++ {
			c[i] = 0
			if r[i] != 0 {
				c[i] = -b[i]
				for j = 1; j <= n; j++ {
					c[i] += A[i][j] * x[j]
				}
			}
		}
	}
}

// ldl_mprove_pm - transpiled function from  $GOPATH/src/github.com/Konstantin8105/History_frame3DD/src/HPGmatrix.c:479
/*
 * LDL_MPROVE_PM
 * Improves a solution vector x[1..n] of the partitioned set of linear equations
 *           [A_qq]{x_q} + [A_qr]{x_r} = {b_q}
 *           [A_rq]{x_q} + [A_rr]{x_r} = {b_r}+{c_r}
 *           where {b_q}, {b_r}, and {x_r} are known and
 *           where {x_q} and {c_r} are unknown
 * by reducing the residual r_q
 *           A_qq r_q = {b_q} - [A_qq]{x_q+r_q} + [A_qr]{x_r}
 * The matrix A[1..n][1..n], and the vectors b[1..n] and x[1..n]
 * are input, as is the dimension n.   The matrix [A_qq] is the L D L'
 * decomposition of the original system matrix, as returned by ldl_dcmp_pm().
 * Also input is the diagonal vector, {d} of [D] of the L D L' decompositon.
 * On output, only {x} and {c} are modified to an improved set of values.
 * The partial right-hand-side vectors, {b_q} and {b_r}, are returned unchanged.
 * Further, the calculations in ldl_mprove_pm do not involve b_r.
 *
 * usage: double **A, *d, *b, *x, rms_resid;
 * 	int   n, ok, *q, *r;
 *	ldl_mprove_pm ( A, n, d, b, x, q, r, &rms_resid, &ok );
 *
 * H.P. Gavin, Civil Engineering, Duke University, hpgavin@duke.edu
 * 2001-05-01, 2014-05-14
 */ //
/**< the system matrix, and L of the L D L' decomp.*/ //
/**< the dimension of the matrix		*/ //
/**< diagonal of D in the  L D L' - decomp'n    */ //
/**< the right hand side vector			*/ //
/**< part of the solution vector		*/ //
/**< the part of the solution vector in the rhs */ //
/**< q[j]=1 if  b[j] is known; q[j]=0 otherwise	*/ //
/**< r[j]=1 if  x[j] is known; r[j]=0 otherwise	*/ //
/**< root-mean-square of residual error	*/ //
/**< 1: >10% reduction in rms_resid; 0: not	*/ //
// accumulate the r.h.s. in double precision
// the residual error
// update to partial r.h.s. vector, c
// the RMS error of the mprvd solution
// allocate memory for a vector	of doubles
// calculate the r.h.s. of ...
//  [A_qq]{dx_q} = {b_q} - [A_qq]*{x_q} - [A_qr]*{x_r}
//  {dx_r} is left unchanged at 0.0;
// A_qq in upper triangle only
// else dx[i] = 0.0; // x[i];
// solve for the residual error term, A is already factored
/*  enough improvement    */ //
/*  update the solution 2014-05-14   */ //
/* return the new residual   */ //
/* the solution has improved */ //
//
func ldl_mprove_pm(A [][]float64, n int, d []float64, b []float64, x []float64, c []float64, q []int, r []int, rms_resid []float64, ok []int) {
	var sdp float64
	var dx []float64
	var dc []float64
	var rms_resid_new float64
	var j int
	var i int
	var pd int
	dx = dvector(1, int32(n))
	dc = dvector(1, int32(n))
	for i = 1; i <= n; i++ {
		dx[i] = 0
	}
	for i = 1; i <= n; i++ {
		if q[i] != 0 {
			sdp = b[i]
			for j = 1; j <= n; j++ {
				if q[j] != 0 {
					if i <= j {
						sdp -= A[i][j] * x[j]
					} else {
						sdp -= A[j][i] * x[j]
					}
				}
			}
			for j = 1; j <= n; j++ {
				if r[j] != 0 {
					sdp -= A[i][j] * x[j]
				}
			}
			dx[i] = sdp
		}
	}
	ldl_dcmp_pm(A, n, d, dx, dx, dc, q, r, 0, 1, (*[100000000]int)(unsafe.Pointer(&pd))[:])
	for i = 1; i <= n; i++ {
		if q[i] != 0 {
			rms_resid_new += dx[i] * dx[i]
		}
	}
	rms_resid_new = math.Sqrt(rms_resid_new / float64(n))
	ok[0] = 0
	if rms_resid_new/rms_resid[0] < 0.9 {
		for i = 1; i <= n; i++ {
			if q[i] != 0 {
				x[i] += dx[i]
			}
			if r[i] != 0 {
				c[i] += dc[i]
			}
		}
		rms_resid[0] = rms_resid_new
		ok[0] = 1
	}
	free_dvector(dx, 1, int32(n))
	free_dvector(dc, 1, int32(n))
}

// PSB_update - transpiled function from  $GOPATH/src/github.com/Konstantin8105/History_frame3DD/src/HPGmatrix.c:553
/*
 * PSB_UPDATE
 * Update secant stiffness matrix via the Powell-Symmetric-Broyden update eqn.
 *
 *       B = B - (f*d' + d*f') / (d' * d) + f'*d * d*d' / (d' * d)^2 ;
 *
 * H.P. Gavin, Civil Engineering, Duke University, hpgavin@duke.edu  24 Oct 2012
 */ //
/**< secant stiffness matrix            */ //
/**< out-of-balance force vector        */ //
/**< incremental displacement vector    */ //
/**< matrix dimension is n-by-n         */ //
/*  update upper triangle of B[i][j] */ //
//
func PSB_update(B [][]float64, f []float64, d []float64, n int) {
	var i int
	var j int
	var dtd float64
	var ftd float64
	var dtd2 float64
	for i = 1; i <= n; i++ {
		dtd += d[i] * d[i]
	}
	dtd2 = dtd * dtd
	for i = 1; i <= n; i++ {
		ftd += f[i] * d[i]
	}
	for i = 1; i <= n; i++ {
		for j = i; j <= n; j++ {
			B[i][j] -= (f[i]*d[j]+f[j]*d[i])/dtd - ftd*d[i]*d[j]/dtd2
		}
	}
}

// pseudo_inv - transpiled function from  $GOPATH/src/github.com/Konstantin8105/History_frame3DD/src/HPGmatrix.c:580
/*
 * PSEUDO_INV - calculate the pseudo-inverse of A ,
 * 	     Ai = inv ( A'*A + beta * trace(A'*A) * I ) * A'
 *	     beta is a regularization factor, which should be small (1e-10)
 *	     A is m by n      Ai is m by n			      8oct01
 */ //
/* compute A' * A */ //
/* make symmetric */ //
/* trace of AtA */ //
/* add beta I */ //
/*  L D L'  decomp */ //
/* compute inv(AtA) */ //
/* L D L' bksbtn */ //
/*improve the solution*/ //
/* save inv(AtA) */ //
/* make symmetric */ //
/* compute inv(A'*A)*A'	*/ //
//
func pseudo_inv(A [][]float64, Ai [][]float64, n int, m int, beta float64, verbose int) {
	var diag []float64
	var b []float64
	var x []float64
	var AtA [][]float64
	var AtAi [][]float64
	var tmp float64
	var tr_AtA float64
	var error float64
	var i int
	var j int
	var k int
	var ok int
	diag = dvector(1, int32(n))
	b = dvector(1, int32(n))
	x = dvector(1, int32(n))
	AtA = dmatrix(1, int32(n), 1, int32(n))
	AtAi = dmatrix(1, int32(n), 1, int32(n))
	if beta > 1 {
		noarch.Fprintf(stderr, []byte(" pseudo_inv: warning beta = %lf\n\x00"), beta)
	}
	for i = 1; i <= n; i++ {
		b[i] = 0
		x[i] = b[i]
		diag[i] = x[i]
		for j = i; j <= n; j++ {
			AtA[j][i] = 0
			AtA[i][j] = AtA[j][i]
		}
	}
	for i = 1; i <= n; i++ {
		for j = 1; j <= n; j++ {
			tmp = 0
			for k = 1; k <= m; k++ {
				tmp += A[k][i] * A[k][j]
			}
			AtA[i][j] = tmp
		}
	}
	for i = 1; i <= n; i++ {
		for j = i; j <= n; j++ {
			AtA[j][i] = 0.5 * (AtA[i][j] + AtA[j][i])
			AtA[i][j] = AtA[j][i]
		}
	}
	tr_AtA = 0
	for i = 1; i <= n; i++ {
		tr_AtA += AtA[i][i]
	}
	for i = 1; i <= n; i++ {
		AtA[i][i] += beta * tr_AtA
	}
	ldl_dcmp(AtA, n, diag, b, x, 1, 0, (*[100000000]int)(unsafe.Pointer(&ok))[:])
	for j = 1; j <= n; j++ {
		for k = 1; k <= n; k++ {
			b[k] = 0
		}
		b[j] = 1
		ldl_dcmp(AtA, n, diag, b, x, 0, 1, (*[100000000]int)(unsafe.Pointer(&ok))[:])
		if verbose != 0 {
			noarch.Fprintf(stdout, []byte("  RMS matrix error:\x00"))
		}
		error = 1
		ok = 1
		for {
			ldl_mprove(AtA, n, diag, b, x, (*[100000000]float64)(unsafe.Pointer(&error))[:], (*[100000000]int)(unsafe.Pointer(&ok))[:])
			if verbose != 0 {
				noarch.Fprintf(stdout, []byte("%9.2e\x00"), error)
			}
			if noarch.NotInt((ok)) != 0 {
				break
			}
		}
		if verbose != 0 {
			noarch.Fprintf(stdout, []byte("\n\x00"))
		}
		for k = 1; k <= n; k++ {
			AtAi[k][j] = x[k]
		}
	}
	for i = 1; i <= n; i++ {
		for j = i; j <= n; j++ {
			AtAi[j][i] = 0.5 * (AtAi[i][j] + AtAi[j][i])
			AtAi[i][j] = AtAi[j][i]
		}
	}
	for i = 1; i <= n; i++ {
		for j = 1; j <= m; j++ {
			tmp = 0
			for k = 1; k <= n; k++ {
				tmp += AtAi[i][k] * A[j][k]
			}
			Ai[i][j] = tmp
		}
	}
	free_dmatrix(AtAi, 1, int32(n), 1, int32(n))
	free_dmatrix(AtA, 1, int32(n), 1, int32(n))
	free_dvector(x, 1, int32(n))
	free_dvector(b, 1, int32(n))
	free_dvector(diag, 1, int32(n))
}

// prodABj - transpiled function from  $GOPATH/src/github.com/Konstantin8105/History_frame3DD/src/HPGmatrix.c:662
/*
 * PRODABj -  matrix-matrix multiplication for symmetric A	      27apr01
 *		 u = A * B(:,j)
 */ //
//
func prodABj(A [][]float64, B [][]float64, u []float64, n int, j int) {
	var i int
	var k int
	for i = 1; i <= n; i++ {
		u[i] = 0
	}
	for i = 1; i <= n; i++ {
		for k = 1; k <= n; k++ {
			if i <= k {
				u[i] += A[i][k] * B[k][j]
			} else {
				u[i] += A[k][i] * B[k][j]
			}
		}
	}
}

// prodAB - transpiled function from  $GOPATH/src/github.com/Konstantin8105/History_frame3DD/src/HPGmatrix.c:681
/*
 * prodAB - matrix-matrix multiplication      C = A * B			27apr01
 */ //
//
func prodAB(A [][]float64, B [][]float64, C [][]float64, I int, J int, K int) {
	var i int
	var j int
	var k int
	for i = 1; i <= I; i++ {
		for k = 1; k <= K; k++ {
			C[i][k] = 0
		}
	}
	for i = 1; i <= I; i++ {
		for k = 1; k <= K; k++ {
			for j = 1; j <= J; j++ {
				C[i][k] += A[i][j] * B[j][k]
			}
		}
	}
}

// invAB - transpiled function from  $GOPATH/src/github.com/Konstantin8105/History_frame3DD/src/HPGmatrix.c:701
/*
 * INVAB  -  calculate product inv(A) * B
 *	 A is n by n      B is n by m				    6jun07
 */ //
/*   L D L'  decomp */ //
/*   L D L'  bksbtn */ //
/* improve the solution*/ //
//
func invAB(A [][]float64, B [][]float64, n int, m int, AiB [][]float64, ok []int, verbose int) {
	var diag []float64
	var b []float64
	var x []float64
	var error float64
	var i int
	var j int
	var k int
	diag = dvector(1, int32(n))
	x = dvector(1, int32(n))
	b = dvector(1, int32(n))
	for i = 1; i <= n; i++ {
		x[i] = 0
		diag[i] = x[i]
	}
	ldl_dcmp(A, n, diag, b, x, 1, 0, ok)
	if ok[0] < 0 {
		noarch.Fprintf(stderr, []byte(" Make sure that all six\x00"))
		noarch.Fprintf(stderr, []byte(" rigid body translations are restrained!\n\x00"))
	}
	for j = 1; j <= m; j++ {
		for k = 1; k <= n; k++ {
			b[k] = B[k][j]
		}
		ldl_dcmp(A, n, diag, b, x, 0, 1, ok)
		if verbose != 0 {
			noarch.Fprintf(stdout, []byte("    LDL' RMS matrix precision:\x00"))
		}
		error = float64((func() int {
			ok[0] = 1
			return ok[0]
		}()))
		for {
			ldl_mprove(A, n, diag, b, x, (*[100000000]float64)(unsafe.Pointer(&error))[:], ok)
			if verbose != 0 {
				noarch.Fprintf(stdout, []byte("%9.2e\x00"), error)
			}
			if noarch.NotInt((ok[0])) != 0 {
				break
			}
		}
		if verbose != 0 {
			noarch.Fprintf(stdout, []byte("\n\x00"))
		}
		for i = 1; i <= n; i++ {
			AiB[i][j] = x[i]
		}
	}
	free_dvector(diag, 1, int32(n))
	free_dvector(x, 1, int32(n))
	free_dvector(b, 1, int32(n))
}

// xtinvAy - transpiled function from  $GOPATH/src/github.com/Konstantin8105/History_frame3DD/src/HPGmatrix.c:744
/*
 * XTinvAY  -  calculate quadratic form with inverse matrix   X' * inv(A) * Y
 *	   A is n by n    X is n by m     Y is n by m		    15sep01
 */ //
/*   L D L'  decomp */ //
/*   L D L'  bksbtn */ //
/* improve the solution*/ //
//
func xtinvAy(X [][]float64, A [][]float64, Y [][]float64, n int, m int, Ac [][]float64, verbose int) {
	var diag []float64
	var x []float64
	var y []float64
	var error float64
	var i int
	var j int
	var k int
	var ok int
	diag = dvector(1, int32(n))
	x = dvector(1, int32(n))
	y = dvector(1, int32(n))
	for i = 1; i <= n; i++ {
		x[i] = 0
		diag[i] = x[i]
	}
	ldl_dcmp(A, n, diag, y, x, 1, 0, (*[100000000]int)(unsafe.Pointer(&ok))[:])
	for j = 1; j <= m; j++ {
		for k = 1; k <= n; k++ {
			y[k] = Y[k][j]
		}
		ldl_dcmp(A, n, diag, y, x, 0, 1, (*[100000000]int)(unsafe.Pointer(&ok))[:])
		if verbose != 0 {
			noarch.Fprintf(stdout, []byte("    LDL' RMS matrix precision:\x00"))
		}
		error = float64((func() int {
			ok = 1
			return ok
		}()))
		for {
			ldl_mprove(A, n, diag, y, x, (*[100000000]float64)(unsafe.Pointer(&error))[:], (*[100000000]int)(unsafe.Pointer(&ok))[:])
			if verbose != 0 {
				noarch.Fprintf(stdout, []byte("%9.2e\x00"), error)
			}
			if noarch.NotInt((ok)) != 0 {
				break
			}
		}
		if verbose != 0 {
			noarch.Fprintf(stdout, []byte("\n\x00"))
		}
		for i = 1; i <= m; i++ {
			Ac[i][j] = 0
			for k = 1; k <= n; k++ {
				Ac[i][j] += X[k][i] * x[k]
			}
		}
	}
	free_dvector(diag, 1, int32(n))
	free_dvector(x, 1, int32(n))
	free_dvector(y, 1, int32(n))
}

// coord_xfrm - transpiled function from  $GOPATH/src/github.com/Konstantin8105/History_frame3DD/src/HPGmatrix.c:787
/*  COORD_XFRM - coordinate transform of a matrix of column 2-vectors
 *
 * Rr  = [ cosd(theta) -sind(theta) ; sind(theta) cosd(theta) ]*[ Rx ; Ry ];
 */ //
//
func coord_xfrm(Rr [][]float32, R [][]float32, theta float32, n int) {
	var R1 float32
	var R2 float32
	var i int
	for i = 1; i <= n; i++ {
		R1 = float32(math.Cos(math.Mod(float64((theta)), 360)*3.141592653589793/180)*float64(R[1][i]) - math.Sin(math.Mod(float64((theta)), 360)*3.141592653589793/180)*float64(R[2][i]))
		R2 = float32(math.Sin(math.Mod(float64((theta)), 360)*3.141592653589793/180)*float64(R[1][i]) + math.Cos(math.Mod(float64((theta)), 360)*3.141592653589793/180)*float64(R[2][i]))
		Rr[1][i] = R1
		Rr[2][i] = R2
	}
}

// xtAx - transpiled function from  $GOPATH/src/github.com/Konstantin8105/History_frame3DD/src/HPGmatrix.c:806
/*
 * xtAx - carry out matrix-matrix-matrix multiplication for symmetric A  7nov02
 *	 C = X' A X     C is J by J      X is N by J     A is N by N
 */ //
/*  use upper triangle of A */ //
/*  make  C  symmetric */ //
//
func xtAx(A [][]float64, X [][]float64, C [][]float64, N int, J int) {
	var AX [][]float64
	var i int
	var j int
	var k int
	AX = dmatrix(1, int32(N), 1, int32(J))
	for i = 1; i <= J; i++ {
		for j = 1; j <= J; j++ {
			C[i][j] = 0
		}
	}
	for i = 1; i <= N; i++ {
		for j = 1; j <= J; j++ {
			AX[i][j] = 0
		}
	}
	for i = 1; i <= N; i++ {
		for j = 1; j <= J; j++ {
			for k = 1; k <= N; k++ {
				if i <= k {
					AX[i][j] += A[i][k] * X[k][j]
				} else {
					AX[i][j] += A[k][i] * X[k][j]
				}
			}
		}
	}
	for i = 1; i <= J; i++ {
		for j = 1; j <= J; j++ {
			for k = 1; k <= N; k++ {
				C[i][j] += X[k][i] * AX[k][j]
			}
		}
	}
	for i = 1; i <= J; i++ {
		for j = i; j <= J; j++ {
			C[j][i] = 0.5 * (C[i][j] + C[j][i])
			C[i][j] = C[j][i]
		}
	}
	free_dmatrix(AX, 1, int32(N), 1, int32(J))
}

// xtAy - transpiled function from  $GOPATH/src/github.com/Konstantin8105/History_frame3DD/src/HPGmatrix.c:843
/*
 * xtAy - carry out vector-matrix-vector multiplication for symmetric A  7apr94
 */ //
/*  d = A y  */ //
//  A in upper triangle only
/*  xAy = x' A y  */ //
//
func xtAy(x []float64, A [][]float64, y []float64, n int, d []float64) float64 {
	var xtAy float64
	var i int
	var j int
	for i = 1; i <= n; i++ {
		d[i] = 0
		for j = 1; j <= n; j++ {
			if i <= j {
				d[i] += A[i][j] * y[j]
			} else {
				d[i] += A[j][i] * y[j]
			}
		}
	}
	for i = 1; i <= n; i++ {
		xtAy += x[i] * d[i]
	}
	return (xtAy)
}

// invAXinvA - transpiled function from  $GOPATH/src/github.com/Konstantin8105/History_frame3DD/src/HPGmatrix.c:866
/*
 * invAXinvA -  calculate quadratic form with inverse matrix
 *	   replace X with inv(A) * X * inv(A)
 *	   A is n by n and symmetric   X is n by n and symmetric    15sep01
 */ //
/*   L D L'  decomp */ //
/*  compute inv(A)  */ //
/*   L D L'  bksbtn */ //
/* improve the solution*/ //
/*  save inv(A) */ //
/*  make symmetric */ //
/*  compute X * inv(A)   */ //
/*  compute inv(A) * X * inv(A)  */ //
/*  make symmetric */ //
//
func invAXinvA(A [][]float64, X [][]float64, n int, verbose int) {
	var diag []float64
	var b []float64
	var x []float64
	var Ai [][]float64
	var XAi [][]float64
	var Aij float64
	var error float64
	var i int
	var j int
	var k int
	var ok int
	diag = dvector(1, int32(n))
	x = dvector(1, int32(n))
	b = dvector(1, int32(n))
	Ai = dmatrix(1, int32(n), 1, int32(n))
	XAi = dmatrix(1, int32(n), 1, int32(n))
	for i = 1; i <= n; i++ {
		b[i] = 0
		x[i] = b[i]
		diag[i] = x[i]
		for j = 1; j <= n; j++ {
			Ai[i][j] = 0
			XAi[i][j] = Ai[i][j]
		}
	}
	ldl_dcmp(A, n, diag, b, x, 1, 0, (*[100000000]int)(unsafe.Pointer(&ok))[:])
	for j = 1; j <= n; j++ {
		for k = 1; k <= n; k++ {
			b[k] = 0
		}
		b[j] = 1
		ldl_dcmp(A, n, diag, b, x, 0, 1, (*[100000000]int)(unsafe.Pointer(&ok))[:])
		if verbose != 0 {
			noarch.Fprintf(stdout, []byte("    LDL' RMS matrix precision:\x00"))
		}
		error = float64((func() int {
			ok = 1
			return ok
		}()))
		for {
			ldl_mprove(A, n, diag, b, x, (*[100000000]float64)(unsafe.Pointer(&error))[:], (*[100000000]int)(unsafe.Pointer(&ok))[:])
			if verbose != 0 {
				noarch.Fprintf(stdout, []byte("%9.2e\x00"), error)
			}
			if noarch.NotInt((ok)) != 0 {
				break
			}
		}
		if verbose != 0 {
			noarch.Fprintf(stdout, []byte("\n\x00"))
		}
		for k = 1; k <= n; k++ {
			Ai[j][k] = x[k]
		}
	}
	for i = 1; i <= n; i++ {
		for j = i; j <= n; j++ {
			Ai[j][i] = 0.5 * (Ai[i][j] + Ai[j][i])
			Ai[i][j] = Ai[j][i]
		}
	}
	for i = 1; i <= n; i++ {
		for j = 1; j <= n; j++ {
			Aij = 0
			for k = 1; k <= n; k++ {
				Aij += X[i][k] * Ai[k][j]
			}
			XAi[i][j] = Aij
		}
	}
	for i = 1; i <= n; i++ {
		for j = 1; j <= n; j++ {
			Aij = 0
			for k = 1; k <= n; k++ {
				Aij += Ai[i][k] * XAi[k][j]
			}
			X[i][j] = Aij
		}
	}
	for i = 1; i <= n; i++ {
		for j = i; j <= n; j++ {
			X[j][i] = 0.5 * (X[i][j] + X[j][i])
			X[i][j] = X[j][i]
		}
	}
	free_dvector(diag, 1, int32(n))
	free_dvector(x, 1, int32(n))
	free_dvector(b, 1, int32(n))
	free_dmatrix(Ai, 1, int32(n), 1, int32(n))
	free_dmatrix(XAi, 1, int32(n), 1, int32(n))
}

// relative_norm - transpiled function from  $GOPATH/src/github.com/Konstantin8105/History_frame3DD/src/HPGmatrix.c:939
/*
 *  RELATIVE_NORM -  compute the relative 2-norm between two vectors       26dec01
 *       compute the relative 2-norm between two vectors N and D
 *	       return  ( sqrt(sum(N[i]*N[i]) / sqrt(D[i]*D[i]) )
 *
 */ //
//
func relative_norm(N []float64, D []float64, n int) float64 {
	var nN float64
	var nD float64
	var i int
	for i = 1; i <= n; i++ {
		nN += N[i] * N[i]
	}
	for i = 1; i <= n; i++ {
		nD += D[i] * D[i]
	}
	return (math.Sqrt(nN) / math.Sqrt(nD))
}

// Legendre - transpiled function from  $GOPATH/src/github.com/Konstantin8105/History_frame3DD/src/HPGmatrix.c:955
/*
 *  Legendre
 *  compute matrix of the Legendre polynomials and its first two derivitives
 */ //
// save_vector( n, t, "t.dat");
//
func Legendre(order int, t []float32, n int, P [][]float32, Pp [][]float32, Ppp [][]float32) {
	var k int
	var p int
	for p = 1; p <= n; p++ {
		P[0][p] = float32(1)
		P[1][p] = t[p]
		P[2][p] = float32(1.5*float64(t[p])*float64(t[p]) - 0.5)
		P[3][p] = float32(2.5*float64(t[p])*float64(t[p])*float64(t[p]) - 1.5*float64(t[p]))
		Pp[0][p] = float32(0)
		Pp[1][p] = float32(1)
		Pp[2][p] = float32(3 * float64(t[p]))
		Pp[3][p] = float32(7.5*float64(t[p])*float64(t[p]) - 1.5)
		Ppp[0][p] = float32(0)
		Ppp[1][p] = float32(0)
		Ppp[2][p] = float32(3)
		Ppp[3][p] = float32(15 * float64(t[p]))
		for k = 4; k <= order; k++ {
			P[k][p] = float32((2-1/float64(k))*float64(t[p])*float64(P[k-1][p]) - (1-1/float64(k))*float64(P[k-2][p]))
			Pp[k][p] = float32((2-1/float64(k))*float64(P[k-1][p]+t[p]*Pp[k-1][p]) - (1-1/float64(k))*float64(Pp[k-2][p]))
			Ppp[k][p] = float32((2-1/float64(k))*float64(2*Pp[k-1][p]+t[p]*Ppp[k-1][p]) - (1-1/float64(k))*float64(Ppp[k-2][p]))
		}
	}
}

// color - transpiled function from  $GOPATH/src/github.com/Konstantin8105/History_frame3DD/src/HPGutil.c:35
/*  HPGutil.c  ---  library of general-purpose utility functions	*/ //
/*
 Copyright (C) 2012 Henri P. Gavin

    HPGutil is free software: you can redistribute it and/or modify
    it under the terms of the GNU General Public License as published by
    the Free Software Foundation, either version 3 of the License, or
    (at your option) any later version.

    HPGutil is distributed in the hope that it will be useful,
    but WITHOUT ANY WARRANTY; without even the implied warranty of
    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
    GNU General Public License for more details.

    You should have received a copy of the GNU General Public License
    along with HPGutil.  If not, see <http://www.gnu.org/licenses/>.
*/ //
/*
 * COLOR - change color on the screen ...
 * Screen   Color  Scheme  : 0 = white on black, 1 = bright
 * first digit= 3  for text color	  first digit= 4  for  background color
 * second digit codes:
 * 0=black, 1=red, 2=green, 3=gold, 4=blue, 5=magenta, 6=cyan, 7=white
 * http://en.wikipedia.org/wiki/ANSI_escape_code
 */ //
/*  change the screen color      */ //
//
func color(colorCode int) {
	noarch.Fprintf(stderr, []byte("\x1b[%02dm\x00"), colorCode)
	_ = noarch.Fflush(stderr)
}

// textColor - transpiled function from  $GOPATH/src/github.com/Konstantin8105/History_frame3DD/src/HPGutil.c:53
/*
 * TEXTCOLOR - change color of text and background
 * tColor : text color : one of 'k' 'r' 'g' 'y' 'b' 'm' 'c' 'w'
 * bColor : back color : one of 'k' 'r' 'g' 'y' 'b' 'm' 'c' 'w'
 * nbf    : 'n' = normal, 'b' = bright/bold, 'f' = faint
 * uline  : 'u' = underline
 * http://en.wikipedia.org/wiki/ANSI_escape_code
 */ //
// Control Sequence Introducer & reset
// background colors
// black
// red
// green
// yellow
// blue
// magenta
// cyan
// white
// text colors
// black
// red
// green
// yellow
// blue
// magenta
// cyan
// white
//	printf(" tColor = %c   bColor = %c   nbf = %c\n", tColor, bColor, nbf );
// bright
// faint
// underline
// Select Graphic Rendition (SGR)
//
func textColor(tColor byte, bColor byte, nbf byte, uline byte) {
	noarch.Fprintf(stderr, []byte("\x1b[%02d\x00"), 0)
	if int(bColor) == int('k') {
		noarch.Fprintf(stderr, []byte(";%02d\x00"), 40)
	}
	if int(bColor) == int('r') {
		noarch.Fprintf(stderr, []byte(";%02d\x00"), 41)
	}
	if int(bColor) == int('g') {
		noarch.Fprintf(stderr, []byte(";%02d\x00"), 42)
	}
	if int(bColor) == int('y') {
		noarch.Fprintf(stderr, []byte(";%02d\x00"), 43)
	}
	if int(bColor) == int('b') {
		noarch.Fprintf(stderr, []byte(";%02d\x00"), 44)
	}
	if int(bColor) == int('m') {
		noarch.Fprintf(stderr, []byte(";%02d\x00"), 45)
	}
	if int(bColor) == int('c') {
		noarch.Fprintf(stderr, []byte(";%02d\x00"), 46)
	}
	if int(bColor) == int('w') {
		noarch.Fprintf(stderr, []byte(";%02d\x00"), 47)
	}
	if int(tColor) == int('k') {
		noarch.Fprintf(stderr, []byte(";%02d\x00"), 30)
	}
	if int(tColor) == int('r') {
		noarch.Fprintf(stderr, []byte(";%02d\x00"), 31)
	}
	if int(tColor) == int('g') {
		noarch.Fprintf(stderr, []byte(";%02d\x00"), 32)
	}
	if int(tColor) == int('y') {
		noarch.Fprintf(stderr, []byte(";%02d\x00"), 33)
	}
	if int(tColor) == int('b') {
		noarch.Fprintf(stderr, []byte(";%02d\x00"), 34)
	}
	if int(tColor) == int('m') {
		noarch.Fprintf(stderr, []byte(";%02d\x00"), 35)
	}
	if int(tColor) == int('c') {
		noarch.Fprintf(stderr, []byte(";%02d\x00"), 36)
	}
	if int(tColor) == int('w') {
		noarch.Fprintf(stderr, []byte(";%02d\x00"), 37)
	}
	if int(nbf) == int('b') {
		noarch.Fprintf(stderr, []byte(";%02d\x00"), 1)
	}
	if int(nbf) == int('f') {
		noarch.Fprintf(stderr, []byte(";%02d\x00"), 2)
	}
	if int(uline) == int('u') {
		noarch.Fprintf(stderr, []byte(";%02d\x00"), 4)
	}
	noarch.Fprintf(stderr, []byte("m\x00"))
	_ = noarch.Fflush(stderr)
}

// errorMsg - transpiled function from  $GOPATH/src/github.com/Konstantin8105/History_frame3DD/src/HPGutil.c:94
/*
 * ERRORMSG -  write a diagnostic error message in color
 */ //
//
func errorMsg(errString []byte) {
	noarch.Fprintf(stderr, []byte("\n\n\x00"))
	noarch.Fflush(stderr)
	color(1)
	color(41)
	color(37)
	noarch.Fprintf(stderr, []byte("  %s  \x00"), errString)
	noarch.Fflush(stderr)
	color(0)
	noarch.Fprintf(stderr, []byte("\n\n\x00"))
}

// openFile - transpiled function from  $GOPATH/src/github.com/Konstantin8105/History_frame3DD/src/HPGutil.c:114
// Warning (*ast.SwitchStmt):  /home/konstantin/go/src/github.com/Konstantin8105/History_frame3DD/src/HPGutil.c:126 :Unsupport case
/*
 * OPENFILE  -  open a file or print a diagnostic error message
 */ //
// open file
//
func openFile(path []byte, fileName []byte, mode []byte, usage []byte) (c4goDefaultReturn *noarch.File) {
	var fp *noarch.File
	var pathToFile []byte = make([]byte, 512)
	var errMsg []byte = make([]byte, 512)
	if mode == nil {
		return nil
	}
	noarch.Sprintf(pathToFile, []byte("%s%s\x00"), path, fileName)
	if (func() *noarch.File {
		fp = noarch.Fopen(pathToFile, mode)
		return fp
	}()) == nil {
		noarch.Sprintf(errMsg, []byte(" openFile: \x00"))
		switch int(mode[0]) {
		case 'w':
			noarch.Sprintf(errMsg, []byte("%s%s\n  usage: %s\x00"), []byte("cannot write to file: \x00"), pathToFile, usage)
		case 'r':
			noarch.Sprintf(errMsg, []byte("%s%s\n  usage: %s\x00"), []byte("cannot read from file: \x00"), pathToFile, usage)
		case 'a':
			noarch.Sprintf(errMsg, []byte("%s%s\n  usage: %s\x00"), []byte("cannot append to file: \x00"), pathToFile, usage)
		default:
			{
				noarch.Sprintf(errMsg, []byte("%s%s\n  usage: %s\x00"), []byte("cannot open file: \x00"), pathToFile, usage)
			}
		}
		errorMsg(errMsg)
		os.Exit(1)
	} else {
		return fp
	}
	return
}

// scanLine - transpiled function from  $GOPATH/src/github.com/Konstantin8105/History_frame3DD/src/HPGutil.c:155
/*
 * SCANLINE -  scan through a line until a 'a' is reached, like getline() 3feb94
 */ //
//
func scanLine(fp *noarch.File, lim int, s []byte, a byte) int {
	var c int
	var i int = -1
	for func() int {
		lim--
		return lim
	}() > 0 && (func() int {
		c = noarch.Fgetc(fp)
		return c
	}()) != -1 && c != int(a) {
		s[func() int {
			i++
			return i
		}()] = byte(c)
	}
	s[func() int {
		i++
		return i
	}()] = '\x00'
	return i
}

// scanLabel - transpiled function from  $GOPATH/src/github.com/Konstantin8105/History_frame3DD/src/HPGutil.c:168
/*
 * SCANLABEL -  scan through a line until a '"' is reached, like getline()
 */ //
// scan to first delimitter char
// read the label between delimitters
//
func scanLabel(fp *noarch.File, lim int, s []byte, a byte) int {
	var c int
	var i int = -1
	for func() int {
		lim--
		return lim
	}() > 0 && (func() int {
		c = noarch.Fgetc(fp)
		return c
	}()) != -1 && c != int(a) {
	}
	for func() int {
		lim--
		return lim
	}() > 0 && (func() int {
		c = noarch.Fgetc(fp)
		return c
	}()) != -1 && c != int(a) {
		s[func() int {
			i++
			return i
		}()] = byte(c)
	}
	s[func() int {
		i++
		return i
	}()] = '\x00'
	return i
}

// scanFile - transpiled function from  $GOPATH/src/github.com/Konstantin8105/History_frame3DD/src/HPGutil.c:185
/*
 * SCANFILE -  count the number of lines of multi-column data in a data file,
 * skipping over "head_lines" lines of header information
 */ //
// scan through the header
// count the number of lines of data
// printf ("%% %d data points\n", points);
//
func scanFile(fp *noarch.File, head_lines int, start_chnl int, stop_chnl int) int {
	var points int
	var i int
	var chn int
	var ok int = 1
	var data_value float32
	var ch byte
	for i = 1; i <= head_lines; i++ {
		for int((func() byte {
			ch = byte(noarch.Fgetc(fp))
			return ch
		}())) != int('\n') {
		}
	}
	for {
		for chn = start_chnl; chn <= stop_chnl; chn++ {
			ok = noarch.Fscanf(fp, []byte("%f\x00"), (*[100000000]float32)(unsafe.Pointer(&data_value))[:])
			if ok == 1 {
				points++
			}
		}
		if ok > 0 {
			for int((func() byte {
				ch = byte(noarch.Fgetc(fp))
				return ch
			}())) != int('\n') {
			}
		}
		if !(ok == 1) {
			break
		}
	}
	points = (points / (stop_chnl - start_chnl + 1))
	noarch.Rewind(fp)
	return (points)
}

// getLine - transpiled function from  $GOPATH/src/github.com/Konstantin8105/History_frame3DD/src/HPGutil.c:217
/*
 * GETLINE -  get line form a stream into a character string, return length
 * from K&R	       3feb94
 */ //
/*	if (c == '\n')  s[i++] = c;     */ //
//
func getLine(fp *noarch.File, lim int, s []byte) int {
	var c int
	var i int
	for func() int {
		lim--
		return lim
	}() > 0 && (func() int {
		c = noarch.Fgetc(fp)
		return c
	}()) != -1 && c != int('\n') {
		s[func() int {
			defer func() {
				i++
			}()
			return i
		}()] = byte(c)
	}
	s[func() int {
		defer func() {
			i++
		}()
		return i
	}()] = '\x00'
	return (i)
}

// getTime - transpiled function from  $GOPATH/src/github.com/Konstantin8105/History_frame3DD/src/HPGutil.c:235
/*
 * getTime  parse a numeric time string similar to YYYYMMDDhhmmss
 * The input variables y, m, d, hr, mn, sc are the indices of the string s[]
 * which start the YYYY, MM, DD, hh, mm, ss sections of the time string.
 * The corresponding time is returned in "time_t" format.
 */ //
/*  all times are Universal Time never daylight savings time */ //
// normalize t_tm
//	printf("%d ... %s", (int) t_time, ctime(&t_time) );
//
func getTime(s []byte, y int, m int, d int, hr int, mn int, sc int, os int) noarch.TimeT {
	var temp []byte = make([]byte, 16)
	var t_tm noarch.Tm
	var t_time noarch.TimeT
	t_tm.TmYear = noarch.Atoi(noarch.Strncpy(temp, (*(*[1000000000]byte)(unsafe.Pointer(uintptr(unsafe.Pointer(&s[0])) + (uintptr)(y)*unsafe.Sizeof(s[0]))))[:], int(4))) - 1900
	temp[2] = '\x00'
	t_tm.TmMon = noarch.Atoi(noarch.Strncpy(temp, (*(*[1000000000]byte)(unsafe.Pointer(uintptr(unsafe.Pointer(&s[0])) + (uintptr)(m)*unsafe.Sizeof(s[0]))))[:], int(2))) - 1
	t_tm.TmMday = noarch.Atoi(noarch.Strncpy(temp, (*(*[1000000000]byte)(unsafe.Pointer(uintptr(unsafe.Pointer(&s[0])) + (uintptr)(d)*unsafe.Sizeof(s[0]))))[:], int(2)))
	t_tm.TmHour = noarch.Atoi(noarch.Strncpy(temp, (*(*[1000000000]byte)(unsafe.Pointer(uintptr(unsafe.Pointer(&s[0])) + (uintptr)(hr)*unsafe.Sizeof(s[0]))))[:], int(2)))
	t_tm.TmMin = noarch.Atoi(noarch.Strncpy(temp, (*(*[1000000000]byte)(unsafe.Pointer(uintptr(unsafe.Pointer(&s[0])) + (uintptr)(mn)*unsafe.Sizeof(s[0]))))[:], int(2)))
	t_tm.TmSec = noarch.Atoi(noarch.Strncpy(temp, (*(*[1000000000]byte)(unsafe.Pointer(uintptr(unsafe.Pointer(&s[0])) + (uintptr)(sc)*unsafe.Sizeof(s[0]))))[:], int(2))) + os
	t_tm.TmIsdst = -1
	t_time = noarch.Mktime((*[100000000]noarch.Tm)(unsafe.Pointer(&t_tm))[:])
	return noarch.TimeT(t_time)
}

// showProgress - transpiled function from  $GOPATH/src/github.com/Konstantin8105/History_frame3DD/src/HPGutil.c:265
/*
 * SHOW_PROGRESS  -   show the progress of long computations
 */ //
//
func showProgress(i int, n int, count int) {
	var k int
	var j int
	var line_length int = 55
	var percent_done float32
	percent_done = float32((i)) / float32((n))
	j = int(math.Ceil(float64(percent_done * float32(line_length))))
	for k = 1; k <= line_length+13; k++ {
		noarch.Fprintf(stderr, []byte("\b\x00"))
	}
	for k = 1; k < j; k++ {
		noarch.Fprintf(stderr, []byte(">\x00"))
	}
	for k = j; k < line_length; k++ {
		noarch.Fprintf(stderr, []byte(" \x00"))
	}
	noarch.Fprintf(stderr, []byte(" %5.1f%%\x00"), float64(percent_done)*100)
	noarch.Fprintf(stderr, []byte(" %5d\x00"), count)
	noarch.Fflush(stderr)
}

// sferr - transpiled function from  $GOPATH/src/github.com/Konstantin8105/History_frame3DD/src/HPGutil.c:288
/*
 * SFERR  -  Display error message upon an erronous *scanf operation
 */ //
//
func sferr(s []byte) {
	var errMsg []byte = make([]byte, 512)
	noarch.Sprintf(errMsg, []byte(">> Input Data file error while reading %s\n\x00"), s)
	errorMsg(errMsg)
}

type ptrdiff_t int32
type max_align_t struct {
	__clang_max_align_nonce1 int64
	__clang_max_align_nonce2 float64
}

// NRerror - transpiled function from  $GOPATH/src/github.com/Konstantin8105/History_frame3DD/src/NRutil.c:20
/** @file
Memory allocation functions from Numerical Recipes in C, by Press,
Cambridge University Press, 1988
http://www.nr.com/public-domain.html
*/ //
/* Numerical Recipes standard error handler */ //
//
func NRerror(error_text []byte) {
	noarch.Fprintf(stderr, []byte("Numerical Recipes run-time error...\n\x00"))
	noarch.Fprintf(stderr, []byte("%s\n\x00"), error_text)
	noarch.Fprintf(stderr, []byte("...now exiting to system...\n\x00"))
	os.Exit(1000)
}

// vector - transpiled function from  $GOPATH/src/github.com/Konstantin8105/History_frame3DD/src/NRutil.c:29
/* allocate a float vector with subscript range v[nl..nh] */ //
//
func vector(nl int32, nh int32) []float32 {
	var v []float32
	v = make([]float32, size_t((uint32(nh-nl+1+1)*4))/4)
	if v == nil {
		NRerror([]byte("allocation failure in vector()\x00"))
	}
	return (*(*[1000000000]float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&v[0])) - (uintptr)(int(nl))*unsafe.Sizeof(v[0]))))[:][1:]
}

// ivector - transpiled function from  $GOPATH/src/github.com/Konstantin8105/History_frame3DD/src/NRutil.c:39
/* allocate an int vector with subscript range v[nl..nh] */ //
//-nl+NR_END;
//
func ivector(nl int32, nh int32) []int {
	var v []int
	v = make([]int, size_t((uint32(nh-nl+1+1)*4))/4)
	if v == nil {
		NRerror([]byte("allocation failure in ivector()\x00"))
	}
	var e int = int(-nl + 1)
	return (*(*[1000000000]int)(unsafe.Pointer(uintptr(unsafe.Pointer(&v[0])) + (uintptr)(e)*unsafe.Sizeof(v[0]))))[:]
}

// cvector - transpiled function from  $GOPATH/src/github.com/Konstantin8105/History_frame3DD/src/NRutil.c:50
/* allocate an unsigned char vector with subscript range v[nl..nh] */ //
/* return v-nl+NR_END; */ //
//-nl+NR_END;
//
func cvector(nl int32, nh int32) []uint8 {
	var v []uint8
	v = make([]uint8, size_t((uint32(nh-nl+1+1)*1))/1)
	if v == nil {
		NRerror([]byte("allocation failure in cvector()\x00"))
	}
	var e int = int(-nl + 1)
	return (*(*[1000000000]uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(&v[0])) + (uintptr)(e)*unsafe.Sizeof(v[0]))))[:]
}

// lvector - transpiled function from  $GOPATH/src/github.com/Konstantin8105/History_frame3DD/src/NRutil.c:62
/* allocate an unsigned long vector with subscript range v[nl..nh] */ //
//
func lvector(nl int32, nh int32) []uint32 {
	var v []uint32
	v = make([]uint32, size_t((uint32(nh-nl+1+1)*8))/8)
	if v == nil {
		NRerror([]byte("allocation failure in lvector()\x00"))
	}
	return (*(*[1000000000]uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(&v[0])) - (uintptr)(int(nl))*unsafe.Sizeof(v[0]))))[:][1:]
}

// dvector - transpiled function from  $GOPATH/src/github.com/Konstantin8105/History_frame3DD/src/NRutil.c:72
/* allocate a double vector with subscript range v[nl..nh] */ //
// return v-nl+NR_END;
//-nl+NR_END;
//
func dvector(nl int32, nh int32) []float64 {
	var v []float64
	v = make([]float64, size_t((uint32(nh-nl+1+1)*8))/8)
	if v == nil {
		NRerror([]byte("allocation failure in dvector()\x00"))
	}
	var e int = int(-nl + 1)
	return (*(*[1000000000]float64)(unsafe.Pointer(uintptr(unsafe.Pointer(&v[0])) + (uintptr)(e)*unsafe.Sizeof(v[0]))))[:]
}

// matrix - transpiled function from  $GOPATH/src/github.com/Konstantin8105/History_frame3DD/src/NRutil.c:84
// Warning Function `noarch.InterfaceToFloat32SliceSlice` haven`t implementation
/* allocate a float matrix with subscript range m[nrl..nrh][ncl..nch] */ //
/* allocate pointers to rows */ //
/* allocate rows and set pointers to them */ //
/* return pointer to array of pointers to rows */ //
//
func matrix(nrl int32, nrh int32, ncl int32, nch int32) [][]float32 {
	var i int32
	var nrow int32 = nrh - nrl + 1
	var ncol int32 = nch - ncl + 1
	var m [][]float32
	m = make([][]float32, size_t((uint32(nrow+1)*8))/8)
	if m == nil {
		NRerror([]byte("allocation failure 1 in matrix()\x00"))
	}
	m = m[1:]
	m = (*(*[1000000000][]float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&m[0])) - (uintptr)(int(nrl))*unsafe.Sizeof(m[0]))))[:]
	m[nrl] = make([]float32, size_t((uint32(nrow*ncol+1)*4))/4)
	if m[nrl] == nil {
		NRerror([]byte("allocation failure 2 in matrix()\x00"))
	}
	m[nrl] = m[nrl][1:]
	m[nrl] = (*(*[1000000000]float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&m[nrl][0])) - (uintptr)(int(ncl))*unsafe.Sizeof(m[nrl][0]))))[:]
	for i = nrl + 1; i <= nrh; i++ {
		m[i] = (*(*[1000000000]float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&m[i-1][0])) + (uintptr)(int(ncol))*unsafe.Sizeof(m[i-1][0]))))[:]
	}
	return m
}

// dmatrix - transpiled function from  $GOPATH/src/github.com/Konstantin8105/History_frame3DD/src/NRutil.c:108
// Warning Function `noarch.InterfaceToFloat64SliceSlice` haven`t implementation
/* allocate a double matrix with subscript range m[nrl..nrh][ncl..nch] */ //
/* allocate pointers to rows */ //
/* allocate rows and set pointers to them */ //
/* return pointer to array of pointers to rows */ //
//
func dmatrix(nrl int32, nrh int32, ncl int32, nch int32) [][]float64 {
	var i int32
	var nrow int32 = nrh - nrl + 1
	var ncol int32 = nch - ncl + 1
	var m [][]float64
	m = make([][]float64, size_t((uint32(nrow+1)*8))/8)
	if m == nil {
		NRerror([]byte("allocation failure 1 in matrix()\x00"))
	}
	m = m[1:]
	m = (*(*[1000000000][]float64)(unsafe.Pointer(uintptr(unsafe.Pointer(&m[0])) - (uintptr)(int(nrl))*unsafe.Sizeof(m[0]))))[:]
	m[nrl] = make([]float64, size_t((uint32(nrow*ncol+1)*8))/8)
	if m[nrl] == nil {
		NRerror([]byte("allocation failure 2 in matrix()\x00"))
	}
	m[nrl] = m[nrl][1:]
	m[nrl] = (*(*[1000000000]float64)(unsafe.Pointer(uintptr(unsafe.Pointer(&m[nrl][0])) - (uintptr)(int(ncl))*unsafe.Sizeof(m[nrl][0]))))[:]
	for i = nrl + 1; i <= nrh; i++ {
		m[i] = (*(*[1000000000]float64)(unsafe.Pointer(uintptr(unsafe.Pointer(&m[i-1][0])) + (uintptr)(int(ncol))*unsafe.Sizeof(m[i-1][0]))))[:]
	}
	return m
}

// imatrix - transpiled function from  $GOPATH/src/github.com/Konstantin8105/History_frame3DD/src/NRutil.c:132
// Warning Function `noarch.InterfaceToIntSliceSlice` haven`t implementation
/* allocate a int matrix with subscript range m[nrl..nrh][ncl..nch] */ //
/* allocate pointers to rows */ //
/* allocate rows and set pointers to them */ //
/* return pointer to array of pointers to rows */ //
//
func imatrix(nrl int32, nrh int32, ncl int32, nch int32) [][]int {
	var i int32
	var nrow int32 = nrh - nrl + 1
	var ncol int32 = nch - ncl + 1
	var m [][]int
	m = make([][]int, size_t((uint32(nrow+1)*8))/8)
	if m == nil {
		NRerror([]byte("allocation failure 1 in matrix()\x00"))
	}
	m = m[1:]
	m = (*(*[1000000000][]int)(unsafe.Pointer(uintptr(unsafe.Pointer(&m[0])) - (uintptr)(int(nrl))*unsafe.Sizeof(m[0]))))[:]
	m[nrl] = make([]int, size_t((uint32(nrow*ncol+1)*4))/4)
	if m[nrl] == nil {
		NRerror([]byte("allocation failure 2 in matrix()\x00"))
	}
	m[nrl] = m[nrl][1:]
	m[nrl] = (*(*[1000000000]int)(unsafe.Pointer(uintptr(unsafe.Pointer(&m[nrl][0])) - (uintptr)(int(ncl))*unsafe.Sizeof(m[nrl][0]))))[:]
	for i = nrl + 1; i <= nrh; i++ {
		m[i] = (*(*[1000000000]int)(unsafe.Pointer(uintptr(unsafe.Pointer(&m[i-1][0])) + (uintptr)(int(ncol))*unsafe.Sizeof(m[i-1][0]))))[:]
	}
	return m
}

// subMatrix - transpiled function from  $GOPATH/src/github.com/Konstantin8105/History_frame3DD/src/NRutil.c:157
// Warning Function `noarch.InterfaceToFloat32SliceSlice` haven`t implementation
/* point a subMatrix [newrl..][newcl..] to a[oldrl..oldrh][oldcl..oldch] */ //
/* allocate array of pointers to rows */ //
/* set pointers to rows */ //
/* return pointer to array of pointers to rows */ //
//
func subMatrix(a [][]float32, oldrl int32, oldrh int32, oldcl int32, oldch int32, newrl int32, newcl int32) [][]float32 {
	var i int32
	var j int32
	var nrow int32 = oldrh - oldrl + 1
	var ncol int32 = oldcl - newcl
	var m [][]float32
	m = make([][]float32, size_t((uint32(nrow+1)*8))/8)
	if m == nil {
		NRerror([]byte("allocation failure in subMatrix()\x00"))
	}
	m = m[1:]
	m = (*(*[1000000000][]float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&m[0])) - (uintptr)(int(newrl))*unsafe.Sizeof(m[0]))))[:]
	{
		i = oldrl
		j = newrl
		for j = newrl; i <= oldrh; {
			m[j] = (*(*[1000000000]float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&a[i][0])) + (uintptr)(int(ncol))*unsafe.Sizeof(a[i][0]))))[:]
			i++
			j++
		}
	}
	return m
}

// convert_matrix - transpiled function from  $GOPATH/src/github.com/Konstantin8105/History_frame3DD/src/NRutil.c:177
/* allocate a float matrix m[nrl..nrh][ncl..nch] that points to the matrix
declared in the standard C manner as a[nrow][ncol], where nrow=nrh-nrl+1
and ncol=nch-ncl+1. The routine should be called with the address
&a[0][0] as the first argument. */ //
/* allocate pointers to rows */ //
/* set pointers to rows */ //
/* return pointer to array of pointers to rows */ //
//
func convert_matrix(a []float32, nrl int32, nrh int32, ncl int32, nch int32) [][]float32 {
	var i int32
	var j int32
	var nrow int32 = nrh - nrl + 1
	var ncol int32 = nch - ncl + 1
	var m [][]float32
	m = make([][]float32, size_t((uint32(nrow+1)*8))/8)
	if m == nil {
		NRerror([]byte("allocation failure in convert_matrix()\x00"))
	}
	m = m[1:]
	m = (*(*[1000000000][]float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&m[0])) - (uintptr)(int(nrl))*unsafe.Sizeof(m[0]))))[:]
	m[nrl] = (*(*[1000000000]float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&a[0])) - (uintptr)(int(ncl))*unsafe.Sizeof(a[0]))))[:]
	{
		i = 1
		j = nrl + 1
		for j = nrl + 1; i < nrow; {
			m[j] = (*(*[1000000000]float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&m[j-1][0])) + (uintptr)(int(ncol))*unsafe.Sizeof(m[j-1][0]))))[:]
			i++
			j++
		}
	}
	return m
}

// f3tensor - transpiled function from  $GOPATH/src/github.com/Konstantin8105/History_frame3DD/src/NRutil.c:199
/* allocate a float 3tensor with range t[nrl..nrh][ncl..nch][ndl..ndh] */ //
/* allocate pointers to pointers to rows */ //
/* allocate pointers to rows and set pointers to them */ //
/* allocate rows and set pointers to them */ //
/* return pointer to array of pointers to rows */ //
//
func f3tensor(nrl int32, nrh int32, ncl int32, nch int32, ndl int32, ndh int32) [][][]float32 {
	var i int32
	var j int32
	var nrow int32 = nrh - nrl + 1
	var ncol int32 = nch - ncl + 1
	var ndep int32 = ndh - ndl + 1
	var t [][][]float32
	t = make([][][]float32, size_t((uint32(nrow+1)*8))/8)
	if t == nil {
		NRerror([]byte("allocation failure 1 in f3tensor()\x00"))
	}
	t = t[1:]
	t = (*(*[1000000000][][]float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&t[0])) - (uintptr)(int(nrl))*unsafe.Sizeof(t[0]))))[:]
	t[nrl] = make([][]float32, size_t((uint32(nrow*ncol+1)*8))/8)
	if t[nrl] == nil {
		NRerror([]byte("allocation failure 2 in f3tensor()\x00"))
	}
	t[nrl] = t[nrl][1:]
	t[nrl] = (*(*[1000000000][]float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&t[nrl][0])) - (uintptr)(int(ncl))*unsafe.Sizeof(t[nrl][0]))))[:]
	t[nrl][ncl] = make([]float32, size_t((uint32(nrow*ncol*ndep+1)*4))/4)
	if t[nrl][ncl] == nil {
		NRerror([]byte("allocation failure 3 in f3tensor()\x00"))
	}
	t[nrl][ncl] = t[nrl][ncl][1:]
	t[nrl][ncl] = (*(*[1000000000]float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&t[nrl][ncl][0])) - (uintptr)(int(ndl))*unsafe.Sizeof(t[nrl][ncl][0]))))[:]
	for j = ncl + 1; j <= nch; j++ {
		t[nrl][j] = (*(*[1000000000]float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&t[nrl][j-1][0])) + (uintptr)(int(ndep))*unsafe.Sizeof(t[nrl][j-1][0]))))[:]
	}
	for i = nrl + 1; i <= nrh; i++ {
		t[i] = (*(*[1000000000][]float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&t[i-1][0])) + (uintptr)(int(ncol))*unsafe.Sizeof(t[i-1][0]))))[:]
		t[i][ncl] = (*(*[1000000000]float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&t[i-1][ncl][0])) + (uintptr)(int(ncol*ndep))*unsafe.Sizeof(t[i-1][ncl][0]))))[:]
		for j = ncl + 1; j <= nch; j++ {
			t[i][j] = (*(*[1000000000]float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&t[i][j-1][0])) + (uintptr)(int(ndep))*unsafe.Sizeof(t[i][j-1][0]))))[:]
		}
	}
	return t
}

// free_vector - transpiled function from  $GOPATH/src/github.com/Konstantin8105/History_frame3DD/src/NRutil.c:234
/* free a float vector allocated with vector() */ //
//
func free_vector(v []float32, nl int32, nh int32) {
	_ = ((*(*[1000000000]float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&(*(*[1000000000]float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&v[0])) + (uintptr)(int(nl))*unsafe.Sizeof(v[0]))))[:][0])) - (uintptr)(1)*unsafe.Sizeof((*(*[1000000000]float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&v[0])) + (uintptr)(int(nl))*unsafe.Sizeof(v[0]))))[:][0]))))[:])
}

// free_ivector - transpiled function from  $GOPATH/src/github.com/Konstantin8105/History_frame3DD/src/NRutil.c:240
/* free an int vector allocated with ivector() */ //
//
func free_ivector(v []int, nl int32, nh int32) {
	_ = ((*(*[1000000000]int)(unsafe.Pointer(uintptr(unsafe.Pointer(&(*(*[1000000000]int)(unsafe.Pointer(uintptr(unsafe.Pointer(&v[0])) + (uintptr)(int(nl))*unsafe.Sizeof(v[0]))))[:][0])) - (uintptr)(1)*unsafe.Sizeof((*(*[1000000000]int)(unsafe.Pointer(uintptr(unsafe.Pointer(&v[0])) + (uintptr)(int(nl))*unsafe.Sizeof(v[0]))))[:][0]))))[:])
}

// free_cvector - transpiled function from  $GOPATH/src/github.com/Konstantin8105/History_frame3DD/src/NRutil.c:246
/* free an unsigned char vector allocated with cvector() */ //
//
func free_cvector(v []uint8, nl int32, nh int32) {
	_ = ((*(*[1000000000]uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(&(*(*[1000000000]uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(&v[0])) + (uintptr)(int(nl))*unsafe.Sizeof(v[0]))))[:][0])) - (uintptr)(1)*unsafe.Sizeof((*(*[1000000000]uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(&v[0])) + (uintptr)(int(nl))*unsafe.Sizeof(v[0]))))[:][0]))))[:])
}

// free_lvector - transpiled function from  $GOPATH/src/github.com/Konstantin8105/History_frame3DD/src/NRutil.c:252
/* free an unsigned long vector allocated with lvector() */ //
//
func free_lvector(v []uint32, nl int32, nh int32) {
	_ = ((*(*[1000000000]uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(&(*(*[1000000000]uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(&v[0])) + (uintptr)(int(nl))*unsafe.Sizeof(v[0]))))[:][0])) - (uintptr)(1)*unsafe.Sizeof((*(*[1000000000]uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(&v[0])) + (uintptr)(int(nl))*unsafe.Sizeof(v[0]))))[:][0]))))[:])
}

// free_dvector - transpiled function from  $GOPATH/src/github.com/Konstantin8105/History_frame3DD/src/NRutil.c:258
/* free a double vector allocated with dvector() */ //
//
func free_dvector(v []float64, nl int32, nh int32) {
	_ = ((*(*[1000000000]float64)(unsafe.Pointer(uintptr(unsafe.Pointer(&(*(*[1000000000]float64)(unsafe.Pointer(uintptr(unsafe.Pointer(&v[0])) + (uintptr)(int(nl))*unsafe.Sizeof(v[0]))))[:][0])) - (uintptr)(1)*unsafe.Sizeof((*(*[1000000000]float64)(unsafe.Pointer(uintptr(unsafe.Pointer(&v[0])) + (uintptr)(int(nl))*unsafe.Sizeof(v[0]))))[:][0]))))[:])
}

// free_matrix - transpiled function from  $GOPATH/src/github.com/Konstantin8105/History_frame3DD/src/NRutil.c:264
/* free a float matrix allocated by matrix() */ //
//
func free_matrix(m [][]float32, nrl int32, nrh int32, ncl int32, nch int32) {
	_ = ((*(*[1000000000]float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&(*(*[1000000000]float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&m[nrl][0])) + (uintptr)(int(ncl))*unsafe.Sizeof(m[nrl][0]))))[:][0])) - (uintptr)(1)*unsafe.Sizeof((*(*[1000000000]float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&m[nrl][0])) + (uintptr)(int(ncl))*unsafe.Sizeof(m[nrl][0]))))[:][0]))))[:])
	_ = ((*(*[1000000000][]float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&(*(*[1000000000][]float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&m[0])) + (uintptr)(int(nrl))*unsafe.Sizeof(m[0]))))[:][0])) - (uintptr)(1)*unsafe.Sizeof((*(*[1000000000][]float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&m[0])) + (uintptr)(int(nrl))*unsafe.Sizeof(m[0]))))[:][0]))))[:])
}

// free_dmatrix - transpiled function from  $GOPATH/src/github.com/Konstantin8105/History_frame3DD/src/NRutil.c:271
/* free a double matrix allocated by dmatrix() */ //
//
func free_dmatrix(m [][]float64, nrl int32, nrh int32, ncl int32, nch int32) {
	_ = ((*(*[1000000000]float64)(unsafe.Pointer(uintptr(unsafe.Pointer(&(*(*[1000000000]float64)(unsafe.Pointer(uintptr(unsafe.Pointer(&m[nrl][0])) + (uintptr)(int(ncl))*unsafe.Sizeof(m[nrl][0]))))[:][0])) - (uintptr)(1)*unsafe.Sizeof((*(*[1000000000]float64)(unsafe.Pointer(uintptr(unsafe.Pointer(&m[nrl][0])) + (uintptr)(int(ncl))*unsafe.Sizeof(m[nrl][0]))))[:][0]))))[:])
	_ = ((*(*[1000000000][]float64)(unsafe.Pointer(uintptr(unsafe.Pointer(&(*(*[1000000000][]float64)(unsafe.Pointer(uintptr(unsafe.Pointer(&m[0])) + (uintptr)(int(nrl))*unsafe.Sizeof(m[0]))))[:][0])) - (uintptr)(1)*unsafe.Sizeof((*(*[1000000000][]float64)(unsafe.Pointer(uintptr(unsafe.Pointer(&m[0])) + (uintptr)(int(nrl))*unsafe.Sizeof(m[0]))))[:][0]))))[:])
}

// free_imatrix - transpiled function from  $GOPATH/src/github.com/Konstantin8105/History_frame3DD/src/NRutil.c:278
/* free an int matrix allocated by imatrix() */ //
//
func free_imatrix(m [][]int, nrl int32, nrh int32, ncl int32, nch int32) {
	_ = ((*(*[1000000000]int)(unsafe.Pointer(uintptr(unsafe.Pointer(&(*(*[1000000000]int)(unsafe.Pointer(uintptr(unsafe.Pointer(&m[nrl][0])) + (uintptr)(int(ncl))*unsafe.Sizeof(m[nrl][0]))))[:][0])) - (uintptr)(1)*unsafe.Sizeof((*(*[1000000000]int)(unsafe.Pointer(uintptr(unsafe.Pointer(&m[nrl][0])) + (uintptr)(int(ncl))*unsafe.Sizeof(m[nrl][0]))))[:][0]))))[:])
	_ = ((*(*[1000000000][]int)(unsafe.Pointer(uintptr(unsafe.Pointer(&(*(*[1000000000][]int)(unsafe.Pointer(uintptr(unsafe.Pointer(&m[0])) + (uintptr)(int(nrl))*unsafe.Sizeof(m[0]))))[:][0])) - (uintptr)(1)*unsafe.Sizeof((*(*[1000000000][]int)(unsafe.Pointer(uintptr(unsafe.Pointer(&m[0])) + (uintptr)(int(nrl))*unsafe.Sizeof(m[0]))))[:][0]))))[:])
}

// free_subMatrix - transpiled function from  $GOPATH/src/github.com/Konstantin8105/History_frame3DD/src/NRutil.c:285
/* free a subMatrix allocated by subMatrix() */ //
//
func free_subMatrix(b [][]float32, nrl int32, nrh int32, ncl int32, nch int32) {
	_ = ((*(*[1000000000][]float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&(*(*[1000000000][]float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&b[0])) + (uintptr)(int(nrl))*unsafe.Sizeof(b[0]))))[:][0])) - (uintptr)(1)*unsafe.Sizeof((*(*[1000000000][]float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&b[0])) + (uintptr)(int(nrl))*unsafe.Sizeof(b[0]))))[:][0]))))[:])
}

// free_convert_matrix - transpiled function from  $GOPATH/src/github.com/Konstantin8105/History_frame3DD/src/NRutil.c:291
/* free a matrix allocated by convert_matrix() */ //
//
func free_convert_matrix(b [][]float32, nrl int32, nrh int32, ncl int32, nch int32) {
	_ = ((*(*[1000000000][]float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&(*(*[1000000000][]float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&b[0])) + (uintptr)(int(nrl))*unsafe.Sizeof(b[0]))))[:][0])) - (uintptr)(1)*unsafe.Sizeof((*(*[1000000000][]float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&b[0])) + (uintptr)(int(nrl))*unsafe.Sizeof(b[0]))))[:][0]))))[:])
}

// free_f3tensor - transpiled function from  $GOPATH/src/github.com/Konstantin8105/History_frame3DD/src/NRutil.c:297
/* free a float f3tensor allocated by f3tensor() */ //
//
func free_f3tensor(t [][][]float32, nrl int32, nrh int32, ncl int32, nch int32, ndl int32, ndh int32) {
	_ = ((*(*[1000000000]float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&(*(*[1000000000]float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&t[nrl][ncl][0])) + (uintptr)(int(ndl))*unsafe.Sizeof(t[nrl][ncl][0]))))[:][0])) - (uintptr)(1)*unsafe.Sizeof((*(*[1000000000]float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&t[nrl][ncl][0])) + (uintptr)(int(ndl))*unsafe.Sizeof(t[nrl][ncl][0]))))[:][0]))))[:])
	_ = ((*(*[1000000000][]float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&(*(*[1000000000][]float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&t[nrl][0])) + (uintptr)(int(ncl))*unsafe.Sizeof(t[nrl][0]))))[:][0])) - (uintptr)(1)*unsafe.Sizeof((*(*[1000000000][]float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&t[nrl][0])) + (uintptr)(int(ncl))*unsafe.Sizeof(t[nrl][0]))))[:][0]))))[:])
	_ = ((*(*[1000000000][][]float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&(*(*[1000000000][][]float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&t[0])) + (uintptr)(int(nrl))*unsafe.Sizeof(t[0]))))[:][0])) - (uintptr)(1)*unsafe.Sizeof((*(*[1000000000][][]float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&t[0])) + (uintptr)(int(nrl))*unsafe.Sizeof(t[0]))))[:][0]))))[:])
}

// Cvector - transpiled function from  $GOPATH/src/github.com/Konstantin8105/History_frame3DD/src/NRutil.c:305
/* allocate storage for a complex vector	*/ //
//
func Cvector(nl int, nh int) []fcomplex {
	var v []fcomplex
	v = make([]fcomplex, uint32((nh-nl+1))*8/8)
	if v == nil {
		NRerror([]byte("allocation failure in Cvector()\x00"))
	}
	return (*(*[1000000000]fcomplex)(unsafe.Pointer(uintptr(unsafe.Pointer(&v[0])) - (uintptr)(nl)*unsafe.Sizeof(v[0]))))[:]
}

// Cmatrix - transpiled function from  $GOPATH/src/github.com/Konstantin8105/History_frame3DD/src/NRutil.c:316
// Warning Function `noarch.InterfaceToFcomplexSliceSlice` haven`t implementation
/* allocate storage for a Complex matrix	*/ //
//
func Cmatrix(nrl int, nrh int, ncl int, nch int) [][]fcomplex {
	var i int
	var m [][]fcomplex
	m = make([][]fcomplex, uint32((nrh-nrl+1))*8/8)
	if m == nil {
		NRerror([]byte("allocation failure 1 in Cmatrix()\x00"))
	}
	m = (*(*[1000000000][]fcomplex)(unsafe.Pointer(uintptr(unsafe.Pointer(&m[0])) - (uintptr)(nrl)*unsafe.Sizeof(m[0]))))[:]
	for i = nrl; i <= nrh; i++ {
		m[i] = make([]fcomplex, uint32((nch-ncl+1))*8/8)
		if m[i] == nil {
			NRerror([]byte("allocation failure 2 in Cmatrix()\x00"))
		}
		m[i] = (*(*[1000000000]fcomplex)(unsafe.Pointer(uintptr(unsafe.Pointer(&m[i][0])) - (uintptr)(ncl)*unsafe.Sizeof(m[i][0]))))[:]
	}
	return m
}

// D3matrix - transpiled function from  $GOPATH/src/github.com/Konstantin8105/History_frame3DD/src/NRutil.c:334
// Warning Function `noarch.Float64SliceSliceSliceToFloat32SliceSliceSlice` haven`t implementation
// Warning Function `noarch.InterfaceToFloat64SliceSlice` haven`t implementation
// Warning Function `noarch.Float64SliceSliceToFloat32SliceSlice` haven`t implementation
// Warning Function `noarch.InterfaceToFloat64SliceSlice` haven`t implementation
// Warning Function `noarch.Float64SliceSliceToFloat32SliceSlice` haven`t implementation
// Warning Function `noarch.Float64SliceToFloat32Slice` haven`t implementation
/* storage for a 3-D matrix */ //
/*
	m=(float ***) malloc((unsigned) (nrh-nrl+1)*sizeof(float*));
	if (!m) NRerror("alloc failure 1 in 3Dmatrix()");
	m -= nrl;
	for (i=nrl;i<=nrh;i++) {
		m[i]=(float **) malloc((unsigned) (nch-ncl+1)*sizeof(float*));
		if (!m[i]) NRerror("alloc failure 2 in 3Dmatrix()");
		m[i] -= ncl;
		for (j=ncl;j<=nch;j++) {
			m[i][j]=(float *)
				malloc((unsigned) (nzh-nzl+1)*sizeof(float));
			if (!m[i][j]) NRerror("alloc failure 3 in 3Dmatrix()");
			m[i][j] -= nzl;
		}
	}
*/ //
//
func D3matrix(nrl int, nrh int, ncl int, nch int, nzl int, nzh int) [][][]float32 {
	var i int
	var j int
	var m [][][]float32
	m = make([][][]float32, uint32((nrh+1))*8/8)
	for i = 0; i <= nrh; i++ {
		m[i] = make([][]float32, uint32((nch+1))*8/8)
		for j = 0; j <= nch; j++ {
			m[i][j] = make([]float32, uint32((nzh+1))*8/4)
		}
	}
	return m
}

// D3dmatrix - transpiled function from  $GOPATH/src/github.com/Konstantin8105/History_frame3DD/src/NRutil.c:367
// Warning Function `noarch.InterfaceToFloat64SliceSlice` haven`t implementation
/* storage for a 3-D matrix */ //
/*
	m=(double ***) malloc((unsigned) (nrh-nrl+1)*sizeof(double*));
	if (!m) NRerror("alloc failure 1 in 3Ddmatrix()");
	m -= nrl;
	for (i=nrl;i<=nrh;i++) {
		m[i]=(double **) malloc((unsigned) (nch-ncl+1)*sizeof(double*));
		if (!m[i]) NRerror("alloc failure 2 in 3Dmatrix()");
		m[i] -= ncl;
		for (j=ncl;j<=nch;j++) {
			m[i][j]=(double *)
				malloc((unsigned) (nzh-nzl+1)*sizeof(double));
			if (!m[i][j]) NRerror("alloc failure 3 in 3Ddmatrix()");
			m[i][j] -= nzl;
		}
	}
*/ //
//
func D3dmatrix(nrl int, nrh int, ncl int, nch int, nzl int, nzh int) [][][]float64 {
	var i int
	var j int
	var m [][][]float64
	m = make([][][]float64, uint32((nrh+1))*8/8)
	for i = 0; i <= nrh; i++ {
		m[i] = make([][]float64, uint32((nch+1))*8/8)
		for j = 0; j <= nch; j++ {
			m[i][j] = make([]float64, uint32((nzh+1))*8/8)
		}
	}
	return m
}

// free_Cvector - transpiled function from  $GOPATH/src/github.com/Konstantin8105/History_frame3DD/src/NRutil.c:401
/* free((void*)#<{(|(char*)|)}># (v+nl)); */ //
//
func free_Cvector(v []fcomplex, nl int, nh int) {
}

// free_Cmatrix - transpiled function from  $GOPATH/src/github.com/Konstantin8105/History_frame3DD/src/NRutil.c:406
/* int	i; */ //
/*  */ //
/* for(i=nrh;i>=nrl;i--) free((void*)#<{(|(char*)|)}># (m[i]+ncl)); */ //
/* free((void*)#<{(|(char*)|)}># (m+nrl)); */ //
//
func free_Cmatrix(m [][]fcomplex, nrl int, nrh int, ncl int, nch int) {
}

// free_D3matrix - transpiled function from  $GOPATH/src/github.com/Konstantin8105/History_frame3DD/src/NRutil.c:414
/* int     i,j; */ //
/*  */ //
/* for(i=nrh;i>=nrl;i--) { */ //
/* 	for(j=nch;j>=ncl;j--) { */ //
/* 		free((void*)#<{(|(char*)|)}># (m[i][j]+nzl)); */ //
/* 	} */ //
/* } */ //
//
func free_D3matrix(m [][][]float32, nrl int, nrh int, ncl int, nch int, nzl int, nzh int) {
}

// free_D3dmatrix - transpiled function from  $GOPATH/src/github.com/Konstantin8105/History_frame3DD/src/NRutil.c:425
/* int     i,j; */ //
/*  */ //
/* for(i=nrh;i>=nrl;i--) { */ //
/* 	for(j=nch;j>=ncl;j--) { */ //
/* 		free((void*)#<{(|(char*)|)}># (m[i][j]+nzl)); */ //
/* 	} */ //
/* } */ //
//
func free_D3dmatrix(m [][][]float64, nrl int, nrh int, ncl int, nch int, nzl int, nzh int) {
}

// show_vector - transpiled function from  $GOPATH/src/github.com/Konstantin8105/History_frame3DD/src/NRutil.c:768
/*
 * SHOW_VECTOR  -  display a vector of dimension [1..n]
 */ //
//
func show_vector(A []float32, n int) {
	var j int
	for j = 1; j <= n; j++ {
		if A[j] != 0 {
			noarch.Fprintf(stdout, []byte("%14.6e\x00"), float64(A[j]))
		} else {
			noarch.Fprintf(stdout, []byte("   0       \x00"))
		}
	}
	noarch.Fprintf(stdout, []byte(" ]';\n\n\x00"))
}

// show_dvector - transpiled function from  $GOPATH/src/github.com/Konstantin8105/History_frame3DD/src/NRutil.c:783
/*
 * SHOW_DVECTOR  -  display a vector of dimension [1..n]
 */ //
//
func show_dvector(A []float64, n int) {
	var j int
	for j = 1; j <= n; j++ {
		if math.Abs(A[j]) >= 1e-99 {
			noarch.Fprintf(stdout, []byte("%14.6e\x00"), A[j])
		} else {
			noarch.Fprintf(stdout, []byte("   0       \x00"))
		}
	}
	noarch.Fprintf(stdout, []byte(" ]';\n\n\x00"))
}

// show_ivector - transpiled function from  $GOPATH/src/github.com/Konstantin8105/History_frame3DD/src/NRutil.c:799
/*
 * SHOW_IVECTOR  -  display a vector of integers of dimension [1..n]
 */ //
//
func show_ivector(A []int, n int) {
	var j int
	for j = 1; j <= n; j++ {
		if A[j] != 0 {
			noarch.Fprintf(stdout, []byte("%11d\x00"), A[j])
		} else {
			noarch.Fprintf(stdout, []byte("   0       \x00"))
		}
	}
	noarch.Fprintf(stdout, []byte(" ]';\n\n\x00"))
}

// show_matrix - transpiled function from  $GOPATH/src/github.com/Konstantin8105/History_frame3DD/src/NRutil.c:815
/*
 * SHOW_MATRIX  -  display a matrix of dimension [1..m][1..n]
 */ //
//
func show_matrix(A [][]float32, m int, n int) {
	var i int
	var j int
	for i = 1; i <= m; i++ {
		for j = 1; j <= n; j++ {
			if A[i][j] != 0 {
				noarch.Fprintf(stdout, []byte("%14.6e\x00"), float64(A[i][j]))
			} else {
				noarch.Fprintf(stdout, []byte("   0       \x00"))
			}
		}
		if i == m {
			noarch.Fprintf(stdout, []byte(" ];\n\n\x00"))
		} else {
			noarch.Fprintf(stdout, []byte(" \n\x00"))
		}
	}
}

// show_dmatrix - transpiled function from  $GOPATH/src/github.com/Konstantin8105/History_frame3DD/src/NRutil.c:833
/*
 * SHOW_DMATRIX  - display a matrix of dimension [1..m][1..n]
 */ //
//
func show_dmatrix(A [][]float64, m int, n int) {
	var i int
	var j int
	for i = 1; i <= m; i++ {
		for j = 1; j <= n; j++ {
			if math.Abs(A[i][j]) > 1e-99 {
				noarch.Fprintf(stdout, []byte("%11.3e\x00"), A[i][j])
			} else {
				noarch.Fprintf(stdout, []byte("   0       \x00"))
			}
		}
		if i == m {
			noarch.Fprintf(stdout, []byte(" ];\n\n\x00"))
		} else {
			noarch.Fprintf(stdout, []byte(" \n\x00"))
		}
	}
}

// save_vector - transpiled function from  $GOPATH/src/github.com/Konstantin8105/History_frame3DD/src/NRutil.c:852
/*
 * SAVE_VECTOR  -  save a vector of dimension [1..n] to the named file
 */ //
//
func save_vector(filename []byte, V []float32, nl int, nh int, mode []byte) {
	var fp_v *noarch.File
	var i int
	var now noarch.TimeT
	if (func() *noarch.File {
		fp_v = noarch.Fopen(filename, mode)
		return fp_v
	}()) == nil {
		noarch.Printf([]byte(" error: cannot open file: '%s' \n\x00"), filename)
		os.Exit(1011)
	}
	_ = noarch.Time((*[100000000]noarch.TimeT)(unsafe.Pointer(&now))[:])
	noarch.Fprintf(fp_v, []byte("%% filename: %s - %s\x00"), filename, noarch.Ctime((*[100000000]noarch.TimeT)(unsafe.Pointer(&now))[:]))
	noarch.Fprintf(fp_v, []byte("%% type: vector\n\x00"))
	noarch.Fprintf(fp_v, []byte("%% rows: %d\n\x00"), 1)
	noarch.Fprintf(fp_v, []byte("%% columns: %d\n\x00"), nh-nl+1)
	for i = nl; i <= nh; i++ {
		if V[i] != 0 {
			noarch.Fprintf(fp_v, []byte("%15.6e\x00"), float64(V[i]))
		} else {
			noarch.Fprintf(fp_v, []byte("    0         \x00"))
		}
		noarch.Fprintf(fp_v, []byte("\n\x00"))
	}
	noarch.Fclose(fp_v)
}

// save_dvector - transpiled function from  $GOPATH/src/github.com/Konstantin8105/History_frame3DD/src/NRutil.c:880
/*
 * SAVE_DVECTOR  -  save a vector of dimension [1..n] to the named file
 */ //
//
func save_dvector(filename []byte, V []float64, nl int, nh int, mode []byte) {
	var fp_v *noarch.File
	var i int
	var now noarch.TimeT
	if (func() *noarch.File {
		fp_v = noarch.Fopen(filename, mode)
		return fp_v
	}()) == nil {
		noarch.Printf([]byte(" error: cannot open file: '%s' \n\x00"), filename)
		os.Exit(1011)
	}
	_ = noarch.Time((*[100000000]noarch.TimeT)(unsafe.Pointer(&now))[:])
	noarch.Fprintf(fp_v, []byte("%% filename: %s - %s\x00"), filename, noarch.Ctime((*[100000000]noarch.TimeT)(unsafe.Pointer(&now))[:]))
	noarch.Fprintf(fp_v, []byte("%% type: vector\n\x00"))
	noarch.Fprintf(fp_v, []byte("%% rows: %d\n\x00"), 1)
	noarch.Fprintf(fp_v, []byte("%% columns: %d\n\x00"), nh-nl+1)
	for i = nl; i <= nh; i++ {
		if V[i] != 0 {
			noarch.Fprintf(fp_v, []byte("%21.12e\x00"), V[i])
		} else {
			noarch.Fprintf(fp_v, []byte("    0                \x00"))
		}
		noarch.Fprintf(fp_v, []byte("\n\x00"))
	}
	noarch.Fclose(fp_v)
}

// save_ivector - transpiled function from  $GOPATH/src/github.com/Konstantin8105/History_frame3DD/src/NRutil.c:908
/*
 * SAVE_IVECTOR  -  save an integer vector of dimension [1..n] to the named file
 */ //
//
func save_ivector(filename []byte, V []int, nl int, nh int, mode []byte) {
	var fp_v *noarch.File
	var i int
	var now noarch.TimeT
	if (func() *noarch.File {
		fp_v = noarch.Fopen(filename, mode)
		return fp_v
	}()) == nil {
		noarch.Printf([]byte(" error: cannot open file: '%s' \n\x00"), filename)
		os.Exit(1012)
	}
	_ = noarch.Time((*[100000000]noarch.TimeT)(unsafe.Pointer(&now))[:])
	noarch.Fprintf(fp_v, []byte("%% filename: %s - %s\x00"), filename, noarch.Ctime((*[100000000]noarch.TimeT)(unsafe.Pointer(&now))[:]))
	noarch.Fprintf(fp_v, []byte("%% type: vector\n\x00"))
	noarch.Fprintf(fp_v, []byte("%% rows: %d\n\x00"), 1)
	noarch.Fprintf(fp_v, []byte("%% columns: %d\n\x00"), nh-nl+1)
	for i = nl; i <= nh; i++ {
		if V[i] != 0 {
			noarch.Fprintf(fp_v, []byte("%15d\x00"), V[i])
		} else {
			noarch.Fprintf(fp_v, []byte("   0         \x00"))
		}
		noarch.Fprintf(fp_v, []byte("\n\x00"))
	}
	noarch.Fclose(fp_v)
}

// save_matrix - transpiled function from  $GOPATH/src/github.com/Konstantin8105/History_frame3DD/src/NRutil.c:936
/*
 * SAVE_MATRIX  -  save a matrix of dimension [ml..mh][nl..nh] to the named file
 */ //
//
func save_matrix(filename []byte, A [][]float32, ml int, mh int, nl int, nh int, transpose int, mode []byte) {
	var fp_m *noarch.File
	var i int
	var j int
	var rows int
	var cols int
	var now noarch.TimeT
	if transpose != 0 {
		rows = nh - nl + 1
	} else {
		rows = mh - ml + 1
	}
	if transpose != 0 {
		cols = mh - ml + 1
	} else {
		cols = nh - nl + 1
	}
	if (func() *noarch.File {
		fp_m = noarch.Fopen(filename, mode)
		return fp_m
	}()) == nil {
		noarch.Printf([]byte(" error: cannot open file: %s \n\x00"), filename)
		os.Exit(1013)
	}
	_ = noarch.Time((*[100000000]noarch.TimeT)(unsafe.Pointer(&now))[:])
	noarch.Fprintf(fp_m, []byte("%% filename: %s - %s\x00"), filename, noarch.Ctime((*[100000000]noarch.TimeT)(unsafe.Pointer(&now))[:]))
	noarch.Fprintf(fp_m, []byte("%% type: matrix \n\x00"))
	noarch.Fprintf(fp_m, []byte("%% rows: %d\n\x00"), rows)
	noarch.Fprintf(fp_m, []byte("%% columns: %d\n\x00"), cols)
	if transpose != 0 {
		for j = nl; j <= nh; j++ {
			for i = ml; i <= mh; i++ {
				if A[i][j] != 0 {
					noarch.Fprintf(fp_m, []byte("%15.6e\x00"), float64(A[i][j]))
				} else {
					noarch.Fprintf(fp_m, []byte("    0          \x00"))
				}
			}
			noarch.Fprintf(fp_m, []byte("\n\x00"))
		}
	} else {
		for i = ml; i <= mh; i++ {
			for j = nl; j <= nh; j++ {
				if A[i][j] != 0 {
					noarch.Fprintf(fp_m, []byte("%15.6e\x00"), float64(A[i][j]))
				} else {
					noarch.Fprintf(fp_m, []byte("    0          \x00"))
				}
			}
			noarch.Fprintf(fp_m, []byte("\n\x00"))
		}
	}
	noarch.Fclose(fp_m)
}

// save_dmatrix - transpiled function from  $GOPATH/src/github.com/Konstantin8105/History_frame3DD/src/NRutil.c:979
/*
 * SAVE_DMATRIX  - save a matrix of dimension [ml..mh][nl..nh] to the named file
 */ //
//
func save_dmatrix(filename []byte, A [][]float64, ml int, mh int, nl int, nh int, transpose int, mode []byte) {
	var fp_m *noarch.File
	var i int
	var j int
	var rows int
	var cols int
	var now noarch.TimeT
	if transpose != 0 {
		rows = nh - nl + 1
	} else {
		rows = mh - ml + 1
	}
	if transpose != 0 {
		cols = mh - ml + 1
	} else {
		cols = nh - nl + 1
	}
	if (func() *noarch.File {
		fp_m = noarch.Fopen(filename, mode)
		return fp_m
	}()) == nil {
		noarch.Printf([]byte(" error: cannot open file: %s \n\x00"), filename)
		os.Exit(1014)
	}
	_ = noarch.Time((*[100000000]noarch.TimeT)(unsafe.Pointer(&now))[:])
	noarch.Fprintf(fp_m, []byte("%% filename: %s - %s\x00"), filename, noarch.Ctime((*[100000000]noarch.TimeT)(unsafe.Pointer(&now))[:]))
	noarch.Fprintf(fp_m, []byte("%% type: matrix \n\x00"))
	noarch.Fprintf(fp_m, []byte("%% rows: %d\n\x00"), rows)
	noarch.Fprintf(fp_m, []byte("%% columns: %d\n\x00"), cols)
	if transpose != 0 {
		for j = nl; j <= nh; j++ {
			for i = ml; i <= mh; i++ {
				if math.Abs(A[i][j]) > 1e-99 {
					noarch.Fprintf(fp_m, []byte("%21.12e\x00"), A[i][j])
				} else {
					noarch.Fprintf(fp_m, []byte("    0                \x00"))
				}
			}
			noarch.Fprintf(fp_m, []byte("\n\x00"))
		}
	} else {
		for i = ml; i <= mh; i++ {
			for j = nl; j <= nh; j++ {
				if math.Abs(A[i][j]) > 1e-99 {
					noarch.Fprintf(fp_m, []byte("%21.12e\x00"), A[i][j])
				} else {
					noarch.Fprintf(fp_m, []byte("    0                \x00"))
				}
			}
			noarch.Fprintf(fp_m, []byte("\n\x00"))
		}
	}
	noarch.Fclose(fp_m)
}

// save_ut_matrix - transpiled function from  $GOPATH/src/github.com/Konstantin8105/History_frame3DD/src/NRutil.c:1024
/*
 * SAVE_UT_MATRIX  - 						     23apr01
 * save a symmetric matrix of dimension [1..n][1..n] to the named file
 *  use only upper-triangular part
 */ //
//
func save_ut_matrix(filename []byte, A [][]float32, n int, mode []byte) {
	var fp_m *noarch.File
	var i int
	var j int
	var now noarch.TimeT
	if (func() *noarch.File {
		fp_m = noarch.Fopen(filename, mode)
		return fp_m
	}()) == nil {
		noarch.Printf([]byte(" error: cannot open file: %s \n\x00"), filename)
		os.Exit(1015)
	}
	_ = noarch.Time((*[100000000]noarch.TimeT)(unsafe.Pointer(&now))[:])
	noarch.Fprintf(fp_m, []byte("%% filename: %s - %s\x00"), filename, noarch.Ctime((*[100000000]noarch.TimeT)(unsafe.Pointer(&now))[:]))
	noarch.Fprintf(fp_m, []byte("%% type: matrix \n\x00"))
	noarch.Fprintf(fp_m, []byte("%% rows: %d\n\x00"), n)
	noarch.Fprintf(fp_m, []byte("%% columns: %d\n\x00"), n)
	for i = 1; i <= n; i++ {
		for j = 1; j <= n; j++ {
			if i > j {
				if math.Abs(float64(A[j][i])) > 1e-99 {
					noarch.Fprintf(fp_m, []byte("%15.6e\x00"), float64(A[j][i]))
				} else {
					noarch.Fprintf(fp_m, []byte("    0          \x00"))
				}
			} else {
				if math.Abs(float64(A[i][j])) > 1e-99 {
					noarch.Fprintf(fp_m, []byte("%15.6e\x00"), float64(A[i][j]))
				} else {
					noarch.Fprintf(fp_m, []byte("    0          \x00"))
				}
			}
		}
		noarch.Fprintf(fp_m, []byte("\n\x00"))
	}
	noarch.Fclose(fp_m)
}

// save_ut_dmatrix - transpiled function from  $GOPATH/src/github.com/Konstantin8105/History_frame3DD/src/NRutil.c:1061
/*
 * SAVE_UT_DMATRIX  - 						23apr01
 * save a symetric matrix of dimension [1..n][1..n] to the named file
 * use only upper-triangular part
 */ //
//
func save_ut_dmatrix(filename []byte, A [][]float64, n int, mode []byte) {
	var fp_m *noarch.File
	var i int
	var j int
	var now noarch.TimeT
	if (func() *noarch.File {
		fp_m = noarch.Fopen(filename, mode)
		return fp_m
	}()) == nil {
		noarch.Printf([]byte(" error: cannot open file: %s \n\x00"), filename)
		os.Exit(1016)
	}
	_ = noarch.Time((*[100000000]noarch.TimeT)(unsafe.Pointer(&now))[:])
	noarch.Fprintf(fp_m, []byte("%% filename: %s - %s\n\x00"), filename, noarch.Ctime((*[100000000]noarch.TimeT)(unsafe.Pointer(&now))[:]))
	noarch.Fprintf(fp_m, []byte("%% type: matrix \n\x00"))
	noarch.Fprintf(fp_m, []byte("%% rows: %d\n\x00"), n)
	noarch.Fprintf(fp_m, []byte("%% columns: %d\n\x00"), n)
	for i = 1; i <= n; i++ {
		for j = 1; j <= n; j++ {
			if i > j {
				if math.Abs(A[j][i]) > 1e-99 {
					noarch.Fprintf(fp_m, []byte("%21.12e\x00"), A[j][i])
				} else {
					noarch.Fprintf(fp_m, []byte("    0                \x00"))
				}
			} else {
				if math.Abs(A[i][j]) > 1e-99 {
					noarch.Fprintf(fp_m, []byte("%21.12e\x00"), A[i][j])
				} else {
					noarch.Fprintf(fp_m, []byte("    0                \x00"))
				}
			}
		}
		noarch.Fprintf(fp_m, []byte("\n\x00"))
	}
	noarch.Fclose(fp_m)
}
func init() {
	stdin = noarch.Stdin
	stdout = noarch.Stdout
	stderr = noarch.Stderr
}

/* itoa moved to frame3dd_io.c */
/* removed strcat -- it's in <string.h> in the standard C library */
/* removed strcpy -- it's in <string.h> in the standard C library */
/* dots moved to frame3dd_io.c */
