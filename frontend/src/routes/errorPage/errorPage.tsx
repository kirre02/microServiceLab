import { Link } from '@tanstack/react-router';

export function ErrorPage() {
    return (
        <div className="error-container">
            <h1 className="error-title">Something went wrong</h1>
            <p className="error-message">We couldn't complete your request.</p>
            <Link to="/" className="error-link">
                Go back to Home
            </Link>
        </div>
    );
}

