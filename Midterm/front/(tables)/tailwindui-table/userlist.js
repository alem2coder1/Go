import { useState, useEffect } from 'react';

const useFetchUsers = () => {
    const [userList, setUserList] = useState([]); // 用户数据状态
    const [loading, setLoading] = useState(true); // 加载状态

    useEffect(() => {
        const fetchUsers = async () => {
            try {
                const response = await fetch('http://localhost:8080/api/users'); // API 请求
                const data = await response.json(); // 解析 JSON 数据
                setUserList(data); // 设置用户数据
                setLoading(false); // 加载完成
            } catch (error) {
                console.error('Error fetching users:', error); // 错误处理
                setLoading(false); // 确保即使有错误也关闭加载状态
            }
        };

        fetchUsers(); // 执行 API 调用
    }, []); // 确保只在初次加载时调用

    return { userList, loading }; // 返回数据和加载状态
};

export default useFetchUsers;