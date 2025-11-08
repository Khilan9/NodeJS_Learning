// Sample quickstart is a basic program that uses Secret Manager.
package main

import (
	"context"
	"fmt"
	"log"

	parametermanager "cloud.google.com/go/parametermanager/apiv1"
	parametermanagerpb "cloud.google.com/go/parametermanager/apiv1/parametermanagerpb"
	secretmanager "cloud.google.com/go/secretmanager/apiv1"
	"cloud.google.com/go/secretmanager/apiv1/secretmanagerpb"
	"google.golang.org/api/iterator"
)

var projectID string = "gsmg4-53270"
var parameterID string = "packer-nginx"
var versionID string = "2"

func createSecret() {

	// Create the client.
	ctx := context.Background()
	client, err := secretmanager.NewClient(ctx)
	if err != nil {
		log.Fatalf("failed to setup client: %v", err)
	}
	defer client.Close()

	// Create the request to create the secret.
	createSecretReq := &secretmanagerpb.CreateSecretRequest{
		Parent:   fmt.Sprintf("projects/%s", projectID),
		SecretId: "my-secret",
		Secret: &secretmanagerpb.Secret{
			Replication: &secretmanagerpb.Replication{
				Replication: &secretmanagerpb.Replication_Automatic_{
					Automatic: &secretmanagerpb.Replication_Automatic{},
				},
			},
		},
	}

	secret, err := client.CreateSecret(ctx, createSecretReq)
	if err != nil {
		log.Fatalf("failed to create secret: %v", err)
	}

	// Declare the payload to store.
	payload := []byte("my super secret data")

	// Build the request.
	addSecretVersionReq := &secretmanagerpb.AddSecretVersionRequest{
		Parent: secret.Name,
		Payload: &secretmanagerpb.SecretPayload{
			Data: payload,
		},
	}

	// Call the API.
	version, err := client.AddSecretVersion(ctx, addSecretVersionReq)
	if err != nil {
		log.Fatalf("failed to add secret version: %v", err)
	}

	// Build the request.
	accessRequest := &secretmanagerpb.AccessSecretVersionRequest{
		Name: version.Name,
	}

	// Call the API.
	result, err := client.AccessSecretVersion(ctx, accessRequest)
	if err != nil {
		log.Fatalf("failed to access secret version: %v", err)
	}

	// Print the secret payload.
	//
	// WARNING: Do not print the secret in a production environment - this
	// snippet is showing how to access the secret material.
	log.Printf("Plaintext: %s", result.Payload.Data)

}

func listParamVersions() {
	ctx := context.Background()
	paramclient, paramerr := parametermanager.NewClient(ctx)
	if paramerr != nil {
		fmt.Printf("failed to create Parameter Manager client: %w", paramerr)
	}
	defer paramclient.Close()

	// Construct the name of the list parameter.
	parent := fmt.Sprintf("projects/%s/locations/global/parameters/%s", projectID, parameterID)

	// Build the request to list parameter versions.
	req := &parametermanagerpb.ListParameterVersionsRequest{
		Parent: parent,
	}

	// Call the API to list parameter versions.
	parameterVersions := paramclient.ListParameterVersions(ctx, req)
	for {
		version, err := parameterVersions.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			fmt.Printf("failed to list parameter versions: %w", err)
		}

		fmt.Printf("Found parameter version %s with disabled state in %v\n", version.Name, version.Disabled)
	}

}

func accessParamVersion() {
	// Create a context and a Parameter Manager client.
	ctx := context.Background()
	client, err := parametermanager.NewClient(ctx)
	if err != nil {
		fmt.Printf("failed to create Parameter Manager client: %w", err)
	}
	defer client.Close()

	// Construct the name of the parameter to get the parameter version.
	name := fmt.Sprintf("projects/%s/locations/global/parameters/%s/versions/%s", projectID, parameterID, versionID)

	// Build the request to get parameter version.
	req := &parametermanagerpb.GetParameterVersionRequest{
		Name: name,
	}

	// Call the API to get parameter version.
	version, err := client.GetParameterVersion(ctx, req)
	if err != nil {
		fmt.Printf("failed to get parameter version: %w", err)
	}

	// Find more details for the Parameter Version object here:
	// https://cloud.google.com/secret-manager/parameter-manager/docs/reference/rest/v1/projects.locations.parameters.versions#ParameterVersion
	fmt.Printf("Found parameter version %s with disabled state in %v\n", version.Name, version.Disabled)
	if !version.Disabled {
		fmt.Printf("Payload: %s\n", version.Payload.Data)
	}

}

func renderParamVersion() {

	ctx := context.Background()
	client, err := parametermanager.NewClient(ctx)
	if err != nil {
		fmt.Printf("failed to create Parameter Manager client: %w", err)
	}
	defer client.Close()

	// Construct the name of the parameter version to get render data.
	name := fmt.Sprintf("projects/%s/locations/global/parameters/%s/versions/%s", projectID, parameterID, versionID)

	// Build the request to render a parameter version.
	req := &parametermanagerpb.RenderParameterVersionRequest{
		Name: name,
	}

	// Call the API to render a parameter version.
	rendered, err := client.RenderParameterVersion(ctx, req)
	if err != nil {
		fmt.Printf("failed to render parameter version: %w", err)
	}

	fmt.Printf("Rendered parameter version: %s\n", rendered.ParameterVersion)

	// If the parameter contains secret references, they will be resolved
	// and the actual secret values will be included in the rendered output.
	// Be cautious with logging or displaying this information.
	fmt.Printf("Rendered payload: %s\n", rendered.RenderedPayload)
}

// func main() {
// 	// createSecret()
// 	// listParamVersions()
// 	// accessParamVersion()
// 	renderParamVersion()
// }
