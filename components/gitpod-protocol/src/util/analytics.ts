/**
 * Copyright (c) 2021 Gitpod GmbH. All rights reserved.
 * Licensed under the GNU Affero General Public License (AGPL).
 * See License-AGPL.txt in the project root for license information.
 */

import Analytics = require("analytics-node");
import { log } from './logging';

export const IAnalyticsWriter = Symbol("IAnalyticsWriter");

type Identity =
    | { userId: string | number }
    | { userId?: string | number; anonymousId: string | number };

export type IdentifyMessage = Identity & {
    traits?: any;
    timestamp?: Date;
    context?: any;
};

export type TrackMessage = Identity & {
    event: string;
    properties?: any;
    timestamp?: Date;
    context?: any;
};

export function newAnalyticsWriterFromEnv(): IAnalyticsWriter {
    if (process.env.ANALYTICS_WRITER === "segment") {
        return new SegmentAnalyticsWriter(process.env.SEGMENT_WRITE_KEY || "");
    } else {
        return new LogAnalyticsWriter();
    }
}

export interface IAnalyticsWriter {

    identify(msg: IdentifyMessage): void;

    track(msg: TrackMessage): void;

}

export class SegmentAnalyticsWriter implements IAnalyticsWriter {
    
    protected readonly analytics: Analytics;

    constructor(writeKey: string) {
        this.analytics = new Analytics(writeKey);
    }

    identify(msg: IdentifyMessage) {
        this.analytics.identify(msg);
    }

    track(msg: TrackMessage) {
        this.track(msg);
    }

}

export class LogAnalyticsWriter implements IAnalyticsWriter {

    identify(msg: IdentifyMessage): void {
        log.debug("analytics identify", msg);
    }
    track(msg: TrackMessage): void {
        log.debug("analytics track", msg);
    }

}
