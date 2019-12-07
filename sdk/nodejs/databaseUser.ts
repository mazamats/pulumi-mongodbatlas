// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

export class DatabaseUser extends pulumi.CustomResource {
    /**
     * Get an existing DatabaseUser resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DatabaseUserState, opts?: pulumi.CustomResourceOptions): DatabaseUser {
        return new DatabaseUser(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'mongodbatlas:index/databaseUser:DatabaseUser';

    /**
     * Returns true if the given object is an instance of DatabaseUser.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DatabaseUser {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DatabaseUser.__pulumiType;
    }

    public readonly databaseName!: pulumi.Output<string>;
    public readonly password!: pulumi.Output<string | undefined>;
    public readonly projectId!: pulumi.Output<string>;
    public readonly roles!: pulumi.Output<outputs.DatabaseUserRole[]>;
    public readonly username!: pulumi.Output<string>;

    /**
     * Create a DatabaseUser resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DatabaseUserArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DatabaseUserArgs | DatabaseUserState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as DatabaseUserState | undefined;
            inputs["databaseName"] = state ? state.databaseName : undefined;
            inputs["password"] = state ? state.password : undefined;
            inputs["projectId"] = state ? state.projectId : undefined;
            inputs["roles"] = state ? state.roles : undefined;
            inputs["username"] = state ? state.username : undefined;
        } else {
            const args = argsOrState as DatabaseUserArgs | undefined;
            if (!args || args.databaseName === undefined) {
                throw new Error("Missing required property 'databaseName'");
            }
            if (!args || args.projectId === undefined) {
                throw new Error("Missing required property 'projectId'");
            }
            if (!args || args.username === undefined) {
                throw new Error("Missing required property 'username'");
            }
            inputs["databaseName"] = args ? args.databaseName : undefined;
            inputs["password"] = args ? args.password : undefined;
            inputs["projectId"] = args ? args.projectId : undefined;
            inputs["roles"] = args ? args.roles : undefined;
            inputs["username"] = args ? args.username : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(DatabaseUser.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering DatabaseUser resources.
 */
export interface DatabaseUserState {
    readonly databaseName?: pulumi.Input<string>;
    readonly password?: pulumi.Input<string>;
    readonly projectId?: pulumi.Input<string>;
    readonly roles?: pulumi.Input<pulumi.Input<inputs.DatabaseUserRole>[]>;
    readonly username?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a DatabaseUser resource.
 */
export interface DatabaseUserArgs {
    readonly databaseName: pulumi.Input<string>;
    readonly password?: pulumi.Input<string>;
    readonly projectId: pulumi.Input<string>;
    readonly roles?: pulumi.Input<pulumi.Input<inputs.DatabaseUserRole>[]>;
    readonly username: pulumi.Input<string>;
}