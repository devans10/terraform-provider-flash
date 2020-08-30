# Protection Group

Provides a Pure Storage Protection Group resource.

## Example Usage

```sh
resource "purestorage_protectiongroup" "example" {
  provider = flash
  name     = "example"
}
```

## Argument Reference

The following arguments are supported:

+ `name` - (Required) Name of the protection group.
+ `hosts` - (Optional) List of hosts in protection group. Conflicts with `volumes` and `hgroups`.
+ `volumes` - (Optional) List of volumes in protection group. Conflicts with `hosts` and `hgroups`.
+ `hgroups` - (Optional) List of hostgroups in the protection group. Conflicts with `hosts` and `volumes`.
+ `targets` - (Optional) List of replication targets for the protection group.
+ `all_for` - (Optional) The retention policy of the protection group. Specifies the length of time to keep the snapshots on the source array before they are eradicated.
+ `days` - (Optional) The retention policy of the protection group. Specifies the number of days to keep the per_day snapshots beyond the all_for period before they are eradicated.
+ `per_day` - (Optional) the retention policy of the protection group. Specifies the number of per_day snapshots to keep beyond the all_for period.
+ `replicate_at` - (Optional) the replication schedule of the protection group. Specifies the preferred time, on the hour, at which to replicate the snapshots.
+ `replicate_blackout` - (Optional) the replication schedule of the protection group. Specifies the range of time at which to suspend replication.
+ `replicate_enabled` - (Optional) Used to enable (true) or disable (false) the protection group replication schedule.
+ `replicate_frequency` - (Optional) the replication schedule of the protection group. Specifies the replication frequency.
+ `snap_at` - (Optional) the snapshot schedule of the protection group. Specifies the preferred time, on the hour, at which to generate the snapshot.
+ `snap_enabled` - (Optional) Used to enable (true) or disable (false) the protection group snapshot schedule.
+ `snap_frequency` - (Optional) Modifies the snapshot schedule of the protection group. Specifies the snapshot frequency.
+ `target_all_for` - (Optional) Modifies the retention policy of the protection group. Specifies the length of time to keep the replicated snapshots on the targets.
+ `target_days` - (Optional) Modifies the retention policy of the protection group. Specifies the number of days to keep the target_per_day replicated snapshots beyond the target_all_for period before they are eradicated.
+ `target_per_day` - (Optional) Modifies the retention policy of the protection group. Specifies the number of per_day replicated snapshots to keep beyond the target_all_for period.

## Attribute Reference

The following attributes are exported:

+ `name` - Name of the protection group.
+ `hosts` - List of hosts in protection group. Conflicts with `volumes` and `hgroups`.
+ `volumes` - List of volumes in protection group. Conflicts with `hosts` and `hgroups`.
+ `hgroups` - List of hostgroups in the protection group. Conflicts with `hosts` and `volumes`.
+ `source` - The source protection group
+ `targets` - List of replication targets for the protection group.
+ `all_for` - The retention policy of the protection group. Specifies the length of time to keep the snapshots on the source array before they are eradicated.
+ `days` - The retention policy of the protection group. Specifies the number of days to keep the per_day snapshots beyond the all_for period before they are eradicated.
+ `per_day` - the retention policy of the protection group. Specifies the number of per_day snapshots to keep beyond the all_for period.
+ `replicate_at` - the replication schedule of the protection group. Specifies the preferred time, on the hour, at which to replicate the snapshots.
+ `replicate_blackout` - the replication schedule of the protection group. Specifies the range of time at which to suspend replication.
+ `replicate_enabled` - Used to enable (true) or disable (false) the protection group replication schedule.
+ `replicate_frequency` - the replication schedule of the protection group. Specifies the replication frequency.
+ `snap_at` - the snapshot schedule of the protection group. Specifies the preferred time, on the hour, at which to generate the snapshot.
+ `snap_enabled` - Used to enable (true) or disable (false) the protection group snapshot schedule.
+ `snap_frequency` - Modifies the snapshot schedule of the protection group. Specifies the snapshot frequency.
+ `target_all_for` - Modifies the retention policy of the protection group. Specifies the length of time to keep the replicated snapshots on the targets.
+ `target_days` - Modifies the retention policy of the protection group. Specifies the number of days to keep the target_per_day replicated snapshots beyond the target_all_for period before they are eradicated.
+ `target_per_day` - Modifies the retention policy of the protection group. Specifies the number of per_day replicated snapshots to keep beyond the target_all_for period.

## Import

Protection groups can be imported using the Protection group name.

```sh
terraform import purestorage_protectiongroup example
```
