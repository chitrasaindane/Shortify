import { sendError } from "./error";
import { getAPIBaseURL } from "./url";

const API_BASE_URL = getAPIBaseURL();

// # "User" API #
const userAPI = {
    // # Get 'User'
    async getUser(clerkUserId) {
        const apiEndpoint = `${API_BASE_URL}/user/?clerk_user_id=${clerkUserId}`;
        const apiRequestOptions = {
            method: 'GET'
        }

        const response = await fetch(
            apiEndpoint,
            apiRequestOptions
        );

        const res = await response.json();

        if (!response.ok) {
            let errMsg = res?.error;
            if (!errMsg) {
                errMsg = '🚫 Client Error: Failed to get the user';
                console.log(errMsg);
                errMsg = "";
            }
            sendError(errMsg);
        }

        return res;
    },

    // # Update 'Username'
    async updateUsername(clerkUserId, username) {
        const apiEndpoint = `${API_BASE_URL}/user/username?clerk_user_id=${clerkUserId}&username=${encodeURIComponent(username)}`;
        const apiRequestOptions = {
            method: 'PUT'
        };

        const response = await fetch(
            apiEndpoint,
            apiRequestOptions
        );

        const res = await response.json();

        if (!response.ok) {
            let errMsg = res?.error;
            if (!errMsg) {
                errMsg = '🚫 Client Error: Failed to update the username';
                console.log(errMsg);
                errMsg = "";
            }
            sendError(errMsg);
        }

        return res;
    }
};

// # "Link" API #
const linkAPI = {
    // # Create 'Link'
    async createLink(clerkUserId, data) {
        const apiEndpoint = `${API_BASE_URL}/link/?clerk_user_id=${clerkUserId}`;
        const apiRequestOptions = {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(
                {
                    data
                }
            )
        };

        const response = await fetch(
            apiEndpoint,
            apiRequestOptions
        );

        const res = await response.json();

        if (!response.ok) {
            let errMsg = res?.error;
            if (!errMsg) {
                errMsg = '🚫 Client Error: Failed to create the link';
                console.log(errMsg);
                errMsg = "";
            }
            sendError(errMsg);
        }

        return res;
    },

    // # Update 'Link'
    async updateLink(clerkUserId, linkId, data) {
        const apiEndpoint = `${API_BASE_URL}/link/?clerk_user_id=${clerkUserId}&link_id=${linkId}`;
        const apiRequestOptions = {
            method: 'PUT',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(
                {
                    data
                }
            )
        };

        const response = await fetch(
            apiEndpoint,
            apiRequestOptions
        );

        const res = await response.json();

        if (!response.ok) {
            let errMsg = res?.error;
            if (!errMsg) {
                errMsg = '🚫 Client Error: Failed to update the link';
                console.log(errMsg);
                errMsg = "";
            }
            sendError(errMsg);
        }

        return res;
    },

    // # Delete 'Link'
    async deleteLink(linkId) {
        const apiEndpoint = `${API_BASE_URL}/link/?link_id=${linkId}`;
        const apiRequestOptions = {
            method: 'DELETE'
        }

        const response = await fetch(
            apiEndpoint,
            apiRequestOptions
        );

        const res = await response.json();

        if (!response.ok) {
            let errMsg = res?.error;
            if (!errMsg) {
                errMsg = '🚫 Client Error: Failed to delete the link';
                console.log(errMsg);
            }
            sendError(errMsg);
        }

        return res;
    },

    // # Get 'Link'
    async getLink(linkId) {
        const apiEndpoint = `${API_BASE_URL}/link/?link_id=${linkId}`;
        const apiRequestOptions = {
            method: 'GET'
        }

        const response = await fetch(
            apiEndpoint,
            apiRequestOptions
        );

        const data = await response.json();

        if (!response.ok) {
            let errMsg = res?.error;
            if (!errMsg) {
                errMsg = '🚫 Client Error: Failed to get the link';
                console.log(errMsg);
            }
            sendError(errMsg);
        }

        return data;
    },

    // # Get 'Links'
    async getLinks(clerkUserId) {
        const apiEndpoint = `${API_BASE_URL}/link/list?clerk_user_id=${clerkUserId}`;
        const apiRequestOptions = {
            method: 'GET'
        }

        const response = await fetch(
            apiEndpoint,
            apiRequestOptions
        );

        const data = await response.json();

        if (!response.ok) {
            let errMsg = res?.error;
            if (!errMsg) {
                errMsg = '🚫 Client Error: Failed to get the links';
                console.log(errMsg);
                errMsg = ""
            }
            sendError(errMsg);
        }

        return data;
    },
};

export {
    userAPI,
    linkAPI
};