// 解析 JWT 令牌以获取用户的 role
export function getRoleFromJwt() {
    const jwtToken = localStorage.getItem("jwtToken");
    if (jwtToken) {
        const decodedPayload = parseJwt(jwtToken);
        return decodedPayload.role;
    }
    return null;
}

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