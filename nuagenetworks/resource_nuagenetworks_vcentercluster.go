package nuagenetworks

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/rvichery/vspk-go/vspk"
)

func resourceVCenterCluster() *schema.Resource {
	return &schema.Resource{
		Create: resourceVCenterClusterCreate,
		Read:   resourceVCenterClusterRead,
		Update: resourceVCenterClusterUpdate,
		Delete: resourceVCenterClusterDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"parent_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"parent_type": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"owner": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vrs_configuration_time_limit": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"v_require_nuage_metadata": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"managed_object_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"last_updated_by": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"data_dns1": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"data_dns2": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"data_gateway": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"data_network_portgroup": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"datapath_sync_timeout": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"scope": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"secondary_data_uplink_dhcp_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"secondary_data_uplink_enabled": {
				Type:     schema.TypeBool,
				Required: true,
			},
			"secondary_data_uplink_interface": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"secondary_data_uplink_mtu": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  1500,
			},
			"secondary_data_uplink_primary_controller": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"secondary_data_uplink_secondary_controller": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"secondary_data_uplink_underlay_id": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  1,
			},
			"secondary_nuage_controller": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"deleted_from_vcenter_data_center": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"memory_size_in_gb": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "DEFAULT_4",
			},
			"remote_syslog_server_ip": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"remote_syslog_server_port": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  514,
			},
			"remote_syslog_server_type": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "NONE",
			},
			"generic_split_activation": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"separate_data_network": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"personality": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"destination_mirror_port": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "no_mirror",
			},
			"metadata_server_ip": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"metadata_server_listen_port": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"metadata_server_port": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"metadata_service_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"network_uplink_interface": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"network_uplink_interface_gateway": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"network_uplink_interface_ip": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"network_uplink_interface_netmask": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"revertive_controller_enabled": {
				Type:     schema.TypeBool,
				Required: true,
			},
			"revertive_timer": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"nfs_log_server": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"nfs_mount_path": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"mgmt_dns1": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"mgmt_dns2": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"mgmt_gateway": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"mgmt_network_portgroup": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dhcp_relay_server": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"mirror_network_portgroup": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"disable_gro_on_datapath": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"disable_lro_on_datapath": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"site_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"allow_data_dhcp": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"allow_mgmt_dhcp": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"flow_eviction_threshold": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"vm_network_portgroup": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"enable_vrs_resource_reservation": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"entity_scope": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"configured_metrics_push_interval": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  60,
			},
			"portgroup_metadata": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"nova_client_version": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"nova_identity_url_version": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"nova_metadata_service_auth_url": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"nova_metadata_service_endpoint": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"nova_metadata_service_password": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"nova_metadata_service_tenant": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"nova_metadata_service_username": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"nova_metadata_shared_secret": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"nova_os_keystone_username": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"nova_project_domain_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"nova_project_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"nova_region_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"nova_user_domain_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"upgrade_package_password": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"upgrade_package_url": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"upgrade_package_username": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"upgrade_script_time_limit": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"cpu_count": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "DEFAULT_2",
			},
			"primary_data_uplink_underlay_id": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  0,
			},
			"primary_nuage_controller": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vrs_password": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vrs_user_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"assoc_vcenter_data_center_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"assoc_vcenter_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"static_route": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"static_route_gateway": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"static_route_netmask": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ntp_server1": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ntp_server2": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"mtu": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"multi_vmssupport": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"multicast_receive_interface": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"multicast_receive_interface_ip": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"multicast_receive_interface_netmask": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"multicast_receive_range": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"multicast_send_interface": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"multicast_send_interface_ip": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"multicast_send_interface_netmask": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"multicast_source_portgroup": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"customized_script_url": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ovf_url": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"avrs_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"avrs_profile": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "AVRS_25G",
			},
			"external_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"parent_vcenter_data_center": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceVCenterClusterCreate(d *schema.ResourceData, m interface{}) error {

	// Initialize VCenterCluster object
	SecondaryDataUplinkEnabled := d.Get("secondary_data_uplink_enabled").(bool)
	RevertiveControllerEnabled := d.Get("revertive_controller_enabled").(bool)
	RevertiveTimer := d.Get("revertive_timer").(int)
	o := &vspk.VCenterCluster{
		Name: d.Get("name").(string),
		SecondaryDataUplinkEnabled: &SecondaryDataUplinkEnabled,
		RevertiveControllerEnabled: &RevertiveControllerEnabled,
		RevertiveTimer:             &RevertiveTimer,
	}
	if attr, ok := d.GetOk("vrs_configuration_time_limit"); ok {
		VRSConfigurationTimeLimit := attr.(int)
		o.VRSConfigurationTimeLimit = &VRSConfigurationTimeLimit
	}
	if attr, ok := d.GetOk("v_require_nuage_metadata"); ok {
		VRequireNuageMetadata := attr.(bool)
		o.VRequireNuageMetadata = &VRequireNuageMetadata
	}
	if attr, ok := d.GetOk("managed_object_id"); ok {
		o.ManagedObjectID = attr.(string)
	}
	if attr, ok := d.GetOk("data_dns1"); ok {
		o.DataDNS1 = attr.(string)
	}
	if attr, ok := d.GetOk("data_dns2"); ok {
		o.DataDNS2 = attr.(string)
	}
	if attr, ok := d.GetOk("data_gateway"); ok {
		o.DataGateway = attr.(string)
	}
	if attr, ok := d.GetOk("data_network_portgroup"); ok {
		o.DataNetworkPortgroup = attr.(string)
	}
	if attr, ok := d.GetOk("datapath_sync_timeout"); ok {
		DatapathSyncTimeout := attr.(int)
		o.DatapathSyncTimeout = &DatapathSyncTimeout
	}
	if attr, ok := d.GetOk("scope"); ok {
		Scope := attr.(bool)
		o.Scope = &Scope
	}
	if attr, ok := d.GetOk("secondary_data_uplink_dhcp_enabled"); ok {
		SecondaryDataUplinkDHCPEnabled := attr.(bool)
		o.SecondaryDataUplinkDHCPEnabled = &SecondaryDataUplinkDHCPEnabled
	}
	if attr, ok := d.GetOk("secondary_data_uplink_interface"); ok {
		o.SecondaryDataUplinkInterface = attr.(string)
	}
	if attr, ok := d.GetOk("secondary_data_uplink_mtu"); ok {
		SecondaryDataUplinkMTU := attr.(int)
		o.SecondaryDataUplinkMTU = &SecondaryDataUplinkMTU
	}
	if attr, ok := d.GetOk("secondary_data_uplink_primary_controller"); ok {
		o.SecondaryDataUplinkPrimaryController = attr.(string)
	}
	if attr, ok := d.GetOk("secondary_data_uplink_secondary_controller"); ok {
		o.SecondaryDataUplinkSecondaryController = attr.(string)
	}
	if attr, ok := d.GetOk("secondary_data_uplink_underlay_id"); ok {
		SecondaryDataUplinkUnderlayID := attr.(int)
		o.SecondaryDataUplinkUnderlayID = &SecondaryDataUplinkUnderlayID
	}
	if attr, ok := d.GetOk("secondary_nuage_controller"); ok {
		o.SecondaryNuageController = attr.(string)
	}
	if attr, ok := d.GetOk("deleted_from_vcenter_data_center"); ok {
		DeletedFromVCenterDataCenter := attr.(bool)
		o.DeletedFromVCenterDataCenter = &DeletedFromVCenterDataCenter
	}
	if attr, ok := d.GetOk("memory_size_in_gb"); ok {
		o.MemorySizeInGB = attr.(string)
	}
	if attr, ok := d.GetOk("remote_syslog_server_ip"); ok {
		o.RemoteSyslogServerIP = attr.(string)
	}
	if attr, ok := d.GetOk("remote_syslog_server_port"); ok {
		RemoteSyslogServerPort := attr.(int)
		o.RemoteSyslogServerPort = &RemoteSyslogServerPort
	}
	if attr, ok := d.GetOk("remote_syslog_server_type"); ok {
		o.RemoteSyslogServerType = attr.(string)
	}
	if attr, ok := d.GetOk("generic_split_activation"); ok {
		GenericSplitActivation := attr.(bool)
		o.GenericSplitActivation = &GenericSplitActivation
	}
	if attr, ok := d.GetOk("separate_data_network"); ok {
		SeparateDataNetwork := attr.(bool)
		o.SeparateDataNetwork = &SeparateDataNetwork
	}
	if attr, ok := d.GetOk("personality"); ok {
		o.Personality = attr.(string)
	}
	if attr, ok := d.GetOk("description"); ok {
		o.Description = attr.(string)
	}
	if attr, ok := d.GetOk("destination_mirror_port"); ok {
		o.DestinationMirrorPort = attr.(string)
	}
	if attr, ok := d.GetOk("metadata_server_ip"); ok {
		o.MetadataServerIP = attr.(string)
	}
	if attr, ok := d.GetOk("metadata_server_listen_port"); ok {
		MetadataServerListenPort := attr.(int)
		o.MetadataServerListenPort = &MetadataServerListenPort
	}
	if attr, ok := d.GetOk("metadata_server_port"); ok {
		MetadataServerPort := attr.(int)
		o.MetadataServerPort = &MetadataServerPort
	}
	if attr, ok := d.GetOk("metadata_service_enabled"); ok {
		MetadataServiceEnabled := attr.(bool)
		o.MetadataServiceEnabled = &MetadataServiceEnabled
	}
	if attr, ok := d.GetOk("network_uplink_interface"); ok {
		o.NetworkUplinkInterface = attr.(string)
	}
	if attr, ok := d.GetOk("network_uplink_interface_gateway"); ok {
		o.NetworkUplinkInterfaceGateway = attr.(string)
	}
	if attr, ok := d.GetOk("network_uplink_interface_ip"); ok {
		o.NetworkUplinkInterfaceIp = attr.(string)
	}
	if attr, ok := d.GetOk("network_uplink_interface_netmask"); ok {
		o.NetworkUplinkInterfaceNetmask = attr.(string)
	}
	if attr, ok := d.GetOk("nfs_log_server"); ok {
		o.NfsLogServer = attr.(string)
	}
	if attr, ok := d.GetOk("nfs_mount_path"); ok {
		o.NfsMountPath = attr.(string)
	}
	if attr, ok := d.GetOk("mgmt_dns1"); ok {
		o.MgmtDNS1 = attr.(string)
	}
	if attr, ok := d.GetOk("mgmt_dns2"); ok {
		o.MgmtDNS2 = attr.(string)
	}
	if attr, ok := d.GetOk("mgmt_gateway"); ok {
		o.MgmtGateway = attr.(string)
	}
	if attr, ok := d.GetOk("mgmt_network_portgroup"); ok {
		o.MgmtNetworkPortgroup = attr.(string)
	}
	if attr, ok := d.GetOk("dhcp_relay_server"); ok {
		o.DhcpRelayServer = attr.(string)
	}
	if attr, ok := d.GetOk("mirror_network_portgroup"); ok {
		o.MirrorNetworkPortgroup = attr.(string)
	}
	if attr, ok := d.GetOk("disable_gro_on_datapath"); ok {
		DisableGROOnDatapath := attr.(bool)
		o.DisableGROOnDatapath = &DisableGROOnDatapath
	}
	if attr, ok := d.GetOk("disable_lro_on_datapath"); ok {
		DisableLROOnDatapath := attr.(bool)
		o.DisableLROOnDatapath = &DisableLROOnDatapath
	}
	if attr, ok := d.GetOk("site_id"); ok {
		o.SiteId = attr.(string)
	}
	if attr, ok := d.GetOk("allow_data_dhcp"); ok {
		AllowDataDHCP := attr.(bool)
		o.AllowDataDHCP = &AllowDataDHCP
	}
	if attr, ok := d.GetOk("allow_mgmt_dhcp"); ok {
		AllowMgmtDHCP := attr.(bool)
		o.AllowMgmtDHCP = &AllowMgmtDHCP
	}
	if attr, ok := d.GetOk("flow_eviction_threshold"); ok {
		FlowEvictionThreshold := attr.(int)
		o.FlowEvictionThreshold = &FlowEvictionThreshold
	}
	if attr, ok := d.GetOk("vm_network_portgroup"); ok {
		o.VmNetworkPortgroup = attr.(string)
	}
	if attr, ok := d.GetOk("enable_vrs_resource_reservation"); ok {
		EnableVRSResourceReservation := attr.(bool)
		o.EnableVRSResourceReservation = &EnableVRSResourceReservation
	}
	if attr, ok := d.GetOk("configured_metrics_push_interval"); ok {
		ConfiguredMetricsPushInterval := attr.(int)
		o.ConfiguredMetricsPushInterval = &ConfiguredMetricsPushInterval
	}
	if attr, ok := d.GetOk("portgroup_metadata"); ok {
		PortgroupMetadata := attr.(bool)
		o.PortgroupMetadata = &PortgroupMetadata
	}
	if attr, ok := d.GetOk("nova_client_version"); ok {
		NovaClientVersion := attr.(int)
		o.NovaClientVersion = &NovaClientVersion
	}
	if attr, ok := d.GetOk("nova_identity_url_version"); ok {
		o.NovaIdentityURLVersion = attr.(string)
	}
	if attr, ok := d.GetOk("nova_metadata_service_auth_url"); ok {
		o.NovaMetadataServiceAuthUrl = attr.(string)
	}
	if attr, ok := d.GetOk("nova_metadata_service_endpoint"); ok {
		o.NovaMetadataServiceEndpoint = attr.(string)
	}
	if attr, ok := d.GetOk("nova_metadata_service_password"); ok {
		o.NovaMetadataServicePassword = attr.(string)
	}
	if attr, ok := d.GetOk("nova_metadata_service_tenant"); ok {
		o.NovaMetadataServiceTenant = attr.(string)
	}
	if attr, ok := d.GetOk("nova_metadata_service_username"); ok {
		o.NovaMetadataServiceUsername = attr.(string)
	}
	if attr, ok := d.GetOk("nova_metadata_shared_secret"); ok {
		o.NovaMetadataSharedSecret = attr.(string)
	}
	if attr, ok := d.GetOk("nova_os_keystone_username"); ok {
		o.NovaOSKeystoneUsername = attr.(string)
	}
	if attr, ok := d.GetOk("nova_project_domain_name"); ok {
		o.NovaProjectDomainName = attr.(string)
	}
	if attr, ok := d.GetOk("nova_project_name"); ok {
		o.NovaProjectName = attr.(string)
	}
	if attr, ok := d.GetOk("nova_region_name"); ok {
		o.NovaRegionName = attr.(string)
	}
	if attr, ok := d.GetOk("nova_user_domain_name"); ok {
		o.NovaUserDomainName = attr.(string)
	}
	if attr, ok := d.GetOk("upgrade_package_password"); ok {
		o.UpgradePackagePassword = attr.(string)
	}
	if attr, ok := d.GetOk("upgrade_package_url"); ok {
		o.UpgradePackageURL = attr.(string)
	}
	if attr, ok := d.GetOk("upgrade_package_username"); ok {
		o.UpgradePackageUsername = attr.(string)
	}
	if attr, ok := d.GetOk("upgrade_script_time_limit"); ok {
		UpgradeScriptTimeLimit := attr.(int)
		o.UpgradeScriptTimeLimit = &UpgradeScriptTimeLimit
	}
	if attr, ok := d.GetOk("cpu_count"); ok {
		o.CpuCount = attr.(string)
	}
	if attr, ok := d.GetOk("primary_data_uplink_underlay_id"); ok {
		PrimaryDataUplinkUnderlayID := attr.(int)
		o.PrimaryDataUplinkUnderlayID = &PrimaryDataUplinkUnderlayID
	}
	if attr, ok := d.GetOk("primary_nuage_controller"); ok {
		o.PrimaryNuageController = attr.(string)
	}
	if attr, ok := d.GetOk("vrs_password"); ok {
		o.VrsPassword = attr.(string)
	}
	if attr, ok := d.GetOk("vrs_user_name"); ok {
		o.VrsUserName = attr.(string)
	}
	if attr, ok := d.GetOk("assoc_vcenter_data_center_id"); ok {
		o.AssocVCenterDataCenterID = attr.(string)
	}
	if attr, ok := d.GetOk("assoc_vcenter_id"); ok {
		o.AssocVCenterID = attr.(string)
	}
	if attr, ok := d.GetOk("static_route"); ok {
		o.StaticRoute = attr.(string)
	}
	if attr, ok := d.GetOk("static_route_gateway"); ok {
		o.StaticRouteGateway = attr.(string)
	}
	if attr, ok := d.GetOk("static_route_netmask"); ok {
		o.StaticRouteNetmask = attr.(string)
	}
	if attr, ok := d.GetOk("ntp_server1"); ok {
		o.NtpServer1 = attr.(string)
	}
	if attr, ok := d.GetOk("ntp_server2"); ok {
		o.NtpServer2 = attr.(string)
	}
	if attr, ok := d.GetOk("mtu"); ok {
		Mtu := attr.(int)
		o.Mtu = &Mtu
	}
	if attr, ok := d.GetOk("multi_vmssupport"); ok {
		MultiVMSsupport := attr.(bool)
		o.MultiVMSsupport = &MultiVMSsupport
	}
	if attr, ok := d.GetOk("multicast_receive_interface"); ok {
		o.MulticastReceiveInterface = attr.(string)
	}
	if attr, ok := d.GetOk("multicast_receive_interface_ip"); ok {
		o.MulticastReceiveInterfaceIP = attr.(string)
	}
	if attr, ok := d.GetOk("multicast_receive_interface_netmask"); ok {
		o.MulticastReceiveInterfaceNetmask = attr.(string)
	}
	if attr, ok := d.GetOk("multicast_receive_range"); ok {
		o.MulticastReceiveRange = attr.(string)
	}
	if attr, ok := d.GetOk("multicast_send_interface"); ok {
		o.MulticastSendInterface = attr.(string)
	}
	if attr, ok := d.GetOk("multicast_send_interface_ip"); ok {
		o.MulticastSendInterfaceIP = attr.(string)
	}
	if attr, ok := d.GetOk("multicast_send_interface_netmask"); ok {
		o.MulticastSendInterfaceNetmask = attr.(string)
	}
	if attr, ok := d.GetOk("multicast_source_portgroup"); ok {
		o.MulticastSourcePortgroup = attr.(string)
	}
	if attr, ok := d.GetOk("customized_script_url"); ok {
		o.CustomizedScriptURL = attr.(string)
	}
	if attr, ok := d.GetOk("ovf_url"); ok {
		o.OvfURL = attr.(string)
	}
	if attr, ok := d.GetOk("avrs_enabled"); ok {
		AvrsEnabled := attr.(bool)
		o.AvrsEnabled = &AvrsEnabled
	}
	if attr, ok := d.GetOk("avrs_profile"); ok {
		o.AvrsProfile = attr.(string)
	}
	if attr, ok := d.GetOk("external_id"); ok {
		o.ExternalID = attr.(string)
	}
	parent := &vspk.VCenterDataCenter{ID: d.Get("parent_vcenter_data_center").(string)}
	err := parent.CreateVCenterCluster(o)
	if err != nil {
		return err
	}

	d.SetId(o.Identifier())
	return resourceVCenterClusterRead(d, m)
}

