import { Button } from "@/components/ui/button";
import AddIcon from "../icons/add-icon";
import { Avatar, AvatarFallback, AvatarImage } from "@/components/ui/avatar";
import { usePathname } from "next/navigation";
import CustomButton from "./button";

const MainNav = () => {
  const pathname = usePathname();
  const tickets = "/dashboard/tickets";
  const history = "/dashboard/history";
  const users = "/admin/users";
  const staff = "/admin/staff";
  let title = "Welcome, IT!";

  if (pathname === tickets) {
    title = "Tickets";
  } else if (pathname === history) {
    title = "History";
  } else if (pathname === users) {
    title = "Users";
  } else if (pathname === staff) {
    title = "Staff";
  }

  return (
    <nav
      className={`flex items-center justify-between xl:items-center py-6 mb-7`}
    >
      <div className="md:block xl:gap-0 gap-2.5 flex flex-col">
        {/* <MenuIcon className="md:hidden" onClick={handleClick} /> */}
        <p className="text-[#7B7B7B] text-xs">3rd October, 2023</p>
        <h1 className="md:text-2xl text-lg text-[#0A0A0A]">{title}</h1>
      </div>
      <div className="flex gap-2 md:gap-7">
        <CustomButton>
          <AddIcon />
          <p className="hidden md:block">Create New</p>
        </CustomButton>
        <Avatar>
          <AvatarImage src="" />
          <AvatarFallback className="bg-[#F4F4F4] text-xs md:text-[0.95rem] text-[#101723]">
            IT
          </AvatarFallback>
        </Avatar>
      </div>
    </nav>
  );
};

export default MainNav;
