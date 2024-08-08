const API_URL = import.meta.env.VITE_API_URL as string;

function deleteAllCookies() {
    // Get all cookies as a single string
    const cookies = document.cookie.split(";");

    // Loop through all cookies
    for (let i = 0; i < cookies.length; i++) {
        // Get cookie name
        const cookie = cookies[i].split("=");
        const cookieName = cookie[0].trim();

        // Set cookie expiration date to past
        document.cookie = cookieName + "=; expires=Thu, 01 Jan 1970 00:00:00 GMT; path=/";
    }
}

export async function login(username: string, password: string) {
    deleteAllCookies();
    const response = await fetch(`${API_URL}/login`, {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({ username, password }),
        credentials: "include",
    });
    return response;
}

function getToken() {
    const cookies = document.cookie.split("; ");
    const tokenRegex = /.*token=.*/;
    const tokenRow = cookies.filter(row => tokenRegex.test(row)) as string[];
    if (tokenRow.length === 0) {
        throw new Error("Unable to get cookie");
    }
    const token = tokenRow[0].split("=")[1];
    return token;
}

async function apiFetch(url: string, method="GET") {
    const token = getToken();
    return await fetch(url, {
        method: method,
        headers: {
            Authorization: token,
        }
    })
}

// async function apiPost(url: string, body: string) {
//     const token = getToken();
//     return await fetch(url, {
//         method: 'POST',
//         body,
//         headers: {
//             Authorization: token,
//         }
//     })
// }

interface MeResponse {
    username: string;
}

export async function getMe() {
    const url = `${API_URL}/me`;
    const response = await apiFetch(url);
    console.log(response);
    if (!response.ok) {
        return undefined;
    }
    const data = await response.json() as MeResponse;
    return data.username;
}