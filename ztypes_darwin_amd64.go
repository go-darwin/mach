// Code generated by cmd/cgo -godefs; DO NOT EDIT.
// cgo -godefs types_darwin_amd64.go

package mach

type (
	ClockRes      = int32
	ErrorType     = int32
	HostNamePort  = uint32
	IPCSpace      = uint32
	MachMsgHeader = struct {
		Bits         uint32
		Size         uint32
		Remote_port  uint32
		Local_port   uint32
		Voucher_port uint32
		Id           int32
	}
	MachMsgOption     = int32
	MachMsgPriority   = uint32
	MachMsgReturn     = int32
	MachMsgSize       = uint32
	MachMsgTimeout    = uint32
	MachMsgTypeName   = uint32
	MachMsgTypeNumber = uint32
	MachPort          = uint32
	MachPortDelta     = int32
	MachPortFlavor    = int32
	MachPortName      = uint32
	MachPortNameArray = *uint32
	MachPortRight     = uint32
	MachTimespec      = struct {
		Sec  uint32
		Nsec int32
	}
	MachVMAddress                 = uint64
	MachVMOffset                  = uint64
	MachVMSize                    = uint64
	MachVoucherAttrKey            = uint32
	MachVoucherAttrRawRecipe      = *uint8
	MachVoucherAttrRawRecipeArray = *uint8
	MemEntryNamePort              = uint32
	MemoryObjectOffset            = uint64
	MemoryObjectSize              = uint64
	Natural                       = uint32
	SleepType                     = int32
	ThreadPort                    = uint32
	VMMap                         = uint32
	VMProt                        = int32
	VMPurgable                    = int32
)