func resourceVCenterClusterRead(d *schema.ResourceData, m interface{}) error {
	o := &vspk.VCenterCluster{
		ID: d.Id(),
	}

	err := o.Fetch()
	if err != nil {
		d.SetId("")
		return nil
	}

	d.Set("vrs_configuration_time_limit", o.VRSConfigurationTimeLimit)
	d.Set("v_require_nuage_metadata", o.VRequireNuageMetadata)
	d.Set("name", o.Name)
	d.Set("managed_object_id", o.ManagedObjectID)
	d.Set("last_updated_by", o.LastUpdatedBy)
	d.Set("data_dns1", o.DataDNS1)
	d.Set("data_dns2", o.DataDNS2)
	d.Set("data_gateway", o.DataGateway)
	d.Set("data_network_portgroup", o.DataNetworkPortgroup)
	d.Set("datapath_sync_timeout", o.DatapathSyncTimeout)
	d.Set("scope", o.Scope)
	d.Set("secondary_data_uplink_dhcp_enabled", o.SecondaryDataUplinkDHCPEnabled)
	d.Set("secondary_data_uplink_enabled", o.SecondaryDataUplinkEnabled)
	d.Set("secondary_data_uplink_interface", o.SecondaryDataUplinkInterface)
	d.Set("secondary_data_uplink_mtu", o.SecondaryDataUplinkMTU)
	d.Set("secondary_data_uplink_primary_controller", o.SecondaryDataUplinkPrimaryController)
	d.Set("secondary_data_uplink_secondary_controller", o.SecondaryDataUplinkSecondaryController)
	d.Set("secondary_data_uplink_underlay_id", o.SecondaryDataUplinkUnderlayID)
	d.Set("secondary_nuage_controller", o.SecondaryNuageController)
	d.Set("deleted_from_vcenter_data_center", o.DeletedFromVCenterDataCenter)
	d.Set("memory_size_in_gb", o.MemorySizeInGB)
	d.Set("remote_syslog_server_ip", o.RemoteSyslogServerIP)
	d.Set("remote_syslog_server_port", o.RemoteSyslogServerPort)
	d.Set("remote_syslog_server_type", o.RemoteSyslogServerType)
	d.Set("generic_split_activation", o.GenericSplitActivation)
	d.Set("separate_data_network", o.SeparateDataNetwork)
	d.Set("personality", o.Personality)
	d.Set("description", o.Description)
	d.Set("destination_mirror_port", o.DestinationMirrorPort)
	d.Set("metadata_server_ip", o.MetadataServerIP)
	d.Set("metadata_server_listen_port", o.MetadataServerListenPort)
	d.Set("metadata_server_port", o.MetadataServerPort)
	d.Set("metadata_service_enabled", o.MetadataServiceEnabled)
	d.Set("network_uplink_interface", o.NetworkUplinkInterface)
	d.Set("network_uplink_interface_gateway", o.NetworkUplinkInterfaceGateway)
	d.Set("network_uplink_interface_ip", o.NetworkUplinkInterfaceIp)
	d.Set("network_uplink_interface_netmask", o.NetworkUplinkInterfaceNetmask)
	d.Set("revertive_controller_enabled", o.RevertiveControllerEnabled)
	d.Set("revertive_timer", o.RevertiveTimer)
	d.Set("nfs_log_server", o.NfsLogServer)
	d.Set("nfs_mount_path", o.NfsMountPath)
	d.Set("mgmt_dns1", o.MgmtDNS1)
	d.Set("mgmt_dns2", o.MgmtDNS2)
	d.Set("mgmt_gateway", o.MgmtGateway)
	d.Set("mgmt_network_portgroup", o.MgmtNetworkPortgroup)
	d.Set("dhcp_relay_server", o.DhcpRelayServer)
	d.Set("mirror_network_portgroup", o.MirrorNetworkPortgroup)
	d.Set("disable_gro_on_datapath", o.DisableGROOnDatapath)
	d.Set("disable_lro_on_datapath", o.DisableLROOnDatapath)
	d.Set("site_id", o.SiteId)
	d.Set("allow_data_dhcp", o.AllowDataDHCP)
	d.Set("allow_mgmt_dhcp", o.AllowMgmtDHCP)
	d.Set("flow_eviction_threshold", o.FlowEvictionThreshold)
	d.Set("vm_network_portgroup", o.VmNetworkPortgroup)
	d.Set("enable_vrs_resource_reservation", o.EnableVRSResourceReservation)
	d.Set("entity_scope", o.EntityScope)
	d.Set("configured_metrics_push_interval", o.ConfiguredMetricsPushInterval)
	d.Set("portgroup_metadata", o.PortgroupMetadata)
	d.Set("nova_client_version", o.NovaClientVersion)
	d.Set("nova_identity_url_version", o.NovaIdentityURLVersion)
	d.Set("nova_metadata_service_auth_url", o.NovaMetadataServiceAuthUrl)
	d.Set("nova_metadata_service_endpoint", o.NovaMetadataServiceEndpoint)
	d.Set("nova_metadata_service_password", o.NovaMetadataServicePassword)
	d.Set("nova_metadata_service_tenant", o.NovaMetadataServiceTenant)
	d.Set("nova_metadata_service_username", o.NovaMetadataServiceUsername)
	d.Set("nova_metadata_shared_secret", o.NovaMetadataSharedSecret)
	d.Set("nova_os_keystone_username", o.NovaOSKeystoneUsername)
	d.Set("nova_project_domain_name", o.NovaProjectDomainName)
	d.Set("nova_project_name", o.NovaProjectName)
	d.Set("nova_region_name", o.NovaRegionName)
	d.Set("nova_user_domain_name", o.NovaUserDomainName)
	d.Set("upgrade_package_password", o.UpgradePackagePassword)
	d.Set("upgrade_package_url", o.UpgradePackageURL)
	d.Set("upgrade_package_username", o.UpgradePackageUsername)
	d.Set("upgrade_script_time_limit", o.UpgradeScriptTimeLimit)
	d.Set("cpu_count", o.CpuCount)
	d.Set("primary_data_uplink_underlay_id", o.PrimaryDataUplinkUnderlayID)
	d.Set("primary_nuage_controller", o.PrimaryNuageController)
	d.Set("vrs_password", o.VrsPassword)
	d.Set("vrs_user_name", o.VrsUserName)
	d.Set("assoc_vcenter_data_center_id", o.AssocVCenterDataCenterID)
	d.Set("assoc_vcenter_id", o.AssocVCenterID)
	d.Set("static_route", o.StaticRoute)
	d.Set("static_route_gateway", o.StaticRouteGateway)
	d.Set("static_route_netmask", o.StaticRouteNetmask)
	d.Set("ntp_server1", o.NtpServer1)
	d.Set("ntp_server2", o.NtpServer2)
	d.Set("mtu", o.Mtu)
	d.Set("multi_vmssupport", o.MultiVMSsupport)
	d.Set("multicast_receive_interface", o.MulticastReceiveInterface)
	d.Set("multicast_receive_interface_ip", o.MulticastReceiveInterfaceIP)
	d.Set("multicast_receive_interface_netmask", o.MulticastReceiveInterfaceNetmask)
	d.Set("multicast_receive_range", o.MulticastReceiveRange)
	d.Set("multicast_send_interface", o.MulticastSendInterface)
	d.Set("multicast_send_interface_ip", o.MulticastSendInterfaceIP)
	d.Set("multicast_send_interface_netmask", o.MulticastSendInterfaceNetmask)
	d.Set("multicast_source_portgroup", o.MulticastSourcePortgroup)
	d.Set("customized_script_url", o.CustomizedScriptURL)
	d.Set("ovf_url", o.OvfURL)
	d.Set("avrs_enabled", o.AvrsEnabled)
	d.Set("avrs_profile", o.AvrsProfile)
	d.Set("external_id", o.ExternalID)

	d.Set("id", o.Identifier())
	d.Set("parent_id", o.ParentID)
	d.Set("parent_type", o.ParentType)
	d.Set("owner", o.Owner)

	return nil
}

