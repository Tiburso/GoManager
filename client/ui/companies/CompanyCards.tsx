import CompanyCard from './CompanyCard';

import style from '@/styles/company/CompanyCards.module.css';

import { Company } from '@/lib/types';

export default function CompanyCards({ companies }: { companies: Company[] }) {
  return (
    <div className={style.cards}>
      {companies.map((company: Company) => (
        <CompanyCard key={company.id} company={company} />
      ))}
    </div>
  );
}
