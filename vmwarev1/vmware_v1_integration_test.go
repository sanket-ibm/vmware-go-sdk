// +build integration

/**
 * (C) Copyright IBM Corp. 2023.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package vmwarev1_test

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/IBM/go-sdk-core/v5/core"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.ibm.com/VMWSolutions/vmware-go-sdk/vmwarev1"
)

/**
 * This file contains an integration test for the vmwarev1 package.
 *
 * Notes:
 *
 * The integration test will automatically skip tests if the required config file is not available.
 */

var _ = Describe(`VmwareV1 Integration Tests`, func() {
	const externalConfigFile = "../vmware_v1.env"

	var (
		err          error
		vmwareService *vmwarev1.VmwareV1
		serviceURL   string
		config       map[string]string
	)

	var shouldSkipTest = func() {
		Skip("External configuration is not available, skipping tests...")
	}

	Describe(`External configuration`, func() {
		It("Successfully load the configuration", func() {
			_, err = os.Stat(externalConfigFile)
			if err != nil {
				Skip("External configuration file not found, skipping tests: " + err.Error())
			}

			os.Setenv("IBM_CREDENTIALS_FILE", externalConfigFile)
			config, err = core.GetServiceProperties(vmwarev1.DefaultServiceName)
			if err != nil {
				Skip("Error loading service properties, skipping tests: " + err.Error())
			}
			serviceURL = config["URL"]
			if serviceURL == "" {
				Skip("Unable to load service URL configuration property, skipping tests")
			}

			fmt.Fprintf(GinkgoWriter, "Service URL: %v\n", serviceURL)
			shouldSkipTest = func() {}
		})
	})

	Describe(`Client initialization`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It("Successfully construct the service client instance", func() {
			vmwareServiceOptions := &vmwarev1.VmwareV1Options{}

			vmwareService, err = vmwarev1.NewVmwareV1UsingExternalConfig(vmwareServiceOptions)
			Expect(err).To(BeNil())
			Expect(vmwareService).ToNot(BeNil())
			Expect(vmwareService.Service.Options.URL).To(Equal(serviceURL))

			core.SetLogger(core.NewLogger(core.LevelDebug, log.New(GinkgoWriter, "", log.LstdFlags), log.New(GinkgoWriter, "", log.LstdFlags)))
			vmwareService.EnableRetries(4, 30*time.Second)
		})
	})

	Describe(`CreateDirectorSites - Create a director site instance`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateDirectorSites(createDirectorSitesOptions *CreateDirectorSitesOptions)`, func() {
			fileSharesPrototypeModel := &vmwarev1.FileSharesPrototype{
				STORAGEPOINTTWOFIVEIOPSGB: core.Int64Ptr(int64(0)),
				STORAGETWOIOPSGB: core.Int64Ptr(int64(0)),
				STORAGEFOURIOPSGB: core.Int64Ptr(int64(0)),
				STORAGETENIOPSGB: core.Int64Ptr(int64(0)),
			}

			clusterPrototypeModel := &vmwarev1.ClusterPrototype{
				Name: core.StringPtr("testString"),
				HostCount: core.Int64Ptr(int64(2)),
				HostProfile: core.StringPtr("testString"),
				FileShares: fileSharesPrototypeModel,
			}

			pvdcPrototypeModel := &vmwarev1.PVDCPrototype{
				Name: core.StringPtr("testString"),
				DataCenterName: core.StringPtr("testString"),
				Clusters: []vmwarev1.ClusterPrototype{*clusterPrototypeModel},
			}

			resourceGroupIdentityModel := &vmwarev1.ResourceGroupIdentity{
				ID: core.StringPtr("testString"),
			}

			serviceIdentityModel := &vmwarev1.ServiceIdentity{
				Name: core.StringPtr("veeam"),
			}

			createDirectorSitesOptions := &vmwarev1.CreateDirectorSitesOptions{
				Name: core.StringPtr("testString"),
				Pvdcs: []vmwarev1.PVDCPrototype{*pvdcPrototypeModel},
				ResourceGroup: resourceGroupIdentityModel,
				Services: []vmwarev1.ServiceIdentity{*serviceIdentityModel},
				AcceptLanguage: core.StringPtr("testString"),
				XGlobalTransactionID: core.StringPtr("testString"),
			}

			directorSite, response, err := vmwareService.CreateDirectorSites(createDirectorSitesOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(directorSite).ToNot(BeNil())
		})
	})

	Describe(`ListDirectorSites - List director site instances`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListDirectorSites(listDirectorSitesOptions *ListDirectorSitesOptions)`, func() {
			listDirectorSitesOptions := &vmwarev1.ListDirectorSitesOptions{
				AcceptLanguage: core.StringPtr("testString"),
				XGlobalTransactionID: core.StringPtr("testString"),
			}

			directorSiteCollection, response, err := vmwareService.ListDirectorSites(listDirectorSitesOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(directorSiteCollection).ToNot(BeNil())
		})
	})

	Describe(`GetDirectorSite - Get a director site instance`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetDirectorSite(getDirectorSiteOptions *GetDirectorSiteOptions)`, func() {
			getDirectorSiteOptions := &vmwarev1.GetDirectorSiteOptions{
				ID: core.StringPtr("testString"),
				AcceptLanguage: core.StringPtr("testString"),
				XGlobalTransactionID: core.StringPtr("testString"),
			}

			directorSite, response, err := vmwareService.GetDirectorSite(getDirectorSiteOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(directorSite).ToNot(BeNil())
		})
	})

	Describe(`ListDirectorSitesPvdcs - List the provider virtual data centers in a director site instance`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListDirectorSitesPvdcs(listDirectorSitesPvdcsOptions *ListDirectorSitesPvdcsOptions)`, func() {
			listDirectorSitesPvdcsOptions := &vmwarev1.ListDirectorSitesPvdcsOptions{
				SiteID: core.StringPtr("testString"),
				AcceptLanguage: core.StringPtr("testString"),
				XGlobalTransactionID: core.StringPtr("testString"),
			}

			pvdcCollection, response, err := vmwareService.ListDirectorSitesPvdcs(listDirectorSitesPvdcsOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(pvdcCollection).ToNot(BeNil())
		})
	})

	Describe(`CreateDirectorSitesPvdcs - Create a provider virtual data center instance in a specified director site`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateDirectorSitesPvdcs(createDirectorSitesPvdcsOptions *CreateDirectorSitesPvdcsOptions)`, func() {
			fileSharesPrototypeModel := &vmwarev1.FileSharesPrototype{
				STORAGEPOINTTWOFIVEIOPSGB: core.Int64Ptr(int64(0)),
				STORAGETWOIOPSGB: core.Int64Ptr(int64(0)),
				STORAGEFOURIOPSGB: core.Int64Ptr(int64(0)),
				STORAGETENIOPSGB: core.Int64Ptr(int64(0)),
			}

			clusterPrototypeModel := &vmwarev1.ClusterPrototype{
				Name: core.StringPtr("testString"),
				HostCount: core.Int64Ptr(int64(2)),
				HostProfile: core.StringPtr("testString"),
				FileShares: fileSharesPrototypeModel,
			}

			createDirectorSitesPvdcsOptions := &vmwarev1.CreateDirectorSitesPvdcsOptions{
				SiteID: core.StringPtr("testString"),
				Name: core.StringPtr("testString"),
				DataCenterName: core.StringPtr("testString"),
				Clusters: []vmwarev1.ClusterPrototype{*clusterPrototypeModel},
				AcceptLanguage: core.StringPtr("testString"),
				XGlobalTransactionID: core.StringPtr("testString"),
			}

			pvdc, response, err := vmwareService.CreateDirectorSitesPvdcs(createDirectorSitesPvdcsOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(pvdc).ToNot(BeNil())
		})
	})

	Describe(`GetDirectorSitesPvdcs - Get the specified provider virtual data center in a director site instance`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetDirectorSitesPvdcs(getDirectorSitesPvdcsOptions *GetDirectorSitesPvdcsOptions)`, func() {
			getDirectorSitesPvdcsOptions := &vmwarev1.GetDirectorSitesPvdcsOptions{
				SiteID: core.StringPtr("testString"),
				ID: core.StringPtr("testString"),
				AcceptLanguage: core.StringPtr("testString"),
				XGlobalTransactionID: core.StringPtr("testString"),
			}

			pvdc, response, err := vmwareService.GetDirectorSitesPvdcs(getDirectorSitesPvdcsOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(pvdc).ToNot(BeNil())
		})
	})

	Describe(`ListDirectorSitesPvdcsClusters - List clusters`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListDirectorSitesPvdcsClusters(listDirectorSitesPvdcsClustersOptions *ListDirectorSitesPvdcsClustersOptions)`, func() {
			listDirectorSitesPvdcsClustersOptions := &vmwarev1.ListDirectorSitesPvdcsClustersOptions{
				SiteID: core.StringPtr("testString"),
				PvdcID: core.StringPtr("testString"),
				AcceptLanguage: core.StringPtr("testString"),
				XGlobalTransactionID: core.StringPtr("testString"),
			}

			clusterCollection, response, err := vmwareService.ListDirectorSitesPvdcsClusters(listDirectorSitesPvdcsClustersOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(clusterCollection).ToNot(BeNil())
		})
	})

	Describe(`CreateDirectorSitesPvdcsClusters - Create a cluster`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateDirectorSitesPvdcsClusters(createDirectorSitesPvdcsClustersOptions *CreateDirectorSitesPvdcsClustersOptions)`, func() {
			fileSharesPrototypeModel := &vmwarev1.FileSharesPrototype{
				STORAGEPOINTTWOFIVEIOPSGB: core.Int64Ptr(int64(0)),
				STORAGETWOIOPSGB: core.Int64Ptr(int64(0)),
				STORAGEFOURIOPSGB: core.Int64Ptr(int64(0)),
				STORAGETENIOPSGB: core.Int64Ptr(int64(0)),
			}

			createDirectorSitesPvdcsClustersOptions := &vmwarev1.CreateDirectorSitesPvdcsClustersOptions{
				SiteID: core.StringPtr("testString"),
				PvdcID: core.StringPtr("testString"),
				Name: core.StringPtr("testString"),
				HostCount: core.Int64Ptr(int64(2)),
				HostProfile: core.StringPtr("testString"),
				FileShares: fileSharesPrototypeModel,
				AcceptLanguage: core.StringPtr("testString"),
				XGlobalTransactionID: core.StringPtr("testString"),
			}

			cluster, response, err := vmwareService.CreateDirectorSitesPvdcsClusters(createDirectorSitesPvdcsClustersOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(cluster).ToNot(BeNil())
		})
	})

	Describe(`GetDirectorInstancesPvdcsCluster - Get a cluster`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetDirectorInstancesPvdcsCluster(getDirectorInstancesPvdcsClusterOptions *GetDirectorInstancesPvdcsClusterOptions)`, func() {
			getDirectorInstancesPvdcsClusterOptions := &vmwarev1.GetDirectorInstancesPvdcsClusterOptions{
				SiteID: core.StringPtr("testString"),
				ID: core.StringPtr("testString"),
				PvdcID: core.StringPtr("testString"),
				AcceptLanguage: core.StringPtr("testString"),
				XGlobalTransactionID: core.StringPtr("testString"),
			}

			cluster, response, err := vmwareService.GetDirectorInstancesPvdcsCluster(getDirectorInstancesPvdcsClusterOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(cluster).ToNot(BeNil())
		})
	})

	Describe(`UpdateDirectorSitesPvdcsCluster - Update a cluster`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UpdateDirectorSitesPvdcsCluster(updateDirectorSitesPvdcsClusterOptions *UpdateDirectorSitesPvdcsClusterOptions)`, func() {
			fileSharesPrototypeModel := &vmwarev1.FileSharesPrototype{
				STORAGEPOINTTWOFIVEIOPSGB: core.Int64Ptr(int64(0)),
				STORAGETWOIOPSGB: core.Int64Ptr(int64(0)),
				STORAGEFOURIOPSGB: core.Int64Ptr(int64(0)),
				STORAGETENIOPSGB: core.Int64Ptr(int64(0)),
			}

			clusterPatchModel := &vmwarev1.ClusterPatch{
				FileShares: fileSharesPrototypeModel,
				HostCount: core.Int64Ptr(int64(2)),
			}
			clusterPatchModelAsPatch, asPatchErr := clusterPatchModel.AsPatch()
			Expect(asPatchErr).To(BeNil())

			updateDirectorSitesPvdcsClusterOptions := &vmwarev1.UpdateDirectorSitesPvdcsClusterOptions{
				SiteID: core.StringPtr("testString"),
				ID: core.StringPtr("testString"),
				PvdcID: core.StringPtr("testString"),
				Body: clusterPatchModelAsPatch,
				AcceptLanguage: core.StringPtr("testString"),
				XGlobalTransactionID: core.StringPtr("testString"),
			}

			updateCluster, response, err := vmwareService.UpdateDirectorSitesPvdcsCluster(updateDirectorSitesPvdcsClusterOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(updateCluster).ToNot(BeNil())
		})
	})

	Describe(`ListDirectorSiteRegions - List regions`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListDirectorSiteRegions(listDirectorSiteRegionsOptions *ListDirectorSiteRegionsOptions)`, func() {
			listDirectorSiteRegionsOptions := &vmwarev1.ListDirectorSiteRegionsOptions{
				AcceptLanguage: core.StringPtr("testString"),
				XGlobalTransactionID: core.StringPtr("testString"),
			}

			directorSiteRegionCollection, response, err := vmwareService.ListDirectorSiteRegions(listDirectorSiteRegionsOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(directorSiteRegionCollection).ToNot(BeNil())
		})
	})

	Describe(`ListMultitenantDirectorSites - List multitenant director sites`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListMultitenantDirectorSites(listMultitenantDirectorSitesOptions *ListMultitenantDirectorSitesOptions)`, func() {
			listMultitenantDirectorSitesOptions := &vmwarev1.ListMultitenantDirectorSitesOptions{
				AcceptLanguage: core.StringPtr("testString"),
				XGlobalTransactionID: core.StringPtr("testString"),
			}

			multitenantDirectorSiteCollection, response, err := vmwareService.ListMultitenantDirectorSites(listMultitenantDirectorSitesOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(multitenantDirectorSiteCollection).ToNot(BeNil())
		})
	})

	Describe(`ListDirectorSiteHostProfiles - List host profiles`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListDirectorSiteHostProfiles(listDirectorSiteHostProfilesOptions *ListDirectorSiteHostProfilesOptions)`, func() {
			listDirectorSiteHostProfilesOptions := &vmwarev1.ListDirectorSiteHostProfilesOptions{
				AcceptLanguage: core.StringPtr("testString"),
				XGlobalTransactionID: core.StringPtr("testString"),
			}

			directorSiteHostProfileCollection, response, err := vmwareService.ListDirectorSiteHostProfiles(listDirectorSiteHostProfilesOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(directorSiteHostProfileCollection).ToNot(BeNil())
		})
	})

	Describe(`ListVdcs - List virtual data centers`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListVdcs(listVdcsOptions *ListVdcsOptions)`, func() {
			listVdcsOptions := &vmwarev1.ListVdcsOptions{
				AcceptLanguage: core.StringPtr("testString"),
			}

			vdcCollection, response, err := vmwareService.ListVdcs(listVdcsOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(vdcCollection).ToNot(BeNil())
		})
	})

	Describe(`CreateVdc - Create a virtual data center`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateVdc(createVdcOptions *CreateVdcOptions)`, func() {
			vdcProviderTypeModel := &vmwarev1.VDCProviderType{
				Name: core.StringPtr("paygo"),
			}

			directorSitePvdcModel := &vmwarev1.DirectorSitePVDC{
				ID: core.StringPtr("testString"),
				ProviderType: vdcProviderTypeModel,
			}

			vdcDirectorSitePrototypeModel := &vmwarev1.VDCDirectorSitePrototype{
				ID: core.StringPtr("testString"),
				Pvdc: directorSitePvdcModel,
			}

			vdcEdgePrototypeModel := &vmwarev1.VDCEdgePrototype{
				Size: core.StringPtr("medium"),
				Type: core.StringPtr("performance"),
			}

			resourceGroupIdentityModel := &vmwarev1.ResourceGroupIdentity{
				ID: core.StringPtr("testString"),
			}

			createVdcOptions := &vmwarev1.CreateVdcOptions{
				Name: core.StringPtr("testString"),
				DirectorSite: vdcDirectorSitePrototypeModel,
				Edge: vdcEdgePrototypeModel,
				FastProvisioningEnabled: core.BoolPtr(true),
				ResourceGroup: resourceGroupIdentityModel,
				Cpu: core.Int64Ptr(int64(0)),
				Ram: core.Int64Ptr(int64(0)),
				RhelByol: core.BoolPtr(false),
				WindowsByol: core.BoolPtr(false),
				AcceptLanguage: core.StringPtr("testString"),
			}

			vdc, response, err := vmwareService.CreateVdc(createVdcOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(vdc).ToNot(BeNil())
		})
	})

	Describe(`GetVdc - Get a virtual data center`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetVdc(getVdcOptions *GetVdcOptions)`, func() {
			getVdcOptions := &vmwarev1.GetVdcOptions{
				ID: core.StringPtr("testString"),
				AcceptLanguage: core.StringPtr("testString"),
			}

			vdc, response, err := vmwareService.GetVdc(getVdcOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(vdc).ToNot(BeNil())
		})
	})

	Describe(`UpdateVdc - Update a virtual data center`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UpdateVdc(updateVdcOptions *UpdateVdcOptions)`, func() {
			vdcPatchModel := &vmwarev1.VDCPatch{
				Cpu: core.Int64Ptr(int64(0)),
				FastProvisioningEnabled: core.BoolPtr(true),
				Ram: core.Int64Ptr(int64(0)),
			}
			vdcPatchModelAsPatch, asPatchErr := vdcPatchModel.AsPatch()
			Expect(asPatchErr).To(BeNil())

			updateVdcOptions := &vmwarev1.UpdateVdcOptions{
				ID: core.StringPtr("testString"),
				VDCPatch: vdcPatchModelAsPatch,
				AcceptLanguage: core.StringPtr("testString"),
			}

			vdc, response, err := vmwareService.UpdateVdc(updateVdcOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(vdc).ToNot(BeNil())
		})
	})

	Describe(`DeleteVdc - Delete a virtual data center`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteVdc(deleteVdcOptions *DeleteVdcOptions)`, func() {
			deleteVdcOptions := &vmwarev1.DeleteVdcOptions{
				ID: core.StringPtr("testString"),
				AcceptLanguage: core.StringPtr("testString"),
			}

			vdc, response, err := vmwareService.DeleteVdc(deleteVdcOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(vdc).ToNot(BeNil())
		})
	})

	Describe(`DeleteDirectorSitesPvdcsCluster - Delete a cluster`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteDirectorSitesPvdcsCluster(deleteDirectorSitesPvdcsClusterOptions *DeleteDirectorSitesPvdcsClusterOptions)`, func() {
			deleteDirectorSitesPvdcsClusterOptions := &vmwarev1.DeleteDirectorSitesPvdcsClusterOptions{
				SiteID: core.StringPtr("testString"),
				ID: core.StringPtr("testString"),
				PvdcID: core.StringPtr("testString"),
				AcceptLanguage: core.StringPtr("testString"),
				XGlobalTransactionID: core.StringPtr("testString"),
			}

			clusterSummary, response, err := vmwareService.DeleteDirectorSitesPvdcsCluster(deleteDirectorSitesPvdcsClusterOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(clusterSummary).ToNot(BeNil())
		})
	})

	Describe(`DeleteDirectorSite - Delete a director site instance`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteDirectorSite(deleteDirectorSiteOptions *DeleteDirectorSiteOptions)`, func() {
			deleteDirectorSiteOptions := &vmwarev1.DeleteDirectorSiteOptions{
				ID: core.StringPtr("testString"),
				AcceptLanguage: core.StringPtr("testString"),
				XGlobalTransactionID: core.StringPtr("testString"),
			}

			directorSite, response, err := vmwareService.DeleteDirectorSite(deleteDirectorSiteOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(directorSite).ToNot(BeNil())
		})
	})
})

//
// Utility functions are declared in the unit test file
//
