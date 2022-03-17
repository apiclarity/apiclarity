import React from 'react';
import { useParams, useRouteMatch } from 'react-router-dom';
import BackRouteButton from 'components/BackRouteButton';
import Title from 'components/Title';
import TabbedPageContainer from 'components/TabbedPageContainer';
import PageContainer from 'components/PageContainer';
import Loader from 'components/Loader';
import { useFetch } from 'hooks';
import { formatDate } from 'utils/utils';
import Details from './Details';
import Specs from './Specs';
import { getPlugins, PLUGIN_TYPES } from 'plugins';

import './event-details.scss';

const EventDetails = () => {
    const {path, url, isExact} = useRouteMatch();
    const params = useParams();
    const {eventId} = params;

    const [{loading, data}] = useFetch(`apiEvents/${eventId}`);

    if (loading) {
        return <Loader />;
    }

    if (!data) {
        return null;
    }

    const {time} = data;
    const plugins = getPlugins(PLUGIN_TYPES.EVENT_DETAILS);
    const pluginTabs = plugins.map((p) => {
        return {
            title: p.name,
            linkTo: `${url}${p.endpoint}`,
            to: `${path}${p.endpoint}`,
            component: () => <p.component  {...{...data, eventId}}/>
        };
    });

    return (
        <div className="events-details-page">
            <BackRouteButton title="API events" path={url.replace(`/${eventId}`, "")} />
            <Title>{formatDate(time)}</Title>
            <div className="tabbed-container-wrapper">
                <PageContainer className="fixed-sidebar">
                    <div className="sidebar-heading">
                        <div className="title"><div>Event Summary</div></div>
                    </div>
                    <div className="sidebar-content">
                        <Details data={data} />
                    </div>
                </PageContainer>
                <div className="event-details-tab-container">
                    <TabbedPageContainer
                        items={[
                            { title: "Spec", linkTo: url, to: path, exact: true, component: () => <Specs data={data} /> },
                            ...pluginTabs
                        ]}
                        noContentMargin={isExact}
                    />
                </div>
            </div>
        </div>
    )
}

export default EventDetails;
