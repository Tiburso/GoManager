import { Company } from '@/app/lib/types';

import style from '@/app/styles//company/CompanyCard.module.css';

export default function CompanyCard({ company }: { company: Company }) {
  // TODO: change this from tailwind to regular CSS

  return (
    <div className={style.card}>
      <h2>{company.name}</h2>
      <p>{company.company_portal}</p>
    </div>
  );
}
