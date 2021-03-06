# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from . import utilities, tables

class DatabaseUser(pulumi.CustomResource):
    database_name: pulumi.Output[str]
    """
    Database on which the user has the specified role. A role on the `admin` database can include privileges that apply to the other databases.
    """
    password: pulumi.Output[str]
    """
    User's initial password. This is required to create the user but may be removed after. Password may show up in logs, and it will be stored in the state file as plain-text. Password can be changed in the web interface to increase security.
    """
    project_id: pulumi.Output[str]
    """
    The unique ID for the project to create the database user.
    """
    roles: pulumi.Output[list]
    """
    List of user’s roles and the databases / collections on which the roles apply. A role allows the user to perform particular actions on the specified database. A role on the admin database can include privileges that apply to the other databases as well. See Roles below for more details.
    
      * `collectionName` (`str`) - Collection for which the role applies. You can specify a collection for the `read` and `readWrite` roles. If you do not specify a collection for `read` and `readWrite`, the role applies to all collections in the database (excluding some collections in the `system`. database).
      * `database_name` (`str`) - Database on which the user has the specified role. A role on the `admin` database can include privileges that apply to the other databases.
      * `roleName` (`str`)
    """
    username: pulumi.Output[str]
    """
    Username for authenticating to MongoDB.
    """
    def __init__(__self__, resource_name, opts=None, database_name=None, password=None, project_id=None, roles=None, username=None, __props__=None, __name__=None, __opts__=None):
        """
        Create a DatabaseUser resource with the given unique name, props, and options.
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] database_name: Database on which the user has the specified role. A role on the `admin` database can include privileges that apply to the other databases.
        :param pulumi.Input[str] password: User's initial password. This is required to create the user but may be removed after. Password may show up in logs, and it will be stored in the state file as plain-text. Password can be changed in the web interface to increase security.
        :param pulumi.Input[str] project_id: The unique ID for the project to create the database user.
        :param pulumi.Input[list] roles: List of user’s roles and the databases / collections on which the roles apply. A role allows the user to perform particular actions on the specified database. A role on the admin database can include privileges that apply to the other databases as well. See Roles below for more details.
        :param pulumi.Input[str] username: Username for authenticating to MongoDB.
        
        The **roles** object supports the following:
        
          * `collectionName` (`pulumi.Input[str]`) - Collection for which the role applies. You can specify a collection for the `read` and `readWrite` roles. If you do not specify a collection for `read` and `readWrite`, the role applies to all collections in the database (excluding some collections in the `system`. database).
          * `database_name` (`pulumi.Input[str]`) - Database on which the user has the specified role. A role on the `admin` database can include privileges that apply to the other databases.
          * `roleName` (`pulumi.Input[str]`)

        > This content is derived from https://github.com/terraform-providers/terraform-provider-mongodbatlas/blob/master/website/docs/r/database_user.html.markdown.
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            if database_name is None:
                raise TypeError("Missing required property 'database_name'")
            __props__['database_name'] = database_name
            __props__['password'] = password
            if project_id is None:
                raise TypeError("Missing required property 'project_id'")
            __props__['project_id'] = project_id
            __props__['roles'] = roles
            if username is None:
                raise TypeError("Missing required property 'username'")
            __props__['username'] = username
        super(DatabaseUser, __self__).__init__(
            'mongodbatlas:index/databaseUser:DatabaseUser',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, database_name=None, password=None, project_id=None, roles=None, username=None):
        """
        Get an existing DatabaseUser resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.
        
        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] database_name: Database on which the user has the specified role. A role on the `admin` database can include privileges that apply to the other databases.
        :param pulumi.Input[str] password: User's initial password. This is required to create the user but may be removed after. Password may show up in logs, and it will be stored in the state file as plain-text. Password can be changed in the web interface to increase security.
        :param pulumi.Input[str] project_id: The unique ID for the project to create the database user.
        :param pulumi.Input[list] roles: List of user’s roles and the databases / collections on which the roles apply. A role allows the user to perform particular actions on the specified database. A role on the admin database can include privileges that apply to the other databases as well. See Roles below for more details.
        :param pulumi.Input[str] username: Username for authenticating to MongoDB.
        
        The **roles** object supports the following:
        
          * `collectionName` (`pulumi.Input[str]`) - Collection for which the role applies. You can specify a collection for the `read` and `readWrite` roles. If you do not specify a collection for `read` and `readWrite`, the role applies to all collections in the database (excluding some collections in the `system`. database).
          * `database_name` (`pulumi.Input[str]`) - Database on which the user has the specified role. A role on the `admin` database can include privileges that apply to the other databases.
          * `roleName` (`pulumi.Input[str]`)

        > This content is derived from https://github.com/terraform-providers/terraform-provider-mongodbatlas/blob/master/website/docs/r/database_user.html.markdown.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()
        __props__["database_name"] = database_name
        __props__["password"] = password
        __props__["project_id"] = project_id
        __props__["roles"] = roles
        __props__["username"] = username
        return DatabaseUser(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

