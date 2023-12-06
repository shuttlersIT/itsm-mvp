import Link from "next/link";

export default function Home() {
  return (
    <main>
      <Link href="/dashboard">
        <p className="text-center hover:bg-[#F4F4F4] py-2.5 px-2.5 mt-10 w-max mx-auto rounded-3xl shadow-[0px_0px_32px_0px_rgba(204,204,204,0.25)]">
          Click here to Sign In!
        </p>
      </Link>
    </main>
  );
}
