import Header from "@/src/app/components/Header";
import Sidebar from "@/src/app/components/Sidebar";
import NavigationBar from "@/src/app/components/Navigation";
import { GroupProvider } from "../contexts/GroupContext";

/**
 * Protected layout component that wraps authenticated pages.
 * Provides consistent layout structure with header, sidebar, and navigation.
 * Wraps content in GroupProvider context.
 */
export default function ProtectedLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  return (
    <div className="h-screen flex flex-col">
      <Header />
      <div className="flex-1 flex h-[calc(100vh-1200px)]">
        <GroupProvider>
          <Sidebar />
          <div className="flex flex-col w-full">
            <NavigationBar />
            <main className="flex-1 overflow-y-auto">{children}</main>
          </div>
        </GroupProvider>
      </div>
    </div>
  );
}
