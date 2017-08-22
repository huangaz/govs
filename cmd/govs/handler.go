/*
 * Copyright 2016 Xiaomi Corporation. All rights reserved.
 * Use of this source code is governed by a BSD-style
 * license that can be found in the LICENSE file.
 *
 * Authors:    Yu Bo <yubo@xiaomi.com>
 */
package main

import (
	"fmt"
	"os"

	"github.com/dpvs/govs"
	"github.com/yubo/gotool/flags"
)

func init() {
	//flags.CommandLine.Usage = fmt.Sprintf("Usage: %s COMMAND [OPTIONS] host[:port]\n\n",
	//	os.Args[0])

	flags.FirstCmd.BoolVar(&govs.FirstCmd.ADD, "A", false, "add")
	flags.FirstCmd.BoolVar(&govs.FirstCmd.EDIT, "E", false, "edit Service")
	flags.FirstCmd.BoolVar(&govs.FirstCmd.DEL, "D", false, "del Service")
	flags.FirstCmd.BoolVar(&govs.FirstCmd.ADDDEST, "a", false, "add Real Server")
	flags.FirstCmd.BoolVar(&govs.FirstCmd.EDITDEST, "e", false, "edit Real Server")
	flags.FirstCmd.BoolVar(&govs.FirstCmd.DELDEST, "d", false, "del Real Server")
	flags.FirstCmd.BoolVar(&govs.FirstCmd.FLUSH, "C", false, "flush")
	flags.FirstCmd.BoolVar(&govs.FirstCmd.LIST, "L", false, "list Service")
	flags.FirstCmd.BoolVar(&govs.FirstCmd.LIST, "l", false, "list Service")
	flags.FirstCmd.BoolVar(&govs.FirstCmd.ZERO, "Z", false, "zero")
	flags.FirstCmd.BoolVar(&govs.FirstCmd.TIMEOUT, "TAG_SET", false, "timeout")
	flags.FirstCmd.BoolVar(&govs.FirstCmd.USAGE, "h", false, "show usage")
	flags.FirstCmd.BoolVar(&govs.FirstCmd.VERSION, "V", false, "show version")
	flags.FirstCmd.BoolVar(&govs.FirstCmd.ADDLADDR, "P", false, "add Local Address")
	flags.FirstCmd.BoolVar(&govs.FirstCmd.DELLADDR, "Q", false, "del Local Address")
	flags.FirstCmd.BoolVar(&govs.FirstCmd.GETLADDR, "G", false, "get Local Address")
	flags.FirstCmd.BoolVar(&govs.FirstCmd.STATUS, "S", false, "show status")

	flags.OthersCmd.StringVar(&govs.CmdOpt.TCP, "t", "", "tcp service")
	flags.OthersCmd.StringVar(&govs.CmdOpt.UDP, "u", "", "udp service")
	flags.OthersCmd.Var(&govs.CmdOpt.Netmask, "M", "netmask deafult 0.0.0.0")
	flags.OthersCmd.StringVar(&govs.CmdOpt.Sched_name, "s", "rr", "scheduler name rr/wrr")
	flags.OthersCmd.UintVar(&govs.CmdOpt.Flags, "flags", 0, "the service flags")
	flags.OthersCmd.Var(&govs.CmdOpt.Daddr, "r", "service-address is host[:port]")
	flags.OthersCmd.IntVar(&govs.CmdOpt.Weight, "w", 0, "capacity of real server")
	flags.OthersCmd.UintVar(&govs.CmdOpt.U_threshold, "x", 0, "upper threshold of connections")
	flags.OthersCmd.UintVar(&govs.CmdOpt.L_threshold, "y", 0, "lower threshold of connections")
	flags.OthersCmd.Var(&govs.CmdOpt.Lip, "z", "local-address is host")
	flags.OthersCmd.StringVar(&govs.CmdOpt.Typ, "type", "io", "type of the stats name(io/w/we/dev/ctl/mem)")
	flags.OthersCmd.IntVar(&govs.CmdOpt.Id, "i", -1, "id of the stats object")
	flags.OthersCmd.StringVar(&govs.CmdOpt.Timeout_s, "set", "", "set <tcp,tcp_fin,udp>")
	flags.OthersCmd.UintVar(&govs.CmdOpt.Conn_flags, "conn_flags", 0, "the conn flags")
}

