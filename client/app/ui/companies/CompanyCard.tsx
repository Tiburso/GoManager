import { Company } from '@/app/lib/types';

export default function CompanyCard({ company }: { company: Company }) {
  // TODO: change this from tailwind to regular CSS

  return (
    <div className="flex flex-col items-center justify-center w-64 h-64 m-2 p-2 bg-white border-2 border-gray-300 rounded-md">
      <h2 className="text-2xl font-bold">{company.name}</h2>
      <p className="text-lg">{company.id}</p>
      <p className="text-lg">{company.name}</p>
      <p className="text-lg">{company.company_portal}</p>
    </div>
  );
}
