const getUser = async () => {
    try {
        const response = await fetch(`http://localhost:8080/api/users/1`, {
            method: 'GET',
            headers: {
                'Content-Type': 'application/json',
            },
        });

        if (!response.ok) {
            throw new Error('Failed to fetch user data');
        }

        const userData = await response.json();
        return userData;

    } catch (error) {
        console.error("Error fetching user:", error);
        return null;
    }
};

export default getUser;