func resourceVCenterClusterUpdate(d *schema.ResourceData, m interface{}) error {
	o := &vspk.VCenterCluster{
		ID: d.Id(),
	}

	err := o.Fetch()
	if err != nil {
		return err
	}

	o.Name = d.Get("name").(string)
	SecondaryDataUplinkEnabled := d.Get("secondary_data_uplink_enabled").(bool)
	o.SecondaryDataUplinkEnabled = &SecondaryDataUplinkEnabled
	RevertiveControllerEnabled := d.Get("revertive_controller_enabled").(bool)
	o.RevertiveControllerEnabled = &RevertiveControllerEnabled
	RevertiveTimer := d.Get("revertive_timer").(int)
	o.RevertiveTimer = &RevertiveTimer

	if attr, ok := d.GetOk("vrs_configuration_time_limit"); ok {
		VRSConfigurationTimeLimit := attr.(int)
		o.VRSConfigurationTimeLimit = &VRSConfigurationTimeLimit
	}
	if attr, ok := d.GetOk("v_require_nuage_metadata"); ok {
		VRequireNuageMetadata := attr.(bool)
		o.VRequireNuageMetadata = &VRequireNuageMetadata
	}
	if attr, ok := d.GetOk("managed_object_id"); ok {
		o.ManagedObjectID = attr.(string)
	}
	if attr, ok := d.GetOk("data_dns1"); ok {
		o.DataDNS1 = attr.(string)
	}
	if attr, ok := d.GetOk("data_dns2"); ok {
		o.DataDNS2 = attr.(string)
	}
	if attr, ok := d.GetOk("data_gateway"); ok {
		o.DataGateway = attr.(string)
	}
	if attr, ok := d.GetOk("data_network_portgroup"); ok {
		o.DataNetworkPortgroup = attr.(string)
	}
	if attr, ok := d.GetOk("datapath_sync_timeout"); ok {
		DatapathSyncTimeout := attr.(int)
		o.DatapathSyncTimeout = &DatapathSyncTimeout
	}
	if attr, ok := d.GetOk("scope"); ok {
		Scope := attr.(bool)
		o.Scope = &Scope
	}
	if attr, ok := d.GetOk("secondary_data_uplink_dhcp_enabled"); ok {
		SecondaryDataUplinkDHCPEnabled := attr.(bool)
		o.SecondaryDataUplinkDHCPEnabled = &SecondaryDataUplinkDHCPEnabled
	}
	if attr, ok := d.GetOk("secondary_data_uplink_interface"); ok {
		o.SecondaryDataUplinkInterface = attr.(string)
	}
	if attr, ok := d.GetOk("secondary_data_uplink_mtu"); ok {
		SecondaryDataUplinkMTU := attr.(int)
		o.SecondaryDataUplinkMTU = &SecondaryDataUplinkMTU
	}
	if attr, ok := d.GetOk("secondary_data_uplink_primary_controller"); ok {
		o.SecondaryDataUplinkPrimaryController = attr.(string)
	}
	if attr, ok := d.GetOk("secondary_data_uplink_secondary_controller"); ok {
		o.SecondaryDataUplinkSecondaryController = attr.(string)
	}
	if attr, ok := d.GetOk("secondary_data_uplink_underlay_id"); ok {
		SecondaryDataUplinkUnderlayID := attr.(int)
		o.SecondaryDataUplinkUnderlayID = &SecondaryDataUplinkUnderlayID
	}
	if attr, ok := d.GetOk("secondary_nuage_controller"); ok {
		o.SecondaryNuageController = attr.(string)
	}
	if attr, ok := d.GetOk("deleted_from_vcenter_data_center"); ok {
		DeletedFromVCenterDataCenter := attr.(bool)
		o.DeletedFromVCenterDataCenter = &DeletedFromVCenterDataCenter
	}
	if attr, ok := d.GetOk("memory_size_in_gb"); ok {
		o.MemorySizeInGB = attr.(string)
	}
	if attr, ok := d.GetOk("remote_syslog_server_ip"); ok {
		o.RemoteSyslogServerIP = attr.(string)
	}
	if attr, ok := d.GetOk("remote_syslog_server_port"); ok {
		RemoteSyslogServerPort := attr.(int)
		o.RemoteSyslogServerPort = &RemoteSyslogServerPort
	}
	if attr, ok := d.GetOk("remote_syslog_server_type"); ok {
		o.RemoteSyslogServerType = attr.(string)
	}
	if attr, ok := d.GetOk("generic_split_activation"); ok {
		GenericSplitActivation := attr.(bool)
		o.GenericSplitActivation = &GenericSplitActivation
	}
	if attr, ok := d.GetOk("separate_data_network"); ok {
		SeparateDataNetwork := attr.(bool)
		o.SeparateDataNetwork = &SeparateDataNetwork
	}
	if attr, ok := d.GetOk("personality"); ok {
		o.Personality = attr.(string)
	}
	if attr, ok := d.GetOk("description"); ok {
		o.Description = attr.(string)
	}
	if attr, ok := d.GetOk("destination_mirror_port"); ok {
		o.DestinationMirrorPort = attr.(string)
	}
	if attr, ok := d.GetOk("metadata_server_ip"); ok {
		o.MetadataServerIP = attr.(string)
	}
	if attr, ok := d.GetOk("metadata_server_listen_port"); ok {
		MetadataServerListenPort := attr.(int)
		o.MetadataServerListenPort = &MetadataServerListenPort
	}
	if attr, ok := d.GetOk("metadata_server_port"); ok {
		MetadataServerPort := attr.(int)
		o.MetadataServerPort = &MetadataServerPort
	}
	if attr, ok := d.GetOk("metadata_service_enabled"); ok {
		MetadataServiceEnabled := attr.(bool)
		o.MetadataServiceEnabled = &MetadataServiceEnabled
	}
	if attr, ok := d.GetOk("network_uplink_interface"); ok {
		o.NetworkUplinkInterface = attr.(string)
	}
	if attr, ok := d.GetOk("network_uplink_interface_gateway"); ok {
		o.NetworkUplinkInterfaceGateway = attr.(string)
	}
	if attr, ok := d.GetOk("network_uplink_interface_ip"); ok {
		o.NetworkUplinkInterfaceIp = attr.(string)
	}
	if attr, ok := d.GetOk("network_uplink_interface_netmask"); ok {
		o.NetworkUplinkInterfaceNetmask = attr.(string)
	}
	if attr, ok := d.GetOk("nfs_log_server"); ok {
		o.NfsLogServer = attr.(string)
	}
	if attr, ok := d.GetOk("nfs_mount_path"); ok {
		o.NfsMountPath = attr.(string)
	}
	if attr, ok := d.GetOk("mgmt_dns1"); ok {
		o.MgmtDNS1 = attr.(string)
	}
	if attr, ok := d.GetOk("mgmt_dns2"); ok {
		o.MgmtDNS2 = attr.(string)
	}
	if attr, ok := d.GetOk("mgmt_gateway"); ok {
		o.MgmtGateway = attr.(string)
	}
	if attr, ok := d.GetOk("mgmt_network_portgroup"); ok {
		o.MgmtNetworkPortgroup = attr.(string)
	}
	if attr, ok := d.GetOk("dhcp_relay_server"); ok {
		o.DhcpRelayServer = attr.(string)
	}
	if attr, ok := d.GetOk("mirror_network_portgroup"); ok {
		o.MirrorNetworkPortgroup = attr.(string)
	}
	if attr, ok := d.GetOk("disable_gro_on_datapath"); ok {
		DisableGROOnDatapath := attr.(bool)
		o.DisableGROOnDatapath = &DisableGROOnDatapath
	}
	if attr, ok := d.GetOk("disable_lro_on_datapath"); ok {
		DisableLROOnDatapath := attr.(bool)
		o.DisableLROOnDatapath = &DisableLROOnDatapath
	}
	if attr, ok := d.GetOk("site_id"); ok {
		o.SiteId = attr.(string)
	}
	if attr, ok := d.GetOk("allow_data_dhcp"); ok {
		AllowDataDHCP := attr.(bool)
		o.AllowDataDHCP = &AllowDataDHCP
	}
	if attr, ok := d.GetOk("allow_mgmt_dhcp"); ok {
		AllowMgmtDHCP := attr.(bool)
		o.AllowMgmtDHCP = &AllowMgmtDHCP
	}
	if attr, ok := d.GetOk("flow_eviction_threshold"); ok {
		FlowEvictionThreshold := attr.(int)
		o.FlowEvictionThreshold = &FlowEvictionThreshold
	}
	if attr, ok := d.GetOk("vm_network_portgroup"); ok {
		o.VmNetworkPortgroup = attr.(string)
	}
	if attr, ok := d.GetOk("enable_vrs_resource_reservation"); ok {
		EnableVRSResourceReservation := attr.(bool)
		o.EnableVRSResourceReservation = &EnableVRSResourceReservation
	}
	if attr, ok := d.GetOk("configured_metrics_push_interval"); ok {
		ConfiguredMetricsPushInterval := attr.(int)
		o.ConfiguredMetricsPushInterval = &ConfiguredMetricsPushInterval
	}
	if attr, ok := d.GetOk("portgroup_metadata"); ok {
		PortgroupMetadata := attr.(bool)
		o.PortgroupMetadata = &PortgroupMetadata
	}
	if attr, ok := d.GetOk("nova_client_version"); ok {
		NovaClientVersion := attr.(int)
		o.NovaClientVersion = &NovaClientVersion
	}
	if attr, ok := d.GetOk("nova_identity_url_version"); ok {
		o.NovaIdentityURLVersion = attr.(string)
	}
	if attr, ok := d.GetOk("nova_metadata_service_auth_url"); ok {
		o.NovaMetadataServiceAuthUrl = attr.(string)
	}
	if attr, ok := d.GetOk("nova_metadata_service_endpoint"); ok {
		o.NovaMetadataServiceEndpoint = attr.(string)
	}
	if attr, ok := d.GetOk("nova_metadata_service_password"); ok {
		o.NovaMetadataServicePassword = attr.(string)
	}
	if attr, ok := d.GetOk("nova_metadata_service_tenant"); ok {
		o.NovaMetadataServiceTenant = attr.(string)
	}
	if attr, ok := d.GetOk("nova_metadata_service_username"); ok {
		o.NovaMetadataServiceUsername = attr.(string)
	}
	if attr, ok := d.GetOk("nova_metadata_shared_secret"); ok {
		o.NovaMetadataSharedSecret = attr.(string)
	}
	if attr, ok := d.GetOk("nova_os_keystone_username"); ok {
		o.NovaOSKeystoneUsername = attr.(string)
	}
	if attr, ok := d.GetOk("nova_project_domain_name"); ok {
		o.NovaProjectDomainName = attr.(string)
	}
	if attr, ok := d.GetOk("nova_project_name"); ok {
		o.NovaProjectName = attr.(string)
	}
	if attr, ok := d.GetOk("nova_region_name"); ok {
		o.NovaRegionName = attr.(string)
	}
	if attr, ok := d.GetOk("nova_user_domain_name"); ok {
		o.NovaUserDomainName = attr.(string)
	}
	if attr, ok := d.GetOk("upgrade_package_password"); ok {
		o.UpgradePackagePassword = attr.(string)
	}
	if attr, ok := d.GetOk("upgrade_package_url"); ok {
		o.UpgradePackageURL = attr.(string)
	}
	if attr, ok := d.GetOk("upgrade_package_username"); ok {
		o.UpgradePackageUsername = attr.(string)
	}
	if attr, ok := d.GetOk("upgrade_script_time_limit"); ok {
		UpgradeScriptTimeLimit := attr.(int)
		o.UpgradeScriptTimeLimit = &UpgradeScriptTimeLimit
	}
	if attr, ok := d.GetOk("cpu_count"); ok {
		o.CpuCount = attr.(string)
	}
	if attr, ok := d.GetOk("primary_data_uplink_underlay_id"); ok {
		PrimaryDataUplinkUnderlayID := attr.(int)
		o.PrimaryDataUplinkUnderlayID = &PrimaryDataUplinkUnderlayID
	}
	if attr, ok := d.GetOk("primary_nuage_controller"); ok {
		o.PrimaryNuageController = attr.(string)
	}
	if attr, ok := d.GetOk("vrs_password"); ok {
		o.VrsPassword = attr.(string)
	}
	if attr, ok := d.GetOk("vrs_user_name"); ok {
		o.VrsUserName = attr.(string)
	}
	if attr, ok := d.GetOk("assoc_vcenter_data_center_id"); ok {
		o.AssocVCenterDataCenterID = attr.(string)
	}
	if attr, ok := d.GetOk("assoc_vcenter_id"); ok {
		o.AssocVCenterID = attr.(string)
	}
	if attr, ok := d.GetOk("static_route"); ok {
		o.StaticRoute = attr.(string)
	}
	if attr, ok := d.GetOk("static_route_gateway"); ok {
		o.StaticRouteGateway = attr.(string)
	}
	if attr, ok := d.GetOk("static_route_netmask"); ok {
		o.StaticRouteNetmask = attr.(string)
	}
	if attr, ok := d.GetOk("ntp_server1"); ok {
		o.NtpServer1 = attr.(string)
	}
	if attr, ok := d.GetOk("ntp_server2"); ok {
		o.NtpServer2 = attr.(string)
	}
	if attr, ok := d.GetOk("mtu"); ok {
		Mtu := attr.(int)
		o.Mtu = &Mtu
	}
	if attr, ok := d.GetOk("multi_vmssupport"); ok {
		MultiVMSsupport := attr.(bool)
		o.MultiVMSsupport = &MultiVMSsupport
	}
	if attr, ok := d.GetOk("multicast_receive_interface"); ok {
		o.MulticastReceiveInterface = attr.(string)
	}
	if attr, ok := d.GetOk("multicast_receive_interface_ip"); ok {
		o.MulticastReceiveInterfaceIP = attr.(string)
	}
	if attr, ok := d.GetOk("multicast_receive_interface_netmask"); ok {
		o.MulticastReceiveInterfaceNetmask = attr.(string)
	}
	if attr, ok := d.GetOk("multicast_receive_range"); ok {
		o.MulticastReceiveRange = attr.(string)
	}
	if attr, ok := d.GetOk("multicast_send_interface"); ok {
		o.MulticastSendInterface = attr.(string)
	}
	if attr, ok := d.GetOk("multicast_send_interface_ip"); ok {
		o.MulticastSendInterfaceIP = attr.(string)
	}
	if attr, ok := d.GetOk("multicast_send_interface_netmask"); ok {
		o.MulticastSendInterfaceNetmask = attr.(string)
	}
	if attr, ok := d.GetOk("multicast_source_portgroup"); ok {
		o.MulticastSourcePortgroup = attr.(string)
	}
	if attr, ok := d.GetOk("customized_script_url"); ok {
		o.CustomizedScriptURL = attr.(string)
	}
	if attr, ok := d.GetOk("ovf_url"); ok {
		o.OvfURL = attr.(string)
	}
	if attr, ok := d.GetOk("avrs_enabled"); ok {
		AvrsEnabled := attr.(bool)
		o.AvrsEnabled = &AvrsEnabled
	}
	if attr, ok := d.GetOk("avrs_profile"); ok {
		o.AvrsProfile = attr.(string)
	}
	if attr, ok := d.GetOk("external_id"); ok {
		o.ExternalID = attr.(string)
	}

	o.Save()

	return nil
}

func resourceVCenterClusterDelete(d *schema.ResourceData, m interface{}) error {
	o := &vspk.VCenterCluster{
		ID: d.Id(),
	}

	err := o.Delete()
	if err != nil {
		return err
	}

	return nil
}
