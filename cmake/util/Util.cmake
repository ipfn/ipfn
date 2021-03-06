# Original Source: https://github.com/dmlc/tvm/blob/master/cmake/util/Util.cmake
#
# Copyright (c) 2018 TVM Contributors
# Copyright (c) 2018 IPFN Developers
#
# Licensed under Apache-2.0 or MIT License.
#

macro(__ipfn_option variable description value)
  if(NOT DEFINED ${variable})
    set(${variable} ${value} CACHE STRING ${description})
  endif()
endmacro()

#######################################################
# An option that the user can select. Can accept condition to control when option is available for user.
# Usage:
#   ipfn_option(<option_variable> "doc string" <initial value or boolean expression> [IF <condition>])
macro(ipfn_option variable description value)
  set(__value ${value})
  set(__condition "")
  set(__varname "__value")
  foreach(arg ${ARGN})
    if(arg STREQUAL "IF" OR arg STREQUAL "if")
      set(__varname "__condition")
    else()
      list(APPEND ${__varname} ${arg})
    endif()
  endforeach()
  unset(__varname)
  if("${__condition}" STREQUAL "")
    set(__condition 2 GREATER 1)
  endif()

  if(${__condition})
    if("${__value}" MATCHES ";")
      if(${__value})
        __ipfn_option(${variable} "${description}" ON)
      else()
        __ipfn_option(${variable} "${description}" OFF)
      endif()
    elseif(DEFINED ${__value})
      if(${__value})
        __ipfn_option(${variable} "${description}" ON)
      else()
        __ipfn_option(${variable} "${description}" OFF)
      endif()
    else()
      __ipfn_option(${variable} "${description}" "${__value}")
    endif()
  else()
    unset(${variable} CACHE)
  endif()
endmacro()

macro(set_option_force option value)
  set(${option} ${value} CACHE "" INTERNAL FORCE)
endmacro()

function(assign_source_group group)
    foreach(_source IN ITEMS ${ARGN})
        if (IS_ABSOLUTE "${_source}")
            file(RELATIVE_PATH _source_rel "${CMAKE_CURRENT_SOURCE_DIR}" "${_source}")
        else()
            set(_source_rel "${_source}")
        endif()
        get_filename_component(_source_path "${_source_rel}" PATH)
        string(REPLACE "/" "\\" _source_path_msvc "${_source_path}")
        source_group("${group}\\${_source_path_msvc}" FILES "${_source}")
    endforeach()
endfunction(assign_source_group)
