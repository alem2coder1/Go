export const loginUser = async (email, password) => {
    try {
        const requestBody = JSON.stringify({ email, password });

        const response = await fetch("http://localhost:8080/api/users/login", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
            },
            body: requestBody,
        });

        const contentType = response.headers.get("content-type");
        if (contentType && contentType.includes("application/json")) {

            const data = await response.json();
            if (!response.ok) {
                throw new Error(data.error || "Login failed");
            }
            const jwtToken = data.token;
            if (!jwtToken) {
                throw new Error("JWT token is missing in the response");
            }
            const decodedPayload = parseJwt(jwtToken);
            localStorage.setItem("jwtToken", jwtToken);

            return data;
        } else {
            const text = await response.text();
            throw new Error(text || "Unexpected response format");
        }
    } catch (error) {
        console.error("Error during login:", error.message);
        throw new Error(error.message);
    }
};

function parseJwt(token) {
    if (!token || token.split('.').length !== 3) {
        throw new Error("Invalid JWT token format");
    }

    const base64Url = token.split('.')[1];
    const base64 = base64Url.replace(/-/g, '+').replace(/_/g, '/');
    const jsonPayload = decodeURIComponent(atob(base64).split('').map(function(c) {
        return '%' + ('00' + c.charCodeAt(0).toString(16)).slice(-2);
    }).join(''));

    return JSON.parse(jsonPayload);
}

function getRoleFromJwt() {
    const jwtToken = localStorage.getItem("jwtToken");
    if (jwtToken) {
        const decodedPayload = parseJwt(jwtToken);
        return decodedPayload.role;  // 返回角色信息
    }
    return null;
}