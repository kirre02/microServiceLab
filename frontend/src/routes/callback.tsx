import { useNavigate } from '@tanstack/react-router';
import { useEffect } from 'react';

export default function Callback() {
    const navigate = useNavigate();

    useEffect(() => {
        const params = new URLSearchParams(window.location.search);
        const code = params.get("code");

        if (code) {
            fetch("/api/auth/callback", {
                method: "POST",
                headers: { "Content-Type": "application/json" },
                body: JSON.stringify({ code }),
            })
                .then(res => res.json())
                .then(data => {
                    localStorage.setItem("access_token", data.access_token);
                    navigate({ to: '/' });
                })
                .catch(() => navigate({ to: '/error' }));
        } else {
            navigate({ to: '/error' });
        }
    }, [navigate]);

    return <div>Processing login...</div>;
}