type (
	KernInvalidArgs = struct {
		Dummy int32
	}
	MachReplyPortArgs = struct {
		Dummy int32
	}
	ThreadGetSpecialReplyPortArgs = struct {
		Dummy int32
	}
	ThreadSelfTrapArgs = struct {
		Dummy int32
	}
	TaskSelfTrapArgs = struct {
		Dummy int32
	}
	HostSelfTrapArgs = struct {
		Dummy int32
	}
	MachMsgOverwriteTrapArgs = struct {
		Msg_l_       [0]int8
		Msg          uint64
		Msg_r_       [0]int8
		Option_l_    [0]int8
		Option       int32
		Option_r_    [0]int8
		Send_size_l_ [0]int8
		Send_size    uint32
		Send_size_r_ [0]int8
		Rcv_size_l_  [0]int8
		Rcv_size     uint32
		Rcv_size_r_  [0]int8
		Rcv_name_l_  [0]int8
		Rcv_name     uint32
		Rcv_name_r_  [0]int8
		Timeout_l_   [0]int8
		Timeout      uint32
		Timeout_r_   [0]int8
		Priority_l_  [0]int8
		Priority     uint32
		Priority_r_  [0]int8
		Arg8_pad_    [4]int8
		Rcv_msg_l_   [0]int8
		Rcv_msg      uint64
	}
	SemaphoreSignalTrapArgs = struct {
		Name_l_ [0]int8
		Name    uint32
	}
	SemaphoreSignalAllTrapArgs = struct {
		Name_l_ [0]int8
		Name    uint32
	}
	SemaphoreSignalThreadTrapArgs = struct {
		Signal_name_l_ [0]int8
		Signal_name    uint32
		Signal_name_r_ [0]int8
		Thread_name_l_ [0]int8
		Thread_name    uint32
	}
	SemaphoreWaitTrapArgs = struct {
		Name_l_ [0]int8
		Name    uint32
	}
	SemaphoreWaitSignalTrapArgs = struct {
		Wait_name_l_   [0]int8
		Wait_name      uint32
		Wait_name_r_   [0]int8
		Signal_name_l_ [0]int8
		Signal_name    uint32
	}
	SemaphoreTimedwaitTrapArgs = struct {
		Wait_name_l_ [0]int8
		Wait_name    uint32
		Wait_name_r_ [0]int8
		Sec_l_       [0]int8
		Sec          uint32
		Sec_r_       [0]int8
		Nsec_l_      [0]int8
		Nsec         int32
	}
	SemaphoreTimedwaitSignalTrapArgs = struct {
		Wait_name_l_   [0]int8
		Wait_name      uint32
		Wait_name_r_   [0]int8
		Signal_name_l_ [0]int8
		Signal_name    uint32
		Signal_name_r_ [0]int8
		Sec_l_         [0]int8
		Sec            uint32
		Sec_r_         [0]int8
		Nsec_l_        [0]int8
		Nsec           int32
	}
	TaskForPidArgs = struct {
		Target_tport_l_ [0]int8
		Target_tport    uint32
		Target_tport_r_ [0]int8
		Pid_l_          [0]int8
		Pid             int32
		Pid_r_          [0]int8
		T_l_            [0]int8
		T               uint64
	}
	TaskNameForPidArgs = struct {
		Target_tport_l_ [0]int8
		Target_tport    uint32
		Target_tport_r_ [0]int8
		Pid_l_          [0]int8
		Pid             int32
		Pid_r_          [0]int8
		T_l_            [0]int8
		T               uint64
	}
	PidForTaskArgs = struct {
		T_l_   [0]int8
		T      uint32
		T_r_   [0]int8
		Pid_l_ [0]int8
		Pid    uint64
	}
	DebugControlPortForPidArgs = struct {
		Target_tport_l_ [0]int8
		Target_tport    uint32
		Target_tport_r_ [0]int8
		Pid_l_          [0]int8
		Pid             int32
		Pid_r_          [0]int8
		T_l_            [0]int8
		T               uint64
	}
	MacxSwaponArgs = struct {
		Filename_l_ [0]int8
		Filename    uint64
		Filename_r_ [0]int8
		Flags_l_    [0]int8
		Flags       int32
		Flags_r_    [0]int8
		Size_l_     [0]int8
		Size        int32
		Size_r_     [0]int8
		Priority_l_ [0]int8
		Priority    int32
		Priority_r_ [0]int8
		Pad_cgo_0   [4]byte
	}
	MacxSwapoffArgs = struct {
		Filename_l_ [0]int8
		Filename    uint64
		Filename_r_ [0]int8
		Flags_l_    [0]int8
		Flags       int32
		Flags_r_    [0]int8
		Pad_cgo_0   [4]byte
	}
	MacxTriggersArgs = struct {
		Hi_water_l_   [0]int8
		Hi_water      int32
		Hi_water_r_   [0]int8
		Low_water_l_  [0]int8
		Low_water     int32
		Low_water_r_  [0]int8
		Flags_l_      [0]int8
		Flags         int32
		Flags_r_      [0]int8
		Alert_port_l_ [0]int8
		Alert_port    uint32
	}
	MacxBackingStoreSuspendArgs = struct {
		L_      [0]int8
		Suspend uint32
	}
	MacxBackingStoreRecoveryArgs = struct {
		L_  [0]int8
		Pid int32
	}
	SwtchPriArgs = struct {
		L_  [0]int8
		Pri int32
	}
	PfzExitArgs = struct {
		Dummy int32
	}
	SwtchArgs = struct {
		Dummy int32
	}
	ClockSleepTrapArgs = struct {
		Clock_name_l_  [0]int8
		Clock_name     uint32
		Clock_name_r_  [0]int8
		Sleep_type_l_  [0]int8
		Sleep_type     int32
		Sleep_type_r_  [0]int8
		Sleep_sec_l_   [0]int8
		Sleep_sec      int32
		Sleep_sec_r_   [0]int8
		Sleep_nsec_l_  [0]int8
		Sleep_nsec     int32
		Sleep_nsec_r_  [0]int8
		Wakeup_time_l_ [0]int8
		Wakeup_time    uint64
	}
	ThreadSwitchArgs = struct {
		Thread_name_l_ [0]int8
		Thread_name    uint32
		Thread_name_r_ [0]int8
		Option_l_      [0]int8
		Option         int32
		Option_r_      [0]int8
		Option_time_l_ [0]int8
		Option_time    uint32
	}
	MachTimebaseInfoTrapArgs = struct {
		L_   [0]int8
		Info uint64
	}
	MachWaitUntilTrapArgs = struct {
		L_       [0]int8
		Deadline uint64
	}
	MkTimerCreateTrapArgs = struct {
		Dummy int32
	}
	MkTimerDestroyTrapArgs = struct {
		L_   [0]int8
		Name uint32
	}
	MkTimerArmTrapArgs = struct {
		Name_l_        [0]int8
		Name           uint32
		Name_r_        [0]int8
		Expire_time_l_ [0]int8
		Expire_time    uint64
	}
	MkTimerArmLeewayTrapArgs = struct {
		Name_l_           [0]int8
		Name              uint32
		Name_r_           [0]int8
		Mk_timer_flags_l_ [0]int8
		Mk_timer_flags    uint64
		Mk_timer_flags_r_ [0]int8
		Expire_time_l_    [0]int8
		Expire_time       uint64
		Expire_time_r_    [0]int8
		Mk_leeway_l_      [0]int8
		Mk_leeway         uint64
	}
	MkTimerCancelTrapArgs = struct {
		Name_l_        [0]int8
		Name           uint32
		Name_r_        [0]int8
		Result_time_l_ [0]int8
		Result_time    uint64
	}
	KernelrpcMachVmAllocateTrapArgs = struct {
		Target_l_ [0]int8
		Target    uint32
		Target_r_ [0]int8
		Addr_l_   [0]int8
		Addr      uint64
		Addr_r_   [0]int8
		Size_l_   [0]int8
		Size      uint64
		Size_r_   [0]int8
		Flags_l_  [0]int8
		Flags     int32
		Flags_r_  [0]int8
		Pad_cgo_0 [4]byte
	}
	KernelrpcMachVmDeallocateArgs = struct {
		Target_l_  [0]int8
		Target     uint32
		Target_r_  [0]int8
		Address_l_ [0]int8
		Address    uint64
		Address_r_ [0]int8
		Size_l_    [0]int8
		Size       uint64
	}
	TaskDyldProcessInfoNotifyGetTrapArgs = struct {
		Addr_l_       [0]int8
		Addr          uint64
		Addr_r_       [0]int8
		Count_addr_l_ [0]int8
		Count_addr    uint64
	}
	KernelrpcMachVmProtectArgs = struct {
		Target_l_         [0]int8
		Target            uint32
		Target_r_         [0]int8
		Address_l_        [0]int8
		Address           uint64
		Address_r_        [0]int8
		Size_l_           [0]int8
		Size              uint64
		Size_r_           [0]int8
		Set_maximum_l_    [0]int8
		Set_maximum       uint32
		Set_maximum_r_    [0]int8
		New_protection_l_ [0]int8
		New_protection    int32
	}
	KernelrpcMachVmMapTrapArgs = struct {
		Target_l_         [0]int8
		Target            uint32
		Target_r_         [0]int8
		Addr_l_           [0]int8
		Addr              uint64
		Addr_r_           [0]int8
		Size_l_           [0]int8
		Size              uint64
		Size_r_           [0]int8
		Mask_l_           [0]int8
		Mask              uint64
		Mask_r_           [0]int8
		Flags_l_          [0]int8
		Flags             int32
		Flags_r_          [0]int8
		Arg8_pad_         [4]int8
		Cur_protection_l_ [0]int8
		Cur_protection    int32
		Cur_protection_r_ [0]int8
		Pad_cgo_0         [4]byte
	}
	KernelrpcMachVmPurgableControlTrapArgs = struct {
		Target_l_  [0]int8
		Target     uint32
		Target_r_  [0]int8
		Address_l_ [0]int8
		Address    uint64
		Address_r_ [0]int8
		Control_l_ [0]int8
		Control    int32
		Control_r_ [0]int8
		State_l_   [0]int8
		State      uint64
	}
	KernelrpcMachPortAllocateArgs = struct {
		Target_l_ [0]int8
		Target    uint32
		Target_r_ [0]int8
		Right_l_  [0]int8
		Right     uint32
		Right_r_  [0]int8
		Name_l_   [0]int8
		Name      uint64
	}
	KernelrpcMachPortDeallocateArgs = struct {
		Target_l_ [0]int8
		Target    uint32
		Target_r_ [0]int8
		Name_l_   [0]int8
		Name      uint32
	}
	KernelrpcMachPortModRefsArgs = struct {
		Target_l_ [0]int8
		Target    uint32
		Target_r_ [0]int8
		Name_l_   [0]int8
		Name      uint32
		Name_r_   [0]int8
		Right_l_  [0]int8
		Right     uint32
		Right_r_  [0]int8
		Delta_l_  [0]int8
		Delta     int32
	}
	KernelrpcMachPortMoveMemberArgs = struct {
		Target_l_ [0]int8
		Target    uint32
		Target_r_ [0]int8
		Member_l_ [0]int8
		Member    uint32
		Member_r_ [0]int8
		After_l_  [0]int8
		After     uint32
	}
	KernelrpcMachPortInsertRightArgs = struct {
		Target_l_   [0]int8
		Target      uint32
		Target_r_   [0]int8
		Name_l_     [0]int8
		Name        uint32
		Name_r_     [0]int8
		Poly_l_     [0]int8
		Poly        uint32
		Poly_r_     [0]int8
		PolyPoly_l_ [0]int8
		PolyPoly    uint32
	}
	KernelrpcMachPortGetAttributesArgs = struct {
		Target_l_ [0]int8
		Target    uint32
		Target_r_ [0]int8
		Name_l_   [0]int8
		Name      uint32
		Name_r_   [0]int8
		Flavor_l_ [0]int8
		Flavor    int32
		Flavor_r_ [0]int8
		Info_l_   [0]int8
		Info      uint64
		Info_r_   [0]int8
		Count_l_  [0]int8
		Count     uint64
	}
	KernelrpcMachPortInsertMemberArgs = struct {
		Target_l_ [0]int8
		Target    uint32
		Target_r_ [0]int8
		Name_l_   [0]int8
		Name      uint32
		Name_r_   [0]int8
		Pset_l_   [0]int8
		Pset      uint32
	}
	KernelrpcMachPortExtractMemberArgs = struct {
		Target_l_ [0]int8
		Target    uint32
		Target_r_ [0]int8
		Name_l_   [0]int8
		Name      uint32
		Name_r_   [0]int8
		Pset_l_   [0]int8
		Pset      uint32
	}
	KernelrpcMachPortConstructArgs = struct {
		Target_l_  [0]int8
		Target     uint32
		Target_r_  [0]int8
		Options_l_ [0]int8
		Options    uint64
		Options_r_ [0]int8
		Context_l_ [0]int8
		Context    uint64
		Context_r_ [0]int8
		Name_l_    [0]int8
		Name       uint64
	}
	KernelrpcMachPortDestructArgs = struct {
		Target_l_  [0]int8
		Target     uint32
		Target_r_  [0]int8
		Name_l_    [0]int8
		Name       uint32
		Name_r_    [0]int8
		Srdelta_l_ [0]int8
		Srdelta    int32
		Srdelta_r_ [0]int8
		Guard_l_   [0]int8
		Guard      uint64
	}
	KernelrpcMachPortGuardArgs = struct {
		Target_l_ [0]int8
		Target    uint32
		Target_r_ [0]int8
		Name_l_   [0]int8
		Name      uint32
		Name_r_   [0]int8
		Guard_l_  [0]int8
		Guard     uint64
		Guard_r_  [0]int8
		Strict_l_ [0]int8
		Strict    uint32
		Strict_r_ [0]int8
		Pad_cgo_0 [4]byte
	}
	KernelrpcMachPortUnguardArgs = struct {
		Target_l_ [0]int8
		Target    uint32
		Target_r_ [0]int8
		Name_l_   [0]int8
		Name      uint32
		Name_r_   [0]int8
		Guard_l_  [0]int8
		Guard     uint64
	}
	MachGenerateActivityIdArgs = struct {
		Target_l_      [0]int8
		Target         uint32
		Target_r_      [0]int8
		Count_l_       [0]int8
		Count          int32
		Count_r_       [0]int8
		Activity_id_l_ [0]int8
		Activity_id    uint64
	}

	HostCreateMachVoucherArgs = struct {
		Host_l_         [0]int8
		Host            uint32
		Host_r_         [0]int8
		Recipes_l_      [0]int8
		Recipes         *uint8
		Recipes_r_      [0]int8
		Recipes_size_l_ [0]int8
		Recipes_size    int32
		Recipes_size_r_ [0]int8
		Voucher_l_      [0]int8
		Voucher         uint64
	}
	IokitUserClientTrapArgs = struct {
		UserClientRef_l_ [0]int8
		UserClientRef    *byte
		UserClientRef_r_ [0]int8
		Index_l_         [0]int8
		Index            uint32
		Index_r_         [0]int8
		P1_l_            [0]int8
		P1               *byte
		P1_r_            [0]int8
		P2_l_            [0]int8
		P2               *byte
		P2_r_            [0]int8
		P3_l_            [0]int8
		P3               *byte
		P3_r_            [0]int8
		P4_l_            [0]int8
		P4               *byte
		P4_r_            [0]int8
		P5_l_            [0]int8
		P5               *byte
		P5_r_            [0]int8
		Arg8_pad_        [4]int8
		P6_l_            [0]int8
		P6               *byte
	}
	KernelrpcMachPortRequestNotificationArgs = struct {
		Target_l_     [0]int8
		Target        uint32
		Target_r_     [0]int8
		Name_l_       [0]int8
		Name          uint32
		Name_r_       [0]int8
		Msgid_l_      [0]int8
		Msgid         int32
		Msgid_r_      [0]int8
		Sync_l_       [0]int8
		Sync          uint32
		Sync_r_       [0]int8
		Notify_l_     [0]int8
		Notify        uint32
		Notify_r_     [0]int8
		NotifyPoly_l_ [0]int8
		NotifyPoly    uint32
		NotifyPoly_r_ [0]int8
		Previous_l_   [0]int8
		Previous      uint64
	}
	KernelrpcMachPortTypeArgs = struct {
		Target_l_ [0]int8
		Target    uint32
		Target_r_ [0]int8
		Name_l_   [0]int8
		Name      uint32
		Name_r_   [0]int8
		Ptype_l_  [0]int8
		Ptype     uint64
	}
	MachVoucherExtractAttrRecipeArgs = struct {
		Voucher_name_l_ [0]int8
		Voucher_name    uint32
		Voucher_name_r_ [0]int8
		Key_l_          [0]int8
		Key             uint32
		Key_r_          [0]int8
		Recipe_l_       [0]int8
		Recipe          *uint8
		Recipe_r_       [0]int8
		Recipe_size_l_  [0]int8
		Recipe_size     uint64
	}
)
