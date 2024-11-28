import { useState, useEffect } from 'react';
import toast from "react-hot-toast";

const deleteUser = (id) => {
    fetch(`http://localhost:8080/api/users/${id}`, { // 动态替换 {id}
        method: 'DELETE', // 使用 DELETE 请求
    })
        .then((response) => {
            if (response.ok) {
                toast.success("】Delete Successful");
                window.location.reload();
            } else {
                console.error(`Failed to delete user with ID: ${id}`);
            }
        })
        .catch((error) => {
            console.error('Error deleting user:', error);
        });
};

export default deleteUser;