export default function Layout({
  children,
  visitors,
}: {
  children: React.ReactNode;
  visitors: React.ReactNode;
}) {
  return (
    <div className="grid grid-cols-1 md:grid-cols-2 xl:grid-cols-3 gap-4">
      {visitors}
      {children}
    </div>
  );
}
