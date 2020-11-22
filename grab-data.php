<?php

class Helper
{
    public static function getMetaData($link, $all = false)
    {
        $userAgent = 'Mozilla/5.0 (Windows NT 5.1; rv:31.0) Gecko/20100101 Firefox/31.0';
        ini_set('user_agent', $userAgent);
        $html = new \DOMDocument();

        if (strpos($link, 'http') === false) {
            $link = 'http://' . $link;
        }
        // response
        $response = [];
        $sites_html = file_get_contents($link);

        $html = new \DOMDocument();
        @$html->loadHTML($sites_html);

        //Get all meta tags and loop through them.
        $response['title'] = $html->getElementsByTagName('title')[0]->textContent;
        $response['website'] = $link;
        $response['meta'] = [];
        foreach ($html->getElementsByTagName('meta') as $meta) {
            if ($all) {
                if($meta->getAttribute('name')) {
                    $response['meta'][$meta->getAttribute('name')] = $meta->getAttribute('content');
                }else if($meta->getAttribute('rel')) {
                    $response['meta'][$meta->getAttribute('rel')] = $meta->getAttribute('content');
                }else if($meta->getAttribute('itemprop')) {
                    $response['meta'][$meta->getAttribute('itemprop')] = $meta->getAttribute('content');
                }else if($meta->getAttribute('property')) {
                    $response['meta'][$meta->getAttribute('property')] = $meta->getAttribute('content');
                }
            } else {
                if (Helper::has_meta_description($meta)) {
                    $response['description'] = $meta->getAttribute('content');
                };

                if (Helper::has_meta_image($meta)) {
                    $response['image'] = $meta->getAttribute('content');
                }

                if (Helper::has_meta_url($meta)) {
                    $response['url'] = $meta->getAttribute('content');
                }
                if ($meta->getAttribute('name') === 'twitter:domain') {
                    return $meta->getAttribute('content');
                }
            }

        }
        return $response;

    }
    public static function get_meta_editor($link)
    {
        $response = Helper::getMetaData('framer.com');
        return [
            "title" => $response['title'],
            "description" => $response['description'],
            "image" => [
                "url" => $response['image'],
            ],
        ];
    }

    public static function get_meta_title($meta)
    {
        // google
        $meta->getAttribute('itemprop') == 'title';
        // facebook
        $meta->getAttribute('property') == 'og:title';
        //twitter
        $meta->getAttribute('name') == 'twitter:title';
    }

    public static function has_meta_image($meta)
    {
        if ($meta->getAttribute('itemprop') == 'image' || $meta->getAttribute('property') == 'og:image' || $meta->getAttribute('name') == 'twitter:image' || $meta->getAttribute('name') == 'twitter:image:src') { // google
            return true;
        }
        return false;
    }

    public static function has_meta_description($meta)
    {
        if ($meta->getAttribute('itemprop') == 'description' || $meta->getAttribute('property') == 'og:description' || $meta->getAttribute('name') == 'twitter:description') { // google
            return true;
        }
        return false;
    }

    public static function has_meta_url($meta)
    {
        if ($meta->getAttribute('itemprop') == 'url' || $meta->getAttribute('property') == 'og:url' || $meta->getAttribute('name') == 'twitter:url') { // google
            return true;
        }
        return false;
    }
}
