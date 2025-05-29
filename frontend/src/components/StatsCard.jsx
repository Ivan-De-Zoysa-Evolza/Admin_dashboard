function StatsCard({ title, value }) {
  return (
    <div className="p-4 border rounded shadow">
      <h3 className="text-lg font-semibold">{title}</h3>
      <p className="text-2xl">{value !== undefined ? value : "Loading..."}</p>
    </div>
  );
}

export default StatsCard;
