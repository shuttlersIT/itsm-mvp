"use client";

import Image from "next/image";
import logo from "@/public/logo.png";
import Link from "next/link";
import { usePathname } from "next/navigation";
import { MainNavLink, SubNavLink } from "./navlinks";

/* 72px - Narrow 768px-1263px  */
/* 244px - Medium 1264px-1919 */
/* 355px - Wide 1920 */

const Sidebar = () => {
  const pathname = usePathname();
  return (
    <nav
      className={`px-4 py-7 border-r min-[768px]:max-[1263px]:w-[4.5rem] min-[1264px]:max-[1919px]:w-[14rem] min-[1920px]:w-[22.2rem] border-[#EBEBEB] md:block h-screen scroll fixed overflow-auto`}
    >
      <Link href="/">
        <div className="flex gap-2 mb-10">
          <Image src={logo} priority alt="logo" />
        </div>
      </Link>

      {MainNavLink.map((link) => (
        <Link
          href={link.path}
          key={link.title}
          className={`flex items-center gap-2.5 hover:bg-[#F7F7F7] rounded-xl py-2.5 px-5 cursor-pointer mb-5 ${
            pathname === link.path ? "bg-[#F7F7F7]" : "bg-transparent"
          }`}
        >
          {link.icon}
          <p
            className={`text-sm ${
              pathname === link.path ? "font-bold" : "font-medium"
            }`}
          >
            {link.title}
          </p>
        </Link>
      ))}

      <div className="text-[0.875rem] font-medium text-[#565656] mb-5 mt-16">
        <p
          className={`text-[0.85rem] ml-5 mb-5 text-[#060606] ${
            pathname.includes("admin") ? "font-bold" : "font-normal"
          }`}
        >
          ADMIN
        </p>

        {SubNavLink.map((link) => (
          <Link
            href={link.path}
            key={link.title}
            className={`flex items-center gap-4 mb-5 py-2.5 px-5 hover:bg-[#F7F7F7] hover:rounded-xl cursor-pointer ml-1 ${
              pathname === link.path ? "bg-[#F7F7F7]" : "bg-transparent"
            }`}
          >
            <p
              className={`text-sm ${
                pathname === link.path ? "font-bold" : "font-medium"
              }`}
            >
              {link.title}
            </p>
          </Link>
        ))}
      </div>
    </nav>
  );
};

export default Sidebar;
