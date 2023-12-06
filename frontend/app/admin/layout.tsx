"use client";

import MainNav from "../components/main-nav";
import Sidebar from "../components/sidebar";

export default function DashboardLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  return (
    <main className="">
      <Sidebar />
      <section className="max-w-[1440px] scroll min-[768px]:max-[1263px]:w-[calc(100%_-_4.5rem)] min-[1264px]:max-[1919px]:w-[calc(100%_-_14rem)] min-[1920px]:w-[calc(100%_-_22.2rem)] min-[768px]:max-[1263px]:ml-[4.5rem] min-[1264px]:max-[1919px]:ml-[14rem] min-[1920px]:ml-[22.2rem] px-10">
        <MainNav />
        {children}
      </section>
    </main>
  );
}
