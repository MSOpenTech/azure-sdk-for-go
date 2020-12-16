package marketplacefulfillment

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

        // SaasSubscriptionStatus enumerates the values for saas subscription status.
    type SaasSubscriptionStatus string

    const (
                // NotStarted ...
        NotStarted SaasSubscriptionStatus = "NotStarted"
                // PendingFulfillmentStart ...
        PendingFulfillmentStart SaasSubscriptionStatus = "PendingFulfillmentStart"
                // Subscribed ...
        Subscribed SaasSubscriptionStatus = "Subscribed"
                // Suspended ...
        Suspended SaasSubscriptionStatus = "Suspended"
                // Unsubscribed ...
        Unsubscribed SaasSubscriptionStatus = "Unsubscribed"
            )
    // PossibleSaasSubscriptionStatusValues returns an array of possible values for the SaasSubscriptionStatus const type.
    func PossibleSaasSubscriptionStatusValues() []SaasSubscriptionStatus {
        return []SaasSubscriptionStatus{NotStarted,PendingFulfillmentStart,Subscribed,Suspended,Unsubscribed}
    }

        // Status enumerates the values for status.
    type Status string

    const (
                // StatusConflict ...
        StatusConflict Status = "Conflict"
                // StatusFailed ...
        StatusFailed Status = "Failed"
                // StatusInProgress ...
        StatusInProgress Status = "InProgress"
                // StatusNotStarted ...
        StatusNotStarted Status = "NotStarted"
                // StatusSucceeded ...
        StatusSucceeded Status = "Succeeded"
            )
    // PossibleStatusValues returns an array of possible values for the Status const type.
    func PossibleStatusValues() []Status {
        return []Status{StatusConflict,StatusFailed,StatusInProgress,StatusNotStarted,StatusSucceeded}
    }

