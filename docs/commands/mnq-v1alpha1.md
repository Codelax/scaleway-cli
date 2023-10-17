<!-- DO NOT EDIT: this file is automatically generated using scw-doc-gen -->
# Documentation for `scw mnq-v1alpha1`
Messaging and Queuing API.
  
- [MnQ Credentials commands](#mnq-credentials-commands)
  - [Create credentials](#create-credentials)
  - [Delete credentials](#delete-credentials)
  - [Get credentials](#get-credentials)
  - [List credentials](#list-credentials)
  - [Update credentials](#update-credentials)
- [MnQ Namespace commands](#mnq-namespace-commands)
  - [Create a namespace](#create-a-namespace)
  - [Delete a namespace](#delete-a-namespace)
  - [Get a namespace](#get-a-namespace)
  - [List namespaces](#list-namespaces)
  - [Update the name of a namespace](#update-the-name-of-a-namespace)

  
## MnQ Credentials commands

MnQ Credentials commands.


### Create credentials

Create a set of credentials for a Messaging and Queuing namespace, specified by its namespace ID. If creating credentials for a NATS namespace, the `permissions` object must not be included in the request. If creating credentials for an SQS/SNS namespace, the `permissions` object is required, with all three of its child attributes.

**Usage:**

```
scw mnq credential create [arg=value ...]
```


**Args:**

| Name |   | Description |
|------|---|-------------|
| namespace-id | Required | Namespace containing the credentials |
| name | Default: `<generated>` | Name of the credentials |
| permissions.can-publish |  | Defines whether the credentials bearer can publish messages to the service (send messages to SQS queues or publish to SNS topics) |
| permissions.can-receive |  | Defines whether the credentials bearer can receive messages from the service |
| permissions.can-manage |  | Defines whether the credentials bearer can manage the associated resources (SQS queues or SNS topics or subscriptions) |
| region | Default: `fr-par`<br />One of: `fr-par` | Region to target. If none is passed will use default region from the config |



### Delete credentials

Delete a set of credentials, specified by their credential ID. Deleting credentials is irreversible and cannot be undone. The credentials can no longer be used to access the namespace.

**Usage:**

```
scw mnq credential delete <credential-id ...> [arg=value ...]
```


**Args:**

| Name |   | Description |
|------|---|-------------|
| credential-id | Required | ID of the credentials to delete |
| region | Default: `fr-par`<br />One of: `fr-par` | Region to target. If none is passed will use default region from the config |



### Get credentials

Retrieve an existing set of credentials, identified by the `credential_id`. The credentials themselves, as well as their metadata (protocol, namespace ID etc), are returned in the response.

**Usage:**

```
scw mnq credential get <credential-id ...> [arg=value ...]
```


**Args:**

| Name |   | Description |
|------|---|-------------|
| credential-id | Required | ID of the credentials to get |
| region | Default: `fr-par`<br />One of: `fr-par` | Region to target. If none is passed will use default region from the config |



### List credentials

List existing credentials in the specified region. The response contains only the metadata for the credentials, not the credentials themselves (for this, use **Get Credentials**).

**Usage:**

```
scw mnq credential list [arg=value ...]
```


**Args:**

| Name |   | Description |
|------|---|-------------|
| namespace-id |  | Namespace containing the credentials |
| order-by | One of: `id_asc`, `id_desc`, `name_asc`, `name_desc` | Order in which to return results |
| region | Default: `fr-par`<br />One of: `fr-par`, `all` | Region to target. If none is passed will use default region from the config |



### Update credentials

Update a set of credentials. You can update the credentials' name, or (in the case of SQS/SNS credentials only) their permissions. To update the name of NATS credentials, do not include the `permissions` object in your request.

**Usage:**

```
scw mnq credential update <credential-id ...> [arg=value ...]
```


**Args:**

| Name |   | Description |
|------|---|-------------|
| credential-id | Required | ID of the credentials to update |
| name |  | Name of the credentials |
| permissions.can-publish |  | Defines whether the credentials bearer can publish messages to the service (send messages to SQS queues or publish to SNS topics) |
| permissions.can-receive |  | Defines whether the credentials bearer can receive messages from the service |
| permissions.can-manage |  | Defines whether the credentials bearer can manage the associated resources (SQS queues or SNS topics or subscriptions) |
| region | Default: `fr-par`<br />One of: `fr-par` | Region to target. If none is passed will use default region from the config |



## MnQ Namespace commands

MnQ Namespace commands.


### Create a namespace

Create a Messaging and Queuing namespace, set to the desired protocol.

**Usage:**

```
scw mnq namespace create [arg=value ...]
```


**Args:**

| Name |   | Description |
|------|---|-------------|
| name | Default: `<generated>` | Namespace name |
| protocol | Required<br />One of: `unknown`, `nats`, `sqs_sns` | Namespace protocol. You must specify a valid protocol (and not `unknown`) to avoid an error. |
| project-id |  | Project ID to use. If none is passed the default project ID will be used |
| region | Default: `fr-par`<br />One of: `fr-par` | Region to target. If none is passed will use default region from the config |



### Delete a namespace

Delete a Messaging and Queuing namespace, specified by its namespace ID. Note that deleting a namespace is irreversible, and any URLs, credentials and queued messages belonging to this namespace will also be deleted.

**Usage:**

```
scw mnq namespace delete <namespace-id ...> [arg=value ...]
```


**Args:**

| Name |   | Description |
|------|---|-------------|
| namespace-id | Required | ID of the namespace to delete |
| region | Default: `fr-par`<br />One of: `fr-par` | Region to target. If none is passed will use default region from the config |



### Get a namespace

Retrieve information about an existing Messaging and Queuing namespace, identified by its namespace ID. Its full details, including name, endpoint and protocol, are returned in the response.

**Usage:**

```
scw mnq namespace get <namespace-id ...> [arg=value ...]
```


**Args:**

| Name |   | Description |
|------|---|-------------|
| namespace-id | Required | ID of the Namespace to get |
| region | Default: `fr-par`<br />One of: `fr-par` | Region to target. If none is passed will use default region from the config |



### List namespaces

List all Messaging and Queuing namespaces in the specified region, for a Scaleway Organization or Project. By default, the namespaces returned in the list are ordered by creation date in ascending order, though this can be modified via the `order_by` field.

**Usage:**

```
scw mnq namespace list [arg=value ...]
```


**Args:**

| Name |   | Description |
|------|---|-------------|
| project-id |  | Include only namespaces in this Project |
| order-by | One of: `created_at_asc`, `created_at_desc`, `updated_at_asc`, `updated_at_desc`, `id_asc`, `id_desc`, `name_asc`, `name_desc`, `project_id_asc`, `project_id_desc` | Order in which to return results |
| organization-id |  | Include only namespaces in this Organization |
| region | Default: `fr-par`<br />One of: `fr-par`, `all` | Region to target. If none is passed will use default region from the config |



### Update the name of a namespace

Update the name of a Messaging and Queuing namespace, specified by its namespace ID.

**Usage:**

```
scw mnq namespace update <namespace-id ...> [arg=value ...]
```


**Args:**

| Name |   | Description |
|------|---|-------------|
| namespace-id | Required | ID of the Namespace to update |
| name |  | Namespace name |
| region | Default: `fr-par`<br />One of: `fr-par` | Region to target. If none is passed will use default region from the config |


