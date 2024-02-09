'use server';

import { revalidateTag } from 'next/cache';
import { Company } from './types';

import { SERVER_URL } from './consts';

export async function getCompanies() : Promise<Company[] | null> {
  
  try {
      const res = await fetch(`${SERVER_URL}/companies`, {
        next: { tags: ['companies'] }
      });
    
      if (!res.ok) {
        throw new Error(`HTTP error! status: ${res.status}`);
      }

      const data = await res.json();

      console.log(data);

      return data.companies;
  } catch (error) {
    console.error(error);
    return null;
  }
}

export async function addCompany(formData: FormData) : Promise<Company | null> {
  //maybe add some sort of extra validation here
  const name = formData.get('name') as string;
  const candidate_portal = formData.get('candidate_portal') as string;

  const company: Company = {
    name: name,
    candidate_portal: candidate_portal,
  };

  try {
    const res = await fetch(`${SERVER_URL}/companies`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(company),
    });

    if (!res.ok) {
      throw new Error(`HTTP error! status: ${res.status}`);
    }

    const data = await res.json();

    revalidateTag('companies');
    
    return data;

  } catch (error) {
    console.error(error);
    return null;
  }
}

export async function deleteCompany(id: string) {
  try {
    const res = await fetch(`${SERVER_URL}/companies/${id}`, {
      method: 'DELETE',
    });

    if (!res.ok) {
      throw new Error(`HTTP error! status: ${res.status}`);
    }

    revalidateTag('companies');
  } catch (error) {
    console.error(error);
  }
}

export async function updateCompany(id: string, formData: FormData) {
  const name = formData.get('name') as string;
  const candidate_portal = formData.get('candidate_portal') as string;

  try {
    const res = await fetch(`${SERVER_URL}/companies/${id}`, {
      method: 'PUT',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({ name, candidate_portal }),
    });

    if (!res.ok) {
      throw new Error(`HTTP error! status: ${res.status}`);
    }

    revalidateTag('companies');

  } catch (error) {
    console.error(error);
  }
}
