'use server';

import { revalidatePath } from 'next/cache';
import { Company } from './types';

const CompanyData: Company[] = [
  {
    id: '1',
    name: 'Google',
    candidate_portal: 'https://www.google.com/',
  },
  {
    id: '2',
    name: 'Facebook',
    candidate_portal: 'https://www.facebook.com/',
  },
  {
    id: '3',
    name: 'Amazon',
    candidate_portal: 'https://www.amazon.com/',
  },
  {
    id: '4',
    name: 'Apple',
    candidate_portal: 'https://www.apple.com/',
  },
  {
    id: '5',
    name: 'Microsoft',
    candidate_portal: 'https://www.microsoft.com/',
  },
  {
    id: '6',
    name: 'Netflix',
    candidate_portal: 'https://www.netflix.com/',
  },
];

export async function getCompanies() {
  return CompanyData;
}

export async function addCompany(formData: FormData) {
  //maybe add some sort of extra validation here
  const name = formData.get('name') as string;
  const candidate_portal = formData.get('candidate_portal') as string;

  const company: Company = {
    id: (CompanyData.length + 1).toString(),
    name: name,
    candidate_portal: candidate_portal,
  };

  CompanyData.push(company);

  revalidatePath('/companies');
}
