import CompanyCard from './CompanyCard';

import style from '@/app/styles/company/CompanyCards.module.css';

import { Company } from '@/app/lib/types';

const CompanyData: Company[] = [
  {
    id: '1',
    name: 'Google',
    company_portal: 'https://www.google.com/',
  },
  {
    id: '2',
    name: 'Facebook',
    company_portal: 'https://www.facebook.com/',
  },
  {
    id: '3',
    name: 'Amazon',
    company_portal: 'https://www.amazon.com/',
  },
  {
    id: '4',
    name: 'Apple',
    company_portal: 'https://www.apple.com/',
  },
  {
    id: '5',
    name: 'Microsoft',
    company_portal: 'https://www.microsoft.com/',
  },
  {
    id: '6',
    name: 'Netflix',
    company_portal: 'https://www.netflix.com/',
  },
];

export default function CompanyCards() {
  // Change this to a lib call to get the companies from the API
  const companies = CompanyData;

  return (
    <div className={style.cards}>
      {companies.map((company: Company) => (
        <CompanyCard key={company.id} company={company} />
      ))}
    </div>
  );
}
