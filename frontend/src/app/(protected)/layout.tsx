import Header from "../components/Header";
import Sidebar from "../components/Sidebar";

export default function ProtectedLayout({
    children,
}: {
    children: React.ReactNode
}) {

    return (
        <div className="h-screen flex flex-col">
            <Header />
            <div className="flex-1 flex h-[calc(100vh-74px)]">
                <Sidebar/>
                <main className="flex-1 p-8">
                    {children}
                </main>
            </div>
        </div>
    );
}
