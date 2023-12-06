import HistoryIcon from "../icons/history-icon";
import HomeIcon from "../icons/home-icon";
import TicketIcon from "../icons/ticket-icon";

export const MainNavLink = [
  { title: "Home", path: "/dashboard", icon: <HomeIcon /> },
  { title: "Tickets", path: "/dashboard/tickets", icon: <TicketIcon /> },
  { title: "History", path: "/dashboard/history", icon: <HistoryIcon /> },
];

export const SubNavLink = [
  { title: "Users", path: "/admin/users" },
  { title: "Staff", path: "/admin/staff" },
];