func handler() {
	flags.FirstCmd.Parse(os.Args[1:2])
	switch {
	case govs.FirstCmd.ADD:
		fmt.Println("add Service")
		flags.Cmd.Action = add_handle
	case govs.FirstCmd.EDIT:
		fmt.Println("edit Service")
		flags.Cmd.Action = edit_handle
	case govs.FirstCmd.DEL:
		fmt.Println("del Service")
		flags.Cmd.Action = del_handle
	case govs.FirstCmd.ADDDEST:
		fmt.Println("add Real Server")
		flags.Cmd.Action = addsrv_handle
	case govs.FirstCmd.EDITDEST:
		fmt.Println("edit Real Server")
		flags.Cmd.Action = editsrv_handle
	case govs.FirstCmd.DELDEST:
		fmt.Println("del Real Server")
		flags.Cmd.Action = delsrv_handle
	case govs.FirstCmd.ADDLADDR:
		fmt.Println("add Local Address")
		flags.Cmd.Action = addladdr_handle
	case govs.FirstCmd.DELLADDR:
		fmt.Println("del Local Address")
		flags.Cmd.Action = delladdr_handle
	case govs.FirstCmd.GETLADDR:
		fmt.Println("get Local Address")
		flags.Cmd.Action = list_handle
	case govs.FirstCmd.FLUSH:
		fmt.Println("flush")
		flags.Cmd.Action = flush_handle
	case govs.FirstCmd.LIST:
		fmt.Println("list")
		flags.Cmd.Action = list_handle
	case govs.FirstCmd.STATUS:
		fmt.Println("status")
		flags.Cmd.Action = stats_handle
	case govs.FirstCmd.TIMEOUT:
		fmt.Println("timeout")
		flags.Cmd.Action = timeout_handle
	case govs.FirstCmd.USAGE:
		fmt.Println("show usage")
	case govs.FirstCmd.VERSION:
		fmt.Println("gei version")
		flags.Cmd.Action = version_handle
	case govs.FirstCmd.ZERO:
		fmt.Println("zero")
		flags.Cmd.Action = zero_handle
	default:
		fmt.Println("error!!!")
	}
	flags.OthersCmd.Parse(os.Args[2:])
	//generic_opt_check
}

func version_handle(arg interface{}) {
	if version, err := govs.Get_version(); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(version)
	}
}

func info_handle(arg interface{}) {
	if info, err := govs.Get_version(); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(info)
	}
}

func timeout_handle(arg interface{}) {
	opt := arg.(*govs.CallOptions)
	o := &opt.Opt

	if o.Timeout_s != "" {
		if timeout, err := govs.Set_timeout(o); err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(timeout)
		}
	} else {
		if timeout, err := govs.Get_timeout(o); err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(timeout)
		}
	}
}

func list_svc_handle(o *govs.CmdOptions) {

	ret, err := govs.Get_service(o)
	if err != nil {
		fmt.Println(err)
		return
	}

	if ret.Code != 0 {
		fmt.Println(ret.Msg)
		return
	}

	fmt.Println(govs.Svc_title())
	if !govs.FirstCmd.GETLADDR {
		fmt.Println(govs.Dest_title())
	} else {
		fmt.Println(govs.Laddr_title())
	}

	fmt.Println(ret)
	if !govs.FirstCmd.GETLADDR {
		dests, err := govs.Get_dests(o)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(dests)
	} else {
		laddrs, err := govs.Get_laddrs(o)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(laddrs)
	}
	return
}

func list_svcs_handle(o *govs.CmdOptions) {

	ret, err := govs.Get_services(o)

	if err != nil {
		fmt.Println(err)
		return
	}

	if ret.Code != 0 {
		fmt.Println(ret.Msg)
		return
	}

	fmt.Println(govs.Svc_title())
	if !govs.FirstCmd.GETLADDR {
		fmt.Println(govs.Dest_title())
	} else {
		fmt.Println(govs.Laddr_title())
	}

	for _, svc := range ret.Services {
		fmt.Println(svc)
		o.Addr.Ip = svc.Addr
		o.Addr.Port = svc.Port
		o.Protocol = govs.Protocol(svc.Protocol)

		if !govs.FirstCmd.GETLADDR {
			dests, err := govs.Get_dests(o)
			if err != nil || dests.Code != 0 ||
				len(dests.Dests) == 0 {
				//fmt.Println(err)
				continue
			}
			fmt.Println(dests)
		} else {
			laddrs, err := govs.Get_laddrs(o)
			if err != nil || laddrs.Code != 0 ||
				len(laddrs.Laddrs) == 0 {
				//fmt.Println(err)
				return
			}
			fmt.Println(laddrs)
		}
	}

}

