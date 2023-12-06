"use client";
const Dashboard = () => {
  return (
    <div className="flex justify-between text-lg gap-20 font-bold">
      <div className="flex justify-between items-start w-full shadow-[0px_0px_32px_0px_rgba(204,204,204,0.25)] rounded-2xl px-10 py-20">
        <p>Open Tickets</p>
        <p>0</p>
      </div>
      <div className="flex justify-between items-start w-full shadow-[0px_0px_32px_0px_rgba(204,204,204,0.25)] px-10 py-20">
        <p className="">Closed Tickets</p>
        <p>0</p>
      </div>
    </div>
  );
};

export default Dashboard;
