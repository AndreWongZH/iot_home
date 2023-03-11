import { BottomNav } from "@/components/bottomNav";

export default function DashboardLayout({
  children, // will be a page or nested layout
}: {
  children: React.ReactNode,
}) {
  return (
    <>
      {children}
      <div className="h-24"></div>
      <BottomNav />
    </>
  );
}
