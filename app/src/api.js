export function getDashboardStats(from, to) {
  const q = `?from=${from.toISOString()}&to=${to.toISOString()}`;
  return fetchApi('/api/dashboard/stats' + q);
}

export function getDashboardUserChart(from, to) {
  const q = `?from=${from.toISOString()}&to=${to.toISOString()}`;
  return fetchApi('/api/dashboard/userchart' + q);
}

export function getDashboardActivityChart(from, to) {
  const q = `?from=${from.toISOString()}&to=${to.toISOString()}`;
  return fetchApi('/api/dashboard/activitychart' + q);
}

function fetchApi(url, options = {}) {
  const defaulOptions = {
    /*credentials: 'include'*/
  };

  return fetch(url, Object.assign(defaulOptions, options)).then(response => {
    if (response.ok) {
      const res = response.json();
      return res;
    } else {
      console.log(response);
      return response.text().then(err => {
        return Promise.reject({
          message: err,
          status: response.status
        });
      });
    }
  });
}
