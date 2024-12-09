import Header from "../components/Header";

export default function ProtectedLayout({
    children,
}: {
    children: React.ReactNode
}) {
    return (
        <>
            <Header />
            <main>{children}</main>
        </>
    );
}
