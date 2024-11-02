function GrafanaDashboard() {
    return (
        <iframe
            src="http://localhost:3000/d/YOUR_DASHBOARD_UID/YOUR_DASHBOARD_NAME"
            width="100%"
            height="600"
            frameBorder="0"
        />
    );
}

export default GrafanaDashboard;