func list_handle(arg interface{}) {
	opt := arg.(*govs.CallOptions)
	govs.Parse_service(opt)
	o := &opt.Opt

	if o.Addr.Ip != 0 {
		list_svc_handle(o)
		return
	}

	list_svcs_handle(o)
}

func flush_handle(arg interface{}) {
	if reply, err := govs.Set_flush(nil); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(reply)
	}
}

func zero_handle(arg interface{}) {
	opt := arg.(*govs.CallOptions)
	govs.Parse_service(opt)

	if reply, err := govs.Set_zero(&opt.Opt); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(reply)
	}

}

func add_handle(arg interface{}) {
	var err error
	var reply *govs.Vs_cmd_r

	opt := arg.(*govs.CallOptions)
	if err := govs.Parse_service(opt); err != nil {
		fmt.Println(err)
		return
	}
	o := &opt.Opt

	reply, err = govs.Set_add(o)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(reply)
	}
}

func addsrv_handle(arg interface{}) {
	var err error
	var reply *govs.Vs_cmd_r

	opt := arg.(*govs.CallOptions)
	if err := govs.Parse_service(opt); err != nil {
		fmt.Println(err)
		return
	}
	o := &opt.Opt

	reply, err = govs.Set_adddest(o)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(reply)
	}
}

func addladdr_handle(arg interface{}) {
	var err error
	var reply *govs.Vs_cmd_r

	opt := arg.(*govs.CallOptions)
	if err := govs.Parse_service(opt); err != nil {
		fmt.Println(err)
		return
	}
	o := &opt.Opt

	reply, err = govs.Set_addladdr(o)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(reply)
	}
}

func edit_handle(arg interface{}) {
	var err error
	var reply *govs.Vs_cmd_r

	opt := arg.(*govs.CallOptions)
	if err := govs.Parse_service(opt); err != nil {
		fmt.Println(err)
		return
	}
	o := &opt.Opt

	reply, err = govs.Set_edit(o)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(reply)
	}
}

func editsrv_handle(arg interface{}) {
	var err error
	var reply *govs.Vs_cmd_r

	opt := arg.(*govs.CallOptions)
	if err := govs.Parse_service(opt); err != nil {
		fmt.Println(err)
		return
	}
	o := &opt.Opt

	reply, err = govs.Set_editdest(o)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(reply)
	}
}

func del_handle(arg interface{}) {
	var err error
	var reply *govs.Vs_cmd_r

	opt := arg.(*govs.CallOptions)
	if err := govs.Parse_service(opt); err != nil {
		fmt.Println(err)
		return
	}
	o := &opt.Opt

	reply, err = govs.Set_del(o)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(reply)
	}
}

func delsrv_handle(arg interface{}) {
	var err error
	var reply *govs.Vs_cmd_r

	opt := arg.(*govs.CallOptions)
	if err := govs.Parse_service(opt); err != nil {
		fmt.Println(err)
		return
	}
	o := &opt.Opt

	reply, err = govs.Set_deldest(o)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(reply)
	}
}

func delladdr_handle(arg interface{}) {
	var err error
	var reply *govs.Vs_cmd_r

	opt := arg.(*govs.CallOptions)
	if err := govs.Parse_service(opt); err != nil {
		fmt.Println(err)
		return
	}
	o := &opt.Opt

	reply, err = govs.Set_delladdr(o)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(reply)
	}
}

func stats_handle(arg interface{}) {
	id := govs.CmdOpt.Id

	switch govs.CmdOpt.Typ {
	case "io":
		relay, err := govs.Get_stats_io(id)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf(relay)
	case "w":
		relay, err := govs.Get_stats_worker(id)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(relay)
	case "we":
		relay, err := govs.Get_estats_worker(id)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(relay)
	case "dev":
		relay, err := govs.Get_stats_dev(id)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(relay)
	case "ctl":
		relay, err := govs.Get_stats_ctl()
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(relay)
	case "mem":
		relay, err := govs.Get_stats_mem()
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(relay)
	default:
		fmt.Println("govs stats -t io/worker/dev/ctl")
	}
}